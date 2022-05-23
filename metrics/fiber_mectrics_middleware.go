package metrics

import (
	"strconv"
	"strings"
	"time"

	"github.com/Becklyn/go-fx-core/web"
	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"go.uber.org/fx"
)

var currentRequests = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "fiber_requests_current",
		Help: "The current number of active requests",
	},
	[]string{"method", "path"},
)

var totalRequests = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Name: "fiber_requests_total",
		Help: "Total number of requests processed by fiber",
	},
	[]string{"status", "method", "path"},
)

var requestDuration = promauto.NewHistogramVec(
	prometheus.HistogramOpts{
		Name: "fiber_request_duration_seconds",
		Help: "Duration of requests processed by fiber",
		Buckets: []float64{
			0.1, // Instant response
			0.25,
			0.5,
			0.75,
			1, // Maximum acceptable limit
		},
	},
	[]string{"status", "method", "path"},
)

type FiberMetricsMiddleware struct {
	fx.In
}

func useFiberMetricsMiddleware(
	registry *web.FiberMiddlewareRegistry,
	middleware FiberMetricsMiddleware,
) {
	registry.Use(&web.FiberMiddleware{
		Name:    "Fiber metrics",
		Handler: middleware.Handler,
	})
}

func (m *FiberMetricsMiddleware) Handler(ctx *fiber.Ctx) error {
	start := time.Now()
	method := string(ctx.Context().Method())
	path := string(ctx.Context().Path())

	if strings.Contains(path, "metrics") || strings.Contains(path, "health") {
		return ctx.Next()
	}

	currentRequests.WithLabelValues(
		method,
		path,
	).Inc()
	defer currentRequests.WithLabelValues(
		method,
		path,
	).Dec()

	err := ctx.Next()
	status := strconv.Itoa(ctx.Response().StatusCode())
	if err, ok := err.(*fiber.Error); err != nil && ok {
		status = strconv.Itoa(err.Code)
	}

	totalRequests.WithLabelValues(
		status,
		method,
		path,
	).Inc()

	duration := time.Since(start).Seconds()
	requestDuration.WithLabelValues(
		status,
		method,
		path,
	).Observe(duration)

	return nil
}
