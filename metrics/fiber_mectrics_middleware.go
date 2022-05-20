package metrics

import (
	"strconv"

	core "github.com/Becklyn/go-fx-core"
	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/fx"
)

var currentRequests = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Name: "fiber_requests_current",
		Help: "The current number of active requests",
	},
)

var totalRequests = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "fiber_requests_total",
		Help: "Total number of requests processed by fiber",
	},
	[]string{"status", "method", "path"},
)

var requestDuration = prometheus.NewHistogram(
	prometheus.HistogramOpts{
		Name: "fiber_request_duration_seconds",
		Help: "Duration of requests processed by fiber",
		Buckets: []float64{
			0.000000001, // 1ns
			0.000000002,
			0.000000005,
			0.00000001, // 10ns
			0.00000002,
			0.00000005,
			0.0000001, // 100ns
			0.0000002,
			0.0000005,
			0.000001, // 1µs
			0.000002,
			0.000005,
			0.00001, // 10µs
			0.00002,
			0.00005,
			0.0001, // 100µs
			0.0002,
			0.0005,
			0.001, // 1ms
			0.002,
			0.005,
			0.01, // 10ms
			0.02,
			0.05,
			0.1, // 100 ms
			0.2,
			0.5,
			1.0, // 1s
			2.0,
			5.0,
			10.0, // 10s
			15.0,
			20.0,
			30.0,
		},
	},
)

type FiberMetricsMiddleware struct {
	fx.In
}

func useFiberMetricsMiddleware(
	registry *core.FiberMiddlewareRegistry,
	middleware FiberMetricsMiddleware,
) {
	registry.Use(core.FiberMiddleware{
		Name:    "Fiber metrics",
		Handler: middleware.Handler,
	})
}

func (m *FiberMetricsMiddleware) Handler(ctx *fiber.Ctx) error {
	requestDurationTimer := prometheus.NewTimer(requestDuration)
	defer requestDurationTimer.ObserveDuration()

	currentRequests.Inc()
	defer currentRequests.Dec()

	if err := ctx.Next(); err != nil {
		return err
	}

	status := strconv.Itoa(ctx.Response().StatusCode())
	method := ctx.Route().Method
	path := ctx.Route().Path

	totalRequests.WithLabelValues(
		status,
		method,
		path,
	).Inc()

	return nil
}
