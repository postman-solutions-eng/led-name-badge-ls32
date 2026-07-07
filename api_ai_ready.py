from flask import Flask, request
from lednamebadge import SimpleTextAndIcons, LedNameBadge
from postman_catalog import PostmanCatalogError, fetch_catalog_summary
from array import array

import argparse
import threading
import queue
import importlib.util
import os
import base64
import struct
import zlib

DEFAULT_SUMMARY = "Open LED Badge - Free, hackable, and fun! :star: :heart:"

app = Flask(__name__)

# How many real pixels each LED dot becomes in the rendered PNG.
# Bump this for biggg pictures.
ICON_SCALE = 64
ICON_ON_COLOR = (0, 255, 0)
ICON_OFF_COLOR = (0, 0, 0)

# Values (case-insensitive) that turn the optional `meta` payload on for
# GET /predefined-icons. Absent/anything-else keeps the response lean.
_TRUTHY = {'1', 'true', 'yes', 'on'}


def _wants_truthy(value):
    return isinstance(value, str) and value.strip().lower() in _TRUTHY


def _icon_to_png_base64(data, cols, scale=ICON_SCALE):
    """Render an 11-row LED icon bitmap to a scaled-up PNG data URI.

    Icon bitmaps store `cols` byte-columns of 11 bytes each; each byte is a
    horizontal strip of 8 pixels with the most significant bit on the left.
    """
    rows = 11
    width = 8 * cols
    out_width = width * scale
    out_height = rows * scale

    # Build raw RGB scanlines (each prefixed by a 0 filter byte).
    raw = bytearray()
    for y in range(rows):
        line = bytearray()
        for x in range(width):
            block = x // 8
            bit = x % 8
            index = block * rows + y
            byte = data[index] if index < len(data) else 0
            color = ICON_ON_COLOR if (byte & (0x80 >> bit)) else ICON_OFF_COLOR
            line.extend(bytes(color) * scale)
        for _ in range(scale):
            raw.append(0)
            raw.extend(line)

    def _chunk(tag, body):
        return (struct.pack('>I', len(body)) + tag + body
                + struct.pack('>I', zlib.crc32(tag + body) & 0xffffffff))

    ihdr = struct.pack('>IIBBBBB', out_width, out_height, 8, 2, 0, 0, 0)
    png = (b'\x89PNG\r\n\x1a\n'
           + _chunk(b'IHDR', ihdr)
           + _chunk(b'IDAT', zlib.compress(bytes(raw), 9))
           + _chunk(b'IEND', b''))

    return 'data:image/png;base64,' + base64.b64encode(png).decode('ascii')


def _error(message, details, status):
    """Build a consistent machine-readable JSON error body + status tuple."""
    return {'error': message, 'details': str(details)}, status


# --- JSON error handlers -------------------------------------------------
# Guarantee every failure (including unmatched routes, wrong methods, and
# uncaught exceptions) returns the same JSON `ErrorResponse` shape instead of
# Flask's default HTML pages, so an AI agent can always parse the response.

@app.errorhandler(404)
def _handle_404(e):
    return _error('Not found', 'The requested resource does not exist', 404)


@app.errorhandler(405)
def _handle_405(e):
    return _error('Method not allowed', 'The HTTP method is not allowed for this resource', 405)


@app.errorhandler(500)
def _handle_500(e):
    return _error('Internal server error', 'An unexpected error occurred', 500)


def _process_and_write(text, command_queue=None, write_hardware=True):
    """Create the scene buffer for `text` and either write to hardware,
    post to the mock console output via `command_queue`, or both depending on flags.
    """
    creator = SimpleTextAndIcons()
    scene_bitmap = creator.bitmap(text)

    buf = array('B')
    buf.extend(LedNameBadge.header([scene_bitmap[1]], [4], [0], [0], [0], 100))
    buf.extend(scene_bitmap[0])

    if write_hardware:
        try:
            LedNameBadge.write(buf)
        except Exception as e:
            print(f"_process_and_write: hardware write failed: {e}")

    if command_queue is not None:
        # API mock expects only `text` updates
        try:
            command_queue.put({'type': 'update', 'data': {'text': text}})
        except Exception as e:
            print(f"_process_and_write: enqueue failed: {e}")


@app.route('/display-text', methods=['POST'])
def display_text():
    # silent=True so a missing / non-JSON body yields {} instead of crashing
    # with AttributeError (which Flask would surface as an HTML 500).
    data = request.get_json(silent=True) or {}
    text = data.get('text', '')

    # `text` is required and must be non-empty: an agent that omits it should
    # get a clear 400, not a silent 200 that displays a blank badge.
    if not isinstance(text, str) or not text.strip():
        return _error(
            'Missing required field',
            "'text' is required and must be a non-empty string",
            400,
        )

    try:
        # Validate and prepare scene; actual write/mocking handled below.
        creator = SimpleTextAndIcons()
        creator.bitmap(text)
    except (KeyError, ValueError) as e:
        # Genuinely bad input (e.g. unsupported unicode) -> client error.
        return _error('Invalid display string format', e, 400)
    except (FileNotFoundError, OSError) as e:
        # Missing custom image / filesystem / hardware issue -> server error,
        # so an agent escalates instead of retrying a valid request forever.
        return _error('Hardware or file error', e, 500)
    except Exception as e:
        return _error('Internal server error', e, 500)

    # On actual run, the main program will decide whether to write to
    # hardware and/or the mock console by providing globals.
    global _API_COMMAND_QUEUE, _API_WRITE_HARDWARE
    _process_and_write(text, command_queue=globals().get('_API_COMMAND_QUEUE'), write_hardware=globals().get('_API_WRITE_HARDWARE', True))

    return {'status': 'Text displayed on LED', 'text': text}, 200


@app.route('/predefined-icons', methods=['GET'])
def get_predefined_icons():
    # The rendered PNG previews (`meta`) are opt-in: they are large and most
    # callers only need the icon codes. Pass ?meta=true (or 1/yes/on) to include
    # them. By default the `meta` field is omitted entirely.
    include_meta = _wants_truthy(request.args.get('meta'))

    try:
        creator = SimpleTextAndIcons()
        icons = [f':{name}:' for name in creator.bitmap_named.keys()]
        response = {'icons': icons}
        if include_meta:
            response['meta'] = [
                {'name': name, 'image': _icon_to_png_base64(data, cols)}
                for name, (data, cols, _ctrl) in creator.bitmap_named.items()
            ]
        return response, 200
    except Exception as e:
        return _error('Internal server error', e, 500)


def _get_api_key(data):
    key = request.headers.get('X-API-Key') or (data or {}).get('apiKey')
    if isinstance(key, str):
        key = key.strip()
    return key or None


@app.route('/display-summary', methods=['POST'])
def display_summary():
    data = request.get_json(silent=True) or {}
    api_key = _get_api_key(data)
    catalog_summary = None

    if api_key:
        try:
            catalog_summary = fetch_catalog_summary(
                api_key,
                data.get('systemEnvironmentId'),
            )
            summary = catalog_summary['text']
        except PostmanCatalogError as e:
            return _error('Failed to fetch API catalog', e, e.status_code)
    else:
        summary = DEFAULT_SUMMARY

    try:
        creator = SimpleTextAndIcons()
        creator.bitmap(summary)
    except (KeyError, ValueError) as e:
        return _error('Invalid display string format', e, 400)
    except (FileNotFoundError, OSError) as e:
        return _error('Hardware or file error', e, 500)
    except Exception as e:
        return _error('Internal server error', e, 500)

    global _API_COMMAND_QUEUE, _API_WRITE_HARDWARE
    _process_and_write(summary, command_queue=globals().get('_API_COMMAND_QUEUE'), write_hardware=globals().get('_API_WRITE_HARDWARE', True))

    response = {'status': 'Summary displayed on LED'}
    if catalog_summary is not None:
        response.update({
            'source': 'postman-api-catalog',
            'text': summary,
            'systemEnvironment': catalog_summary['systemEnvironment'],
            'serviceCount': catalog_summary['serviceCount'],
            'services': catalog_summary['services'],
            'hasMore': catalog_summary['hasMore'],
        })
    return response, 200


def _load_mock_console_module():
    """Dynamically load `mock-led-display.py` as module `mock_console`.
    This avoids Python import issues with hyphens in the filename.
    """
    base = os.path.dirname(os.path.abspath(__file__))
    path = os.path.join(base, 'mock-led-display.py')
    if not os.path.exists(path):
        raise FileNotFoundError(path)
    spec = importlib.util.spec_from_file_location('mock_console', path)
    module = importlib.util.module_from_spec(spec)
    spec.loader.exec_module(module)
    return module


def main():
    parser = argparse.ArgumentParser(description='LED Name Badge API server (AI-ready)')
    parser.add_argument('--host', default='0.0.0.0')
    parser.add_argument('--port', type=int, default=5001)
    parser.add_argument('--mock', action='store_true', help='Run with mock console instead of writing to hardware')
    parser.add_argument('--both', action='store_true', help='Write to hardware and also show mock console')
    args = parser.parse_args()

    # Decide behavior
    use_mock = args.mock or args.both
    write_hardware = not args.mock

    if use_mock:
        # Prepare shared state and command queue for the console mock
        cmd_q = queue.Queue()
        globals()['_API_COMMAND_QUEUE'] = cmd_q
        globals()['_API_WRITE_HARDWARE'] = args.both

        # Start Flask server in background thread, then run console in foreground
        server_thread = threading.Thread(
            target=app.run,
            kwargs={'host': args.host, 'port': args.port, 'threaded': True, 'use_reloader': False},
            daemon=True,
        )
        server_thread.start()

        # Load and run console (this will block in the main thread)
        mock_console = _load_mock_console_module()
        # mock_console.run_mock accepts display_state and command_queue optionally; pass the queue
        mock_console.run_mock(display_state=None, command_queue=cmd_q)
    else:
        # Default: run API server and write to hardware as before
        globals()['_API_WRITE_HARDWARE'] = True
        app.run(host=args.host, port=args.port)


if __name__ == '__main__':
    main()
