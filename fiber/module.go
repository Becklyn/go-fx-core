package fiber

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(
		newFiber,
	),
	fx.Invoke(
		useFiber,
	),
)
