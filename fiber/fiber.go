package fiber

import (
	"context"

	"github.com/Becklyn/go-fx-core/v2/env"
	"github.com/Becklyn/go-fx-core/v2/metrics"
	"github.com/gofiber/fiber/v2"
	fiberlog "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

type MiddlewareHandlerMap map[string][]fiber.Handler

type FiberParams struct {
	fx.In

	Logger     *logrus.Logger
	Middleware *MiddlewareHandlerMap `optional:"true"`
}

func newFiber(params FiberParams) *fiber.App {
	app := fiber.New()

	app.Use(fiberlog.New(fiberlog.Config{
		Format: "${latency} - ${status} ${method} ${path}\n",
		Output: params.Logger.Writer(),
	}))

	app.Use(errorMiddleware(params.Logger))
	app.Use(metrics.MetricsMiddleware())

	if params.Logger != nil {
		for path, handlers := range *params.Middleware {
			for _, handler := range handlers {
				if handler == nil {
					continue
				}

				if path == "" {
					app.Use(handler)
				} else {
					app.Use(path, handler)
				}
			}
		}
	}

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
