package predefinedicons

import (
	"context"
	restClient "example.com/led-display-api/sdk/internal/clients/rest"
	"example.com/led-display-api/sdk/internal/clients/rest/hooks"
	"example.com/led-display-api/sdk/internal/clients/rest/httptransport"
	"example.com/led-display-api/sdk/internal/configmanager"
	"example.com/led-display-api/sdk/leddisplayapiconfig"
	"example.com/led-display-api/sdk/sharedmodels"
	"time"
)

// Service provides methods to interact with PredefinedIcons-related API endpoints.
// It uses a configuration manager for settings and supports custom hooks for request/response interception.
type Service struct {
	manager                  *configmanager.ConfigManager
	hook                     hooks.Hook
	getPredefinedIconsConfig []leddisplayapiconfig.RequestOption
}

func NewService() *Service {
	return &Service{
		manager: configmanager.NewConfigManager(leddisplayapiconfig.Config{}),
	}
}

// WithConfigManager sets the configuration manager for this service.
// Returns the service instance for method chaining.
func (api *Service) WithConfigManager(manager *configmanager.ConfigManager) *Service {
	api.manager = manager
	return api
}

// WithHook sets a custom hook for request/response interception.
// Returns the service instance for method chaining.
func (api *Service) WithHook(hook hooks.Hook) *Service {
	api.hook = hook
	return api
}

func (api *Service) config() *leddisplayapiconfig.Config {
	return api.manager.GetPredefinedIcons()
}

func (api *Service) getHook() hooks.Hook {
	return api.hook
}

func (api *Service) SetBaseURL(baseURL string) {
	config := api.config()
	config.SetBaseURL(baseURL)
}

func (api *Service) SetTimeout(timeout time.Duration) {
	config := api.config()
	config.SetTimeout(timeout)
}

func (api *Service) SetAPIKey(apiKey string) {
	config := api.config()
	config.SetAPIKey(apiKey)
}

// SetGetPredefinedIconsConfig sets method-level configuration for GetPredefinedIcons.
// Options are applied to every future call to GetPredefinedIcons and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetGetPredefinedIconsConfig(opts ...leddisplayapiconfig.RequestOption) *Service {
	api.getPredefinedIconsConfig = opts
	return api
}

// Returns a list of all available icon codes that can be used in display text. Icons are returned as simple string codes in the `:icon_name:` format.
func (api *Service) GetPredefinedIcons(ctx context.Context, opts ...leddisplayapiconfig.RequestOption) (*GetPredefinedIconsOkResponse, error) {
	config := *api.config()
	for _, opt := range api.getPredefinedIconsConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/predefined-icons").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		Build()

	httpClient := restClient.NewRestClient[GetPredefinedIconsOkResponse, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewLedDisplayAPIError[[]byte](err)
	}

	return &resp.Data, nil
}
