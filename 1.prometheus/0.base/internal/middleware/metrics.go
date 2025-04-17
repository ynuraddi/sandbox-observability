package middleware

import (
	"net/http"
	"prom/internal/helpers"
	"prom/internal/metrics"
	"time"

	"github.com/go-chi/chi/v5"
)

func HttpMetrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sw := helpers.NewStatusResponseWriter(w)
		now := time.Now()

		metrics.HttpRequestsCurrent.Inc()
		defer metrics.HttpRequestsCurrent.Dec()

		next.ServeHTTP(sw, r)

		elapsedSends := time.Since(now).Seconds()
		path := chi.RouteContext(r.Context()).RoutePattern()
		method := chi.RouteContext(r.Context()).RouteMethod
		status := sw.GetStatusString()

		metrics.HttpRequestsTotal.WithLabelValues(path, method, status).Inc()
		metrics.HttpRequestsDurationHistogram.WithLabelValues(path, method).Observe(elapsedSends)
		metrics.HttpRequestsDurationSummary.WithLabelValues(path, method).Observe(elapsedSends)
	})
}
