package core

import (
	"github.com/Becklyn/go-fx-core/env"
	"github.com/Becklyn/go-fx-core/health"
	"go.uber.org/fx"
)

var Module = fx.Options(
	env.Module,
	LogrusModule,
	FiberModule,
	health.Module,
	GrpcModule,
	ValidatorModule,
)
