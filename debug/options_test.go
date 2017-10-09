package debug

import (
	"testing"

	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/stretchr/testify/assert"
)

func TestDebug(t *testing.T) {
	opts := newOptions(Debug(true))

	assert.Equal(t, true, opts.debug)
}

func TestHandler(t *testing.T) {
	handler := http.NewServeMux()
	handler.Handle("/", http.NotFoundHandler())
	opts := newOptions(Handler(handler))

	assert.Equal(t, handler, opts.handler)
}

func TestLogger(t *testing.T) {
	logger := log.NewNopLogger()
	opts := newOptions(Logger(logger))

	assert.Equal(t, logger, opts.logger)
}
