package metrics

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(
		newFiberMetricsMiddleware,
	),
	fx.Invoke(
		usePrometheusMetricsEndpoint,
	),
)
