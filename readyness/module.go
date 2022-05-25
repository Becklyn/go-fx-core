package readyness

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		NewServiceReadyness,
	),
	fx.Invoke(
		useReadynessEndpoint,
	),
)
