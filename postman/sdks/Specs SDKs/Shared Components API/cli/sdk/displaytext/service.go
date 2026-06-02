package displaytext

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

// Service provides methods to interact with DisplayText-related API endpoints.
// It uses a configuration manager for settings and supports custom hooks for request/response interception.
type Service struct {
	manager                 *configmanager.ConfigManager
	hook                    hooks.Hook
	createDisplayTextConfig []leddisplayapiconfig.RequestOption
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
	return api.manager.GetDisplayText()
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

// SetCreateDisplayTextConfig sets method-level configuration for CreateDisplayText.
// Options are applied to every future call to CreateDisplayText and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetCreateDisplayTextConfig(opts ...leddisplayapiconfig.RequestOption) *Service {
	api.createDisplayTextConfig = opts
	return api
}

// Updates the text and visual content displayed on the LED name badge. Accepts text and icon codes in the format `:icon_name:`.
//
// <img src="https://content.pstmn.io/05b1ef2c-9fd3-4f0e-8b2f-76998fb3f6e5/aW1hZ2UucG5n" width="240">
//
// ## Supported Icons
//
// The following icon codes can be embedded in text using the `:icon_name:` syntax:
//
// - `:ball:` - Filled circle
//
// - `:happy:` - Simple smiley face
//
// - `:happy2:` - Larger smiley (2 columns wide)
//
// - `:heart:` - Outline heart
//
// - `:HEART:` - Filled heart
//
// - `:heart2:` - Larger outline heart (2 columns wide)
//
// - `:HEART2:` - Larger filled heart (2 columns wide)
//
// - `:fablab:` - FabLab logo
//
// - `:bicycle:` - Bicycle icon (3 columns wide)
//
// - `:bicycle_r:` - Bicycle facing right (3 columns wide)
//
// - `:owncloud:` - OwnCloud logo (3 columns wide)
//
// - `:octocat:` - GitHub Octocat
//
// - `:smile:` - Smile emoji
//
// - `:star:` - Star icon
//
// - `:sun:` - Sun icon
//
// ## Character Restrictions
//
// **Supported characters:**
//
// - Letters: A-Z, a-z
//
// - Numbers: 0-9
//
// - Special characters: `^ !"$%&/()=?` °}\]\[{@ \~ |<>,;.:-_#'+_\`
//
// - German umlauts: äöüÄÖÜß
//
// - French/European accents: àäòöùüèéêëôöûîïÿçÀÅÄÉÈÊËÖÔÜÛÙŸŠ
//
// **NOT supported:**
//
// - Unicode emoji (e.g., 🌍) - will return a 400 error
func (api *Service) CreateDisplayText(ctx context.Context, createDisplayTextRequest CreateDisplayTextRequest, opts ...leddisplayapiconfig.RequestOption) (*CreateDisplayTextOkResponse, error) {
	config := *api.config()
	for _, opt := range api.createDisplayTextConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/display-text").
		WithConfig(config).
		WithBody(createDisplayTextRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		Build()

	httpClient := restClient.NewRestClient[CreateDisplayTextOkResponse, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewLedDisplayAPIError[[]byte](err)
	}

	return &resp.Data, nil
}
