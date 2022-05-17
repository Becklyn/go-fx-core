package core

import (
	"github.com/Becklyn/go-fx-core/env"
	"go.uber.org/fx"
)

var Module = fx.Options(
	env.Module,
	LogrusModule,
	FiberModule,
	GrpcModule,
	ValidatorModule,
)
