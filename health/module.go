package health

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		NewServiceHealth,
	),
	fx.Invoke(
		useHealthEndpoint,
		useHealthLogger,
	),
)
