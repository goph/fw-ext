package health

import (
	"net/http"

	"github.com/goph/fw"
	"github.com/goph/healthz"
)

// HealthCollector provides an application option which registers a health collector.
func HealthCollector(a *fw.Application) fw.ApplicationOption {
	var options []fw.ApplicationOption

	healthCollector := healthz.Collector{}

	var handler *http.ServeMux

	mux, ok := a.Get("debug_handler")
	if h, ok2 := mux.(*http.ServeMux); !ok || !ok2 {
		handler = http.NewServeMux()
		options = append(options, fw.Entry("debug_handler", handler))
	} else {
		handler = h
	}

	// Add health checks
	handler.Handle("/healthz", healthCollector.Handler(healthz.LivenessCheck))
	handler.Handle("/readiness", healthCollector.Handler(healthz.ReadinessCheck))

	options = append(options, fw.Entry("health_collector", healthCollector))

	return fw.Options(options...)
}

// ApplicationStatus provides an application option which registers status checking in the health checker.
func ApplicationStatus(a *fw.Application) fw.ApplicationOption {
	healthCollector := a.MustGet("health_collector").(healthz.Collector)

	status := healthz.NewStatusChecker(healthz.Healthy)
	healthCollector.RegisterChecker(healthz.ReadinessCheck, status)

	return fw.LifecycleHook(fw.Hook{
		PreStart: func() error {
			status.SetStatus(healthz.Healthy)

			return nil
		},
		PreShutdown: func() error {
			status.SetStatus(healthz.Unhealthy)

			return nil
		},
	})
}
