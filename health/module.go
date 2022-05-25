package health

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		newServiceHealth,
	),
	fx.Invoke(
		useHealthEndpoint,
	),
)
