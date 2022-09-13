package core

import (
	"github.com/Becklyn/go-fx-core/env"
	"github.com/Becklyn/go-fx-core/fiber"
	"github.com/Becklyn/go-fx-core/health"
	"github.com/Becklyn/go-fx-core/logrus"
	"github.com/Becklyn/go-fx-core/metrics"
	"github.com/Becklyn/go-fx-core/middleware"
	"github.com/Becklyn/go-fx-core/readyness"
	"go.uber.org/fx"
)

var Module = fx.Options(
	env.Module,
	logrus.Module,
	metrics.Module,
	readyness.Module,
	health.Module,
	middleware.Module,
	fiber.Module,
)
