package web

import (
	"context"

	"github.com/Becklyn/go-fx-core/env"
	"github.com/gofiber/fiber/v2"
	fiberlog "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

func newFiber(
	logger *logrus.Logger,
	registry *FiberMiddlewareRegistry,
) *fiber.App {
	app := fiber.New()

	app.Use(fiberlog.New(fiberlog.Config{
		Format: "${status} - ${latency} ${method} ${path}\n",
		Output: logger.Writer(),
	}))

	registry.Register(app)

	return app
}

func useFiber(
	lifecycle fx.Lifecycle,
	app *fiber.App,
	logger *logrus.Logger,
) {
	addr := env.StringWithDefault("FIBER_ADDR", ":3000")

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
}
