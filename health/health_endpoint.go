package health

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type HealthEndpoint struct {
	fx.In

	Health *ServiceHealth
}

func useHealthEndpoint(app *fiber.App, endpoint HealthEndpoint) {
	app.Get("/health", endpoint.Handle)
}

func (e *HealthEndpoint) Handle(c *fiber.Ctx) error {
	if healthy, reason := e.Health.IsHealthy(); !healthy {
		return c.Status(fiber.StatusServiceUnavailable).
			SendString(fmt.Sprintf("Service unavailable: %s", reason))
	}
	return c.Status(fiber.StatusOK).SendString("Service healthy")
}
