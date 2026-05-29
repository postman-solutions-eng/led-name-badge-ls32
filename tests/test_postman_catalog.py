from unittest import TestCase
from unittest.mock import patch

from postman_catalog import (
    PostmanCatalogError,
    build_display_text,
    fetch_catalog_summary,
    resolve_system_environment,
    summarize_services,
)


class TestPostmanCatalog(TestCase):
    def test_build_display_text_with_status_breakdown(self):
        services = [
            {'name': 'payments', 'status': 'healthy'},
            {'name': 'billing', 'status': 'warning'},
            {'name': 'users', 'status': 'healthy'},
        ]
        text = build_display_text(services, environment_name='Production')
        self.assertEqual(text, 'Production - 3 svcs (2 ok, 1 warn) :star:')

    def test_build_display_text_empty_services(self):
        text = build_display_text([], environment_name='Staging')
        self.assertEqual(text, 'Staging - 0 svcs :star:')

    def test_build_display_text_has_more_suffix(self):
        services = [{'name': 'payments', 'status': 'healthy'}]
        text = build_display_text(services, has_more=True)
        self.assertEqual(text, 'Catalog - 1 svcs (1 ok) +more :star:')

    def test_summarize_services(self):
        services = [
            {'id': '111', 'name': 'payments', 'status': 'healthy', 'traffic': {}},
            {'id': '222', 'name': 'billing', 'status': 'warning', 'traffic': {}},
        ]
        self.assertEqual(
            summarize_services(services),
            [
                {'id': '111', 'name': 'payments', 'status': 'healthy'},
                {'id': '222', 'name': 'billing', 'status': 'warning'},
            ],
        )

    def test_resolve_system_environment_prefers_production(self):
        environments = [
            {'id': 'a', 'name': 'Dev', 'isProduction': False, 'associationCount': 5},
            {'id': 'b', 'name': 'Production', 'isProduction': True, 'associationCount': 0},
        ]
        with patch('postman_catalog.list_system_environments', return_value=environments):
            env = resolve_system_environment('PMAK-test')
        self.assertEqual(env['id'], 'b')

    def test_resolve_system_environment_uses_associations_when_no_production(self):
        environments = [
            {'id': 'a', 'name': 'Dev', 'isProduction': False, 'associationCount': 0},
            {'id': 'b', 'name': 'QA', 'isProduction': False, 'associationCount': 2},
        ]
        with patch('postman_catalog.list_system_environments', return_value=environments):
            env = resolve_system_environment('PMAK-test')
        self.assertEqual(env['id'], 'b')

    def test_resolve_system_environment_raises_when_empty(self):
        with patch('postman_catalog.list_system_environments', return_value=[]):
            with self.assertRaises(PostmanCatalogError) as ctx:
                resolve_system_environment('PMAK-test')
        self.assertEqual(ctx.exception.status_code, 404)

    def test_sanitize_display_part_replaces_colons(self):
        from postman_catalog import _sanitize_display_part

        self.assertEqual(_sanitize_display_part('env:prod'), 'env-prod')

    def test_build_display_text_avoids_colon_icon_parsing(self):
        counts = {'healthy': 8, 'warning': 5, 'critical': 16, 'inactive': 16}
        services = []
        for status, count in counts.items():
            services.extend({'name': f'svc-{i}', 'status': status} for i in range(count))

        text = build_display_text(services, environment_name='Production', has_more=True)
        self.assertEqual(
            text,
            'Production - 45 svcs (8 ok, 5 warn, 16 crit, 16 off) +more :star:',
        )
        self.assertEqual(text.count(':'), 2)

    def test_ssl_context_uses_ca_bundle_when_available(self):
        with patch('postman_catalog._ca_bundle_path', return_value='/tmp/cacert.pem'), patch(
            'postman_catalog.ssl.create_default_context',
        ) as create_context:
            from postman_catalog import _ssl_context

            _ssl_context()
        create_context.assert_called_once_with(cafile='/tmp/cacert.pem')

    def test_fetch_catalog_summary(self):
        environment = {'id': 'env-1', 'name': 'Production'}
        services = [
            {'id': 'svc-1', 'name': 'payments', 'status': 'healthy'},
            {'id': 'svc-2', 'name': 'billing', 'status': 'critical'},
        ]
        with patch('postman_catalog.resolve_system_environment', return_value=environment), patch(
            'postman_catalog.list_services',
            return_value=(services, 'next-page'),
        ):
            result = fetch_catalog_summary('PMAK-test')

        self.assertEqual(result['serviceCount'], 2)
        self.assertTrue(result['hasMore'])
        self.assertEqual(result['text'], 'Production - 2 svcs (1 ok, 1 crit) +more :star:')
        self.assertEqual(result['systemEnvironment'], {'id': 'env-1', 'name': 'Production'})
