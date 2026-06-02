package leddisplayapiconfig

// Environment type defines the available API environments.
type Environment string

// Environment constants define the available base URLs for different deployment environments.
// Use these constants when configuring the SDK client.
const (
	DefaultEnvironment Environment = "http://127.0.0.1:5001"
)
