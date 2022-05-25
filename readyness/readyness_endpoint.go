package readyness

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type ReadynessEndpoint struct {
	fx.In

	Readyness *ServiceReadyness
}

func useReadynessEndpoint(app *fiber.App, endpoint ReadynessEndpoint) {
	app.Get("/ready", endpoint.Handle)
}

func (e *ReadynessEndpoint) Handle(c *fiber.Ctx) error {
	if ready, component := e.Readyness.IsReady(); !ready {
		return c.Status(fiber.StatusServiceUnavailable).
			SendString(fmt.Sprintf("Service not ready: uninitialized component: %s", component))
	}
	return c.Status(fiber.StatusOK).SendString("Service ready")
}
