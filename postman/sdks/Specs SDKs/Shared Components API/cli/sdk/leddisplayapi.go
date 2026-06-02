package leddisplayapi

import (
	"example.com/led-display-api/sdk/displaysummary"
	"example.com/led-display-api/sdk/displaytext"
	"example.com/led-display-api/sdk/internal/clients/rest/hooks"
	"example.com/led-display-api/sdk/internal/configmanager"
	"example.com/led-display-api/sdk/predefinedicons"
	"time"
)

// LedDisplayAPI is the main SDK client that provides access to all service endpoints.
// It manages configuration, authentication, and service instances with centralized settings.
type LedDisplayAPI struct {
	DisplayText     *displaytext.Service
	PredefinedIcons *predefinedicons.Service
	DisplaySummary  *displaysummary.Service
	manager         *configmanager.ConfigManager
}

func NewLedDisplayAPI(config Config) *LedDisplayAPI {
	displayText := displaytext.NewService()
	predefinedIcons := predefinedicons.NewService()
	displaySummary := displaysummary.NewService()

	manager := configmanager.NewConfigManager(config)
	hook := hooks.NewDefaultHook()
	displayText.WithConfigManager(manager)
	predefinedIcons.WithConfigManager(manager)
	displaySummary.WithConfigManager(manager)
	displayText.WithHook(hook)
	predefinedIcons.WithHook(hook)
	displaySummary.WithHook(hook)

	return &LedDisplayAPI{
		DisplayText:     displayText,
		PredefinedIcons: predefinedIcons,
		DisplaySummary:  displaySummary,
		manager:         manager,
	}
}

func (l *LedDisplayAPI) SetBaseURL(baseURL string) {
	l.manager.SetBaseURL(baseURL)
}

func (l *LedDisplayAPI) SetTimeout(timeout time.Duration) {
	l.manager.SetTimeout(timeout)
}

func (l *LedDisplayAPI) SetAPIKey(apiKey string) {
	l.manager.SetAPIKey(apiKey)
}

// SetEnvironment configures the SDK to use the specified environment's base URL.
func (l *LedDisplayAPI) SetEnvironment(environment Environment) {
	l.manager.SetBaseURL(string(environment))
}

// c029837e0e474b76bc487506e8799df5e3335891efe4fb02bda7a1441840310c
