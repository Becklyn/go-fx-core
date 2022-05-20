package core

import (
	"github.com/Becklyn/go-fx-core/env"
	"github.com/Becklyn/go-fx-core/health"
	"github.com/Becklyn/go-fx-core/metrics"
	"github.com/Becklyn/go-fx-core/web"
	"go.uber.org/fx"
)

var Module = fx.Options(
	env.Module,
	LogrusModule,
	web.Module,
	health.Module,
	metrics.Module,
	GrpcModule,
	ValidatorModule,
)
