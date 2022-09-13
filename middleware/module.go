package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	NewMiddlewareRegistry,
	NewMiddleware,
)

type Middleware interface {
	Use(func(c *fiber.Ctx) error)
}

type MiddlewareRegistry struct {
	handlers []func(c *fiber.Ctx) error
}

func NewMiddlewareRegistry() *MiddlewareRegistry {
	return &MiddlewareRegistry{
		handlers: []func(c *fiber.Ctx) error{},
	}
}

func NewMiddleware(m *MiddlewareRegistry) Middleware {
	return m
}

func (r *MiddlewareRegistry) Use(handler func(c *fiber.Ctx) error) {
	r.handlers = append(r.handlers, handler)
}

func (r *MiddlewareRegistry) Handle(c *fiber.Ctx) error {
	errs := 0
	for _, handler := range r.handlers {
		if err := handler(c); err != nil {
			fiberErr, ok := err.(*fiber.Error)
			if ok && fiberErr.Code == fiber.StatusNotFound &&
				fiberErr.Message == fmt.Sprintf("Cannot %s %s", c.Method(), c.OriginalURL()) {
				errs++
				continue
			} else {
				return err
			}
		}
	}
	if len(r.handlers) > 0 && errs == 0 {
		return nil
	}
	return c.Next()
}
