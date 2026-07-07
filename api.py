from flask import Flask, request
from lednamebadge import SimpleTextAndIcons, LedNameBadge
from postman_catalog import PostmanCatalogError, fetch_catalog_summary
from array import array

import argparse
import threading
import queue
import importlib.util
import os

DEFAULT_SUMMARY = "Open LED Badge - Free, hackable, and fun! :star: :heart:"

app = Flask(__name__)


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
    data = request.get_json()
    text = data.get('text', '')

    try:
        # Validate and prepare scene; actual write/mocking handled in main
        creator = SimpleTextAndIcons()
        creator.bitmap(text)
    except (KeyError, ValueError, FileNotFoundError, OSError) as e:
        return {'error': 'Invalid display string format', 'details': str(e)}, 400
    except Exception as e:
        return {'error': 'Invalid display string format', 'details': str(e)}, 400

    # On actual run, the main program will decide whether to write to
    # hardware and/or the mock console by providing globals.
    global _API_COMMAND_QUEUE, _API_WRITE_HARDWARE
    _process_and_write(text, command_queue=globals().get('_API_COMMAND_QUEUE'), write_hardware=globals().get('_API_WRITE_HARDWARE', True))

    return {'status': 'Text displayed on LED', 'text': text}, 200


@app.route('/predefined-icons', methods=['GET'])
def get_predefined_icons():
    creator = SimpleTextAndIcons()
    icons = [f':{name}:' for name in creator.bitmap_named.keys()]
    return {'icons': icons}, 200


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
            return {'error': 'Failed to fetch API catalog', 'details': str(e)}, e.status_code
    else:
        summary = DEFAULT_SUMMARY

    try:
        creator = SimpleTextAndIcons()
        creator.bitmap(summary)
    except (KeyError, ValueError, FileNotFoundError, OSError) as e:
        return {'error': 'Invalid display string format', 'details': str(e)}, 400

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
    parser = argparse.ArgumentParser(description='LED Name Badge API server')
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
