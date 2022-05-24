package grpc

import "go.uber.org/fx"

var Module = fx.Provide(
	newGrpc,
)
