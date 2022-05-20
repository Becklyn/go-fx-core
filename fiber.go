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
	newFiberMiddlewareRegistry,
	newFiber,
)

type FiberMiddleware struct {
	Name    string
	Route   string
	Handler func(c *fiber.Ctx) error
}

type FiberMiddlewareRegistry []FiberMiddleware

func newFiberMiddlewareRegistry() *FiberMiddlewareRegistry {
	return &FiberMiddlewareRegistry{}
}

func (r *FiberMiddlewareRegistry) Use(middleware FiberMiddleware) {
	*r = append(*r, middleware)
}

func newFiber(
	lifecycle fx.Lifecycle,
	middlewareRegisty *FiberMiddlewareRegistry,
	logger *logrus.Logger,
) *fiber.App {
	addr := env.StringWithDefault("FIBER_ADDR", ":3000")

	app := fiber.New()

	app.Use(fiberlog.New(fiberlog.Config{
		Format: "${status} - ${latency} ${method} ${path}\n",
		Output: logger.Writer(),
	}))

	for _, middleware := range *middlewareRegisty {
		if middleware.Route == "" {
			app.Use(middleware.Handler)
			logger.Infof("Registered %s middleware globally", middleware.Name)
		} else {
			app.Use(middleware.Route, middleware.Handler)
			logger.Infof("Registered %s middleware on route %s", middleware.Name, middleware.Route)
		}
	}

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
