package jaeger

import (
	"testing"

	"github.com/go-kit/kit/log"
	"github.com/stretchr/testify/assert"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

func TestConfig(t *testing.T) {
	config := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type: "const",
		},
	}
	opts := newOptions(Config(config))

	assert.Equal(t, config, opts.config)
}

func TestLogger(t *testing.T) {
	logger := log.NewNopLogger()
	opts := newOptions(Logger(logger))

	assert.Equal(t, logger, opts.logger)
}
