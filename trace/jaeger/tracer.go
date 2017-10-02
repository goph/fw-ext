package jaeger

import (
	"fmt"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/opentracing/opentracing-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

func NewTracer(serviceName string, opts ...TracerOption) opentracing.Tracer {
	o := newOptions(opts...)

	var jaegerOptions []jaegercfg.Option

	if o.logger != nil {
		jaegerOptions = append(jaegerOptions, jaegercfg.Logger(&kitLogger{o.logger}))
	}

	tracer, _, err := o.config.New(serviceName, jaegerOptions...)
	if err != nil {
		panic(err)
	}

	return tracer
}

// kitLogger wraps the application logger instance in a Jaeger compatible one.
type kitLogger struct {
	logger log.Logger
}

// Error implements the github.com/uber/jaeger-client-go/log.Logger interface.
func (l *kitLogger) Error(msg string) {
	level.Error(l.logger).Log("msg", msg)
}

// Infof implements the github.com/uber/jaeger-client-go/log.Logger interface.
func (l *kitLogger) Infof(msg string, args ...interface{}) {
	level.Info(l.logger).Log("msg", fmt.Sprintf(msg, args...))
}
