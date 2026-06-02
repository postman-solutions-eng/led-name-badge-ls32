package sharedmodels

import (
	"example.com/led-display-api/sdk/internal/clients/rest/httptransport"
	"net/http"
)

// LedDisplayAPIError wraps API errors with detailed metadata including status code, headers, and raw response.
// It implements the error interface and provides structured access to error information.
type LedDisplayAPIError[T any] struct {
	Err      error
	Data     *T
	Body     []byte
	Raw      *http.Response
	Metadata LedDisplayAPIErrorMetadata
}

// LedDisplayAPIErrorMetadata contains HTTP metadata associated with an error response.
type LedDisplayAPIErrorMetadata struct {
	Headers    map[string]string
	StatusCode int
}

// NewLedDisplayAPIError creates a new LedDisplayAPIError from an internal transport error.
// It extracts error details, body, status code, and headers into a user-facing error structure.
func NewLedDisplayAPIError[T any](transportError *httptransport.ErrorResponse[T]) *LedDisplayAPIError[T] {
	return &LedDisplayAPIError[T]{
		Err:  transportError.GetError(),
		Data: transportError.Data,
		Body: transportError.GetBody(),
		Raw:  transportError.Raw,
		Metadata: LedDisplayAPIErrorMetadata{
			StatusCode: transportError.GetStatusCode(),
			Headers:    transportError.GetHeaders(),
		},
	}
}

// Error implements the error interface, returning the error message string.
func (e *LedDisplayAPIError[T]) Error() string {
	if e == nil || e.Err == nil {
		return ""
	}
	return e.Err.Error()
}

// Unwrap returns the underlying error, enabling errors.Is and errors.As to traverse the chain.
func (e *LedDisplayAPIError[T]) Unwrap() error {
	return e.Err
}

// GetData returns the deserialized error response data.
// Returns nil if unmarshaling failed or the response body was empty.
func (e *LedDisplayAPIError[T]) GetData() *T {
	return e.Data
}

// GetBody returns the raw response body bytes from the error response.
// Returns nil if no response body was received.
func (e *LedDisplayAPIError[T]) GetBody() []byte {
	return e.Body
}
