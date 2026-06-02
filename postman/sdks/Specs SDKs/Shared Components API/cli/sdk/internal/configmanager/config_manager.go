package configmanager

import (
	"example.com/led-display-api/sdk/leddisplayapiconfig"
	"time"
)

// ConfigManager manages configuration across all services with synchronized updates.
// Provides centralized configuration management and OAuth token handling for multiple services.
type ConfigManager struct {
	displayText     leddisplayapiconfig.Config
	predefinedIcons leddisplayapiconfig.Config
	displaySummary  leddisplayapiconfig.Config
}

// NewConfigManager creates a new configuration manager with the provided config and optional OAuth token service.
// Initializes service-specific configs and sets up OAuth token management if enabled.
func NewConfigManager(config leddisplayapiconfig.Config) *ConfigManager {
	return &ConfigManager{
		displayText:     config,
		predefinedIcons: config,
		displaySummary:  config,
	}
}

// SetBaseURL updates the BaseURL configuration parameter across all services.
// Changes are applied synchronously to all registered service configurations.
func (c *ConfigManager) SetBaseURL(baseURL string) {
	c.displayText.SetBaseURL(baseURL)
	c.predefinedIcons.SetBaseURL(baseURL)
	c.displaySummary.SetBaseURL(baseURL)
}

// SetTimeout updates the Timeout configuration parameter across all services.
// Changes are applied synchronously to all registered service configurations.
func (c *ConfigManager) SetTimeout(timeout time.Duration) {
	c.displayText.SetTimeout(timeout)
	c.predefinedIcons.SetTimeout(timeout)
	c.displaySummary.SetTimeout(timeout)
}

// SetAPIKey updates the APIKey configuration parameter across all services.
// Changes are applied synchronously to all registered service configurations.
func (c *ConfigManager) SetAPIKey(apiKey string) {
	c.displayText.SetAPIKey(apiKey)
	c.predefinedIcons.SetAPIKey(apiKey)
	c.displaySummary.SetAPIKey(apiKey)
}

// SetRetryConfig updates the retry configuration across all services.
// Changes are applied synchronously to all registered service configurations.
func (c *ConfigManager) SetRetryConfig(retry leddisplayapiconfig.RetryConfig) {
	c.displayText.SetRetryConfig(retry)
	c.predefinedIcons.SetRetryConfig(retry)
	c.displaySummary.SetRetryConfig(retry)
}

// GetDisplayText returns the configuration for the DisplayText service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetDisplayText() *leddisplayapiconfig.Config {
	return &c.displayText
}

// GetPredefinedIcons returns the configuration for the PredefinedIcons service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetPredefinedIcons() *leddisplayapiconfig.Config {
	return &c.predefinedIcons
}

// GetDisplaySummary returns the configuration for the DisplaySummary service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetDisplaySummary() *leddisplayapiconfig.Config {
	return &c.displaySummary
}

// GetBaseURL returns the currently configured base URL.
// All services share the same base URL; this reads it from the first service's config.
func (c *ConfigManager) GetBaseURL() string {
	return c.displayText.BaseURL
}
