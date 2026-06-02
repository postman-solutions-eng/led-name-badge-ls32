package displaysummary

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

// Service provides methods to interact with DisplaySummary-related API endpoints.
// It uses a configuration manager for settings and supports custom hooks for request/response interception.
type Service struct {
	manager                    *configmanager.ConfigManager
	hook                       hooks.Hook
	createDisplaySummaryConfig []leddisplayapiconfig.RequestOption
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
	return api.manager.GetDisplaySummary()
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

// SetCreateDisplaySummaryConfig sets method-level configuration for CreateDisplaySummary.
// Options are applied to every future call to CreateDisplaySummary and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetCreateDisplaySummaryConfig(opts ...leddisplayapiconfig.RequestOption) *Service {
	api.createDisplaySummaryConfig = opts
	return api
}

// Displays a summary message on the LED badge.
//
// Without an API key (`X-API-Key` header or `apiKey` in the body), shows the built-in demo summary text. When a key is provided, the endpoint fetches services from the Postman API Catalog and renders a compact health summary on the badge.
//
// Note: The API does not validate the `type` parameter and will return success for any type value provided.
func (api *Service) CreateDisplaySummary(ctx context.Context, createDisplaySummaryRequest CreateDisplaySummaryRequest, opts ...leddisplayapiconfig.RequestOption) (*CreateDisplaySummaryOkResponse, error) {
	config := *api.config()
	for _, opt := range api.createDisplaySummaryConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/display-summary").
		WithConfig(config).
		WithBody(createDisplaySummaryRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		Build()

	httpClient := restClient.NewRestClient[CreateDisplaySummaryOkResponse, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewLedDisplayAPIError[[]byte](err)
	}

	return &resp.Data, nil
}
