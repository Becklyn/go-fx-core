package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type FiberMiddleware struct {
	Name    string
	Route   string
	Handler func(c *fiber.Ctx) error
}

type FiberMiddlewareRegistry struct {
	middleware []*FiberMiddleware

	logger *logrus.Logger
}

func newFiberMiddlewareRegistry() *FiberMiddlewareRegistry {
	return &FiberMiddlewareRegistry{}
}

func (r *FiberMiddlewareRegistry) Use(middleware *FiberMiddleware) {
	r.middleware = append(r.middleware, middleware)
}

func (r *FiberMiddlewareRegistry) Register(app *fiber.App) {
	for _, middleware := range r.middleware {
		if middleware.Route == "" {
			app.Use(middleware.Handler)
			r.logger.WithFields(logrus.Fields{
				"middleware": middleware.Name,
			}).Info("Registered middleware globally")
		} else {
			app.Use(middleware.Route, middleware.Handler)
			r.logger.WithFields(logrus.Fields{
				"middleware": middleware.Name,
				"route":      middleware.Route,
			}).Info("Registered middleware on route")
		}
	}
}
