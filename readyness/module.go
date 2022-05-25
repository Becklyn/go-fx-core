package readyness

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		newServiceReadyness,
	),
	fx.Invoke(
		useReadynessEndpoint,
	),
)
