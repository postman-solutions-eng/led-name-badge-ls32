import json
import os
import ssl
import urllib.error
import urllib.parse
import urllib.request

DEFAULT_POSTMAN_API_BASE_URL = 'https://api.getpostman.com'

_STATUS_SHORT = {
    'healthy': 'ok',
    'warning': 'warn',
    'critical': 'crit',
    'inactive': 'off',
}


class PostmanCatalogError(Exception):
    def __init__(self, message, status_code=502):
        super().__init__(message)
        self.status_code = status_code


def _ca_bundle_path():
    for module_name in ('certifi', 'pip._vendor.certifi'):
        try:
            module = __import__(module_name, fromlist=['where'])
            return module.where()
        except (ImportError, AttributeError):
            continue
    return None


def _ssl_context():
    ca_bundle = _ca_bundle_path()
    if ca_bundle:
        return ssl.create_default_context(cafile=ca_bundle)
    return ssl.create_default_context()


def _api_base_url():
    return os.environ.get('POSTMAN_API_BASE_URL', DEFAULT_POSTMAN_API_BASE_URL).rstrip('/')


def _api_request(api_key, path, params=None):
    query = f'?{urllib.parse.urlencode(params)}' if params else ''
    url = f'{_api_base_url()}{path}{query}'
    request = urllib.request.Request(
        url,
        headers={
            'Accept': 'application/json',
            'x-api-key': api_key,
        },
        method='GET',
    )
    try:
        with urllib.request.urlopen(request, timeout=15, context=_ssl_context()) as response:
            return json.loads(response.read().decode('utf-8'))
    except urllib.error.HTTPError as exc:
        details = exc.read().decode('utf-8', errors='replace')
        raise PostmanCatalogError(
            f'Postman API catalog request failed ({exc.code}): {details}',
            status_code=exc.code if 400 <= exc.code < 500 else 502,
        ) from exc
    except urllib.error.URLError as exc:
        raise PostmanCatalogError(f'Postman API catalog request failed: {exc.reason}') from exc


def list_system_environments(api_key, limit=20):
    payload = _api_request(api_key, '/api-catalog/system-environments', {'limit': limit})
    return payload.get('data') or []


def list_services(api_key, system_environment_id, limit=50):
    payload = _api_request(
        api_key,
        '/api-catalog/services',
        {
            'systemEnvironmentId': system_environment_id,
            'limit': limit,
        },
    )
    return payload.get('data') or [], (payload.get('meta') or {}).get('nextCursor')


def resolve_system_environment(api_key, preferred_id=None):
    if preferred_id:
        return {'id': preferred_id}

    environments = list_system_environments(api_key)
    if not environments:
        raise PostmanCatalogError('No system environments found in API catalog', status_code=404)

    for env in environments:
        if env.get('isProduction'):
            return env

    for env in environments:
        if env.get('associationCount', 0) > 0:
            return env

    return environments[0]


def _sanitize_display_part(value):
    if not value:
        return value
    return value.replace(':', '-')


def build_display_text(services, environment_name=None, has_more=False):
    label = _sanitize_display_part(environment_name) if environment_name else 'Catalog'
    prefix = f'{label} - '

    if not services:
        return f'{prefix}0 svcs :star:'

    counts = {}
    for service in services:
        status = service.get('status', 'unknown')
        counts[status] = counts.get(status, 0) + 1

    total = len(services)
    status_parts = []
    for status in ('healthy', 'warning', 'critical', 'inactive'):
        count = counts.get(status)
        if count:
            short = _STATUS_SHORT.get(status, status)
            status_parts.append(f'{count} {short}')

    status_summary = ', '.join(status_parts) if status_parts else f'{total} svcs'
    suffix = ' +more' if has_more else ''
    return f'{prefix}{total} svcs ({status_summary}){suffix} :star:'


def summarize_services(services):
    return [
        {
            'id': service.get('id'),
            'name': service.get('name'),
            'status': service.get('status'),
        }
        for service in services
    ]


def fetch_catalog_summary(api_key, system_environment_id=None):
    environment = resolve_system_environment(api_key, system_environment_id)
    services, next_cursor = list_services(api_key, environment['id'])
    text = build_display_text(
        services,
        environment_name=environment.get('name'),
        has_more=bool(next_cursor),
    )
    return {
        'text': text,
        'systemEnvironment': {
            'id': environment.get('id'),
            'name': environment.get('name'),
        },
        'serviceCount': len(services),
        'services': summarize_services(services),
        'hasMore': bool(next_cursor),
    }
