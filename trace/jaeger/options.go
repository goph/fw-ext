package jaeger

import (
	"github.com/go-kit/kit/log"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

// options holds a list of options used during the error handler construction.
type options struct {
	config jaegercfg.Configuration
	logger log.Logger
}

// newOptions creates a new options instance,
// applies the provided option list and falls back to defaults where necessary.
func newOptions(opts ...TracerOption) *options {
	o := new(options)

	for _, opt := range opts {
		opt(o)
	}

	return o
}

// TracerOption sets options in the tracer.
type TracerOption func(o *options)

// Config sets a configuration in the tracer.
func Config(c jaegercfg.Configuration) TracerOption {
	return func(o *options) {
		o.config = c
	}
}

// Logger sets a logger in the tracer.
func Logger(l log.Logger) TracerOption {
	return func(o *options) {
		o.logger = l
	}
}
