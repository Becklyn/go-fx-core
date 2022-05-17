package core

import (
	"context"

	"github.com/Becklyn/go-fx-core/env"
	"github.com/gofiber/fiber/v2"
	fiberlog "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

var FiberModule = fx.Provide(
	newFiber,
)

func newFiber(lifecycle fx.Lifecycle, logger *logrus.Logger) *fiber.App {
	addr := env.StringWithDefault("FIBER_ADDR", ":3000")

	app := fiber.New()

	app.Use(fiberlog.New(fiberlog.Config{
		Format: "${status} - ${latency} ${method} ${path}\n",
		Output: logger.Writer(),
	}))

	app.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(fiber.StatusOK)
	})

	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := app.Listen(addr); err != nil {
					logger.Fatal(err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			if err := app.Shutdown(); err != nil {
				logger.Fatal(err)
			}
			return nil
		},
	})

	return app
}
