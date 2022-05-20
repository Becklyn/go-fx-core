package web

import "github.com/gofiber/fiber/v2"

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
