package web

import "go.uber.org/fx"

var Module = fx.Provide(
	newFiberMiddlewareRegistry,
	newFiber,
)
