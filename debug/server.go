package debug

import (
	"context"
	stdlog "log"
	"net"
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/goph/fw"
	"github.com/goph/serverz"
	"github.com/goph/stdlib/expvar"
	"github.com/goph/stdlib/net/http/pprof"
	"github.com/goph/stdlib/x/net/trace"
)

// NewServer creates a new debug server.
func NewServer(addr net.Addr, opts ...DebugOption) serverz.Server {
	o := newOptions(opts...)

	if o.debug {
		// This is probably okay, as this service should not be exposed to public in the first place.
		trace.SetAuth(trace.NoAuth)

		expvar.RegisterRoutes(o.handler)
		pprof.RegisterRoutes(o.handler)
		trace.RegisterRoutes(o.handler)
	}

	return &serverz.AppServer{
		Server: &http.Server{
			Handler:  o.handler,
			ErrorLog: stdlog.New(log.NewStdlibAdapter(level.Error(log.With(o.logger, "server", "debug"))), "", 0),
		},
		Name:   "debug",
		Addr:   addr,
		Logger: o.logger,
	}
}

// DebugServer creates a new debug server option.
func DebugServer(addr net.Addr) fw.ApplicationOption {
	return fw.OptionFunc(func(a *fw.Application) fw.ApplicationOption {
		var handler *http.ServeMux

		mux, ok := a.Get("debug_handler")
		if h, ok2 := mux.(*http.ServeMux); !ok || !ok2 {
			handler = http.NewServeMux()
		} else {
			handler = h
		}

		server := NewServer(
			addr,
			Logger(a.Logger()),
			Handler(handler),
		)

		return fw.LifecycleHook(fw.Hook{
			OnStart: func(ctx context.Context, done chan<- interface{}) error {
				lis, err := net.Listen(addr.Network(), addr.String())
				if err != nil {
					return err
				}

				go func() {
					done <- server.Serve(lis)
				}()

				return nil
			},
			OnShutdown: func(ctx context.Context) error {
				return server.Shutdown(ctx)
			},
		})
	})
}
