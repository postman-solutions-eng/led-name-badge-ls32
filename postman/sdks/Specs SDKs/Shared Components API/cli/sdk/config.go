package leddisplayapi

import (
	"example.com/led-display-api/sdk/leddisplayapiconfig"
	"example.com/led-display-api/sdk/param"
)

// The type aliases below let consumers use a single import path for the entire SDK.
// Internally the concrete types live in leddisplayapiconfig and param.

// Config holds all configuration parameters for the SDK client.
type Config = leddisplayapiconfig.Config

// RequestOption is a function that configures a single request.
type RequestOption = leddisplayapiconfig.RequestOption

// Environment defines the available API base URLs.
type Environment = leddisplayapiconfig.Environment

// RetryConfig holds all runtime-configurable retry parameters.
type RetryConfig = leddisplayapiconfig.RetryConfig

// NewConfig creates a Config with spec-derived defaults.
var NewConfig = leddisplayapiconfig.NewConfig

// NewRetryConfig returns a RetryConfig initialized with spec-derived defaults.
var NewRetryConfig = leddisplayapiconfig.NewRetryConfig

// WithBaseURL returns a RequestOption that overrides BaseURL for a single request.
var WithBaseURL = leddisplayapiconfig.WithBaseURL

// WithTimeout returns a RequestOption that overrides Timeout for a single request.
var WithTimeout = leddisplayapiconfig.WithTimeout

// WithAPIKey returns a RequestOption that overrides APIKey for a single request.
var WithAPIKey = leddisplayapiconfig.WithAPIKey

// WithRetryConfig returns a RequestOption that overrides the RetryConfig for a single request.
var WithRetryConfig = leddisplayapiconfig.WithRetryConfig

// Nullable returns a *param.Nullable[T] set to v — use for nullable fields with a value.
func Nullable[T any](v T) *param.Nullable[T] { return &param.Nullable[T]{Value: v} }

// Null returns a *param.Nullable[T] with IsNull set to true, signalling an explicit JSON null.
func Null[T any]() *param.Nullable[T] { return param.Null[T]() }

// Ptr returns a pointer to v — use when no type-specific helper exists.
func Ptr[T any](v T) *T { return param.Ptr(v) }

// Environment constants for the available API base URLs.
const (
	DefaultEnvironment Environment = leddisplayapiconfig.DefaultEnvironment
)
