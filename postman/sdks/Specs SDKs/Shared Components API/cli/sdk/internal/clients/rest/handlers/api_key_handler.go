package handlers

import (
	"errors"

	"example.com/led-display-api/sdk/internal/clients/rest/httptransport"
)

// APIKeyHandler injects an API key into the request header for authentication.
// It adds the configured API key to the request if present in the config.
// T is the response type, E is the error type.
type APIKeyHandler[T any, E any] struct {
	nextHandler Handler[T, E]
}

// NewAPIKeyHandler creates a new API key authentication handler.
// Returns a handler that will inject the API key header when processing requests.
func NewAPIKeyHandler[T any, E any]() *APIKeyHandler[T, E] {
	return &APIKeyHandler[T, E]{
		nextHandler: nil,
	}
}

// Handle processes a regular request by adding the API key header if configured.
// Clones the request, adds the API key, and passes it to the next handler.
func (h *APIKeyHandler[T, E]) Handle(request httptransport.Request) (*httptransport.Response[T], *httptransport.ErrorResponse[E]) {
	if h.nextHandler == nil {
		err := errors.New("Handler chain terminated without terminating handler")
		return nil, httptransport.NewErrorResponse[E](err, nil)
	}

	nextRequest := request.Clone()

	if request.Config.APIKey == nil {
		return h.nextHandler.Handle(nextRequest)
	}

	nextRequest.SetHeader("X-API-Key", *request.Config.APIKey)

	return h.nextHandler.Handle(nextRequest)
}

// HandleStream processes a streaming request by adding the API key header if configured.
// Clones the request, adds the API key, and passes it to the next handler.
func (h *APIKeyHandler[T, E]) HandleStream(request httptransport.Request) (*httptransport.Stream[T], *httptransport.ErrorResponse[E]) {
	if h.nextHandler == nil {
		err := errors.New("Handler chain terminated without terminating handler")
		return nil, httptransport.NewErrorResponse[E](err, nil)
	}

	nextRequest := request.Clone()

	if request.Config.APIKey == nil {
		return h.nextHandler.HandleStream(nextRequest)
	}

	nextRequest.SetHeader("X-API-Key", *request.Config.APIKey)

	return h.nextHandler.HandleStream(nextRequest)
}

// SetNext sets the next handler in the chain.
// This method is called during chain construction to link handlers together.
func (h *APIKeyHandler[T, E]) SetNext(handler Handler[T, E]) {
	h.nextHandler = handler
}
