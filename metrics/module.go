package metrics

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Invoke(
		usePrometheusMetricsEndpoint,
		useFiberMetricsMiddleware,
	),
)
