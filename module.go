package main

import (
	"github.com/Becklyn/go-fx-core/cors"
	"github.com/Becklyn/go-fx-core/env"
	f "github.com/Becklyn/go-fx-core/fiber"
	"github.com/Becklyn/go-fx-core/health"
	"github.com/Becklyn/go-fx-core/logrus"
	"github.com/Becklyn/go-fx-core/metrics"
	"github.com/Becklyn/go-fx-core/middleware"
	"github.com/Becklyn/go-fx-core/readyness"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

var Module = fx.Options(
	env.Module,
	logrus.Module,
	cors.Module,
	metrics.Module,
	readyness.Module,
	health.Module,
	middleware.Module,
	f.Module,
)

func main() {
	fx.New(
		Module,
		fx.Invoke(func(app *fiber.App) {
			app.Get("/", func(c *fiber.Ctx) error {
				return c.SendString("Hello, World!")
			})
		}),
	).Run()
}
