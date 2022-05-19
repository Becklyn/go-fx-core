package health

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type status struct {
	Healthy bool
	Reason  *string
}

var serviceStatus = status{
	Healthy: true,
}

var Module = fx.Invoke(
	useHealthEndpoint,
)

func useHealthEndpoint(app *fiber.App) {
	app.Get("/health", func(c *fiber.Ctx) error {
		if !serviceStatus.Healthy {
			return c.Status(503).SendString(
				fmt.Sprintf("Not healthy: %s", *serviceStatus.Reason),
			)
		}
		return c.Status(200).SendString("Healthy")
	})
}

func GetStatus() status {
	return serviceStatus
}

func IsHealthy() bool {
	return serviceStatus.Healthy
}

func StatusHealthy() {
	serviceStatus = status{
		Healthy: true,
		Reason:  nil,
	}
}

func StatusNotHealthy(reason string) {
	serviceStatus = status{
		Healthy: false,
		Reason:  &reason,
	}
}
