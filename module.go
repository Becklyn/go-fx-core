package core

import (
	"github.com/Becklyn/go-fx-core/v2/env"
	"github.com/Becklyn/go-fx-core/v2/fiber"
	"github.com/Becklyn/go-fx-core/v2/health"
	"github.com/Becklyn/go-fx-core/v2/logrus"
	"github.com/Becklyn/go-fx-core/v2/metrics"
	"github.com/Becklyn/go-fx-core/v2/readyness"
	"go.uber.org/fx"
)

var Module = fx.Options(
	env.Module,
	logrus.Module,
	metrics.Module,
	readyness.Module,
	health.Module,
	fiber.Module,
)
