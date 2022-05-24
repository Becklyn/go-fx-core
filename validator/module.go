package validator

import "go.uber.org/fx"

var Module = fx.Provide(
	newValidator,
)
