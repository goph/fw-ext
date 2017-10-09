package debug

import (
	"net/http"

	"github.com/go-kit/kit/log"
)

// options holds a list of options used during the error handler construction.
type options struct {
	debug   bool
	handler *http.ServeMux
	logger  log.Logger
}

// newOptions creates a new options instance,
// applies the provided option list and falls back to defaults where necessary.
func newOptions(opts ...DebugOption) *options {
	o := new(options)

	for _, opt := range opts {
		opt(o)
	}

	// Default handler
	if o.handler == nil {
		o.handler = http.NewServeMux()
	}

	return o
}

// DebugOption sets options in the debug server.
type DebugOption func(o *options)

// Debug allows debug mode.
func Debug(d bool) DebugOption {
	return func(o *options) {
		o.debug = d
	}
}

// Handler sets a handler in the debug server.
func Handler(h *http.ServeMux) DebugOption {
	return func(o *options) {
		o.handler = h
	}
}

// Logger sets a logger in the debug server.
func Logger(l log.Logger) DebugOption {
	return func(o *options) {
		o.logger = l
	}
}
