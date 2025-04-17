package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	HttpRequestsTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "http_request_total",
		Help: "Total number of HTTP requests",
	}, []string{"path", "method", "status_code"})

	HttpRequestsCurrent = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "http_request_current",
		Help: "Current number of HTTP requests",
	})

	HttpRequestsInflightMax = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "http_request_inflight_max",
		Help: "Maximum number of concurrent HTTP requests",
	})

	HttpRequestsDurationHistogram = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name: "http_request_duration_seconds",
		Help: "Duration of HTTP requests",
		Buckets: []float64{
			0.1,
			0.2,
			0.25,
			0.5,
			1,
		},
	}, []string{"path", "method"})

	HttpRequestsDurationSummary = promauto.NewSummaryVec(prometheus.SummaryOpts{
		Name: "http_request_duration_seconds_summary",
		Help: "Duration of HTTP requests",
		Objectives: map[float64]float64{
			0.5:   0.05,
			0.95:  0.01,
			0.99:  0.001,
			0.999: 0.0001,
		},
	}, []string{"path", "method"})
)
