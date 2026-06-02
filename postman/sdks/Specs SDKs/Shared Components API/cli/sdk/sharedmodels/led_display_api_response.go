package sharedmodels

import (
	"encoding/json"
	"example.com/led-display-api/sdk/internal/clients/rest/httptransport"
	"net/http"
)

// LedDisplayAPIResponse is the user-facing wrapper for API responses.
// It contains the deserialized data, raw HTTP response, and metadata like headers and status code.
type LedDisplayAPIResponse[T any] struct {
	Data     T
	Raw      *http.Response
	Metadata LedDisplayAPIResponseMetadata
}

// LedDisplayAPIResponseMetadata contains HTTP metadata from the API response.
// Includes status code and headers for inspection and debugging.
type LedDisplayAPIResponseMetadata struct {
	Headers    map[string]string
	StatusCode int
}

// NewLedDisplayAPIResponse creates a new response wrapper from an internal transport response.
// Extracts data and metadata into a user-facing structure.
func NewLedDisplayAPIResponse[T any](resp *httptransport.Response[T]) *LedDisplayAPIResponse[T] {
	return &LedDisplayAPIResponse[T]{
		Data: resp.Data,
		Raw:  resp.Raw,
		Metadata: LedDisplayAPIResponseMetadata{
			StatusCode: resp.StatusCode,
			Headers:    resp.Headers,
		},
	}
}

// GetData returns the deserialized response data.
func (r *LedDisplayAPIResponse[T]) GetData() T {
	return r.Data
}

// String returns a JSON representation of the response for debugging.
// Returns an error message if JSON marshaling fails.
func (r LedDisplayAPIResponse[T]) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: LedDisplayAPIResponse to string"
	}
	return string(jsonData)
}
