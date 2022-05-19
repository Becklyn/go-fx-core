package health

import (
	"context"
	"fmt"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

type status struct {
	Healthy bool
	Reason  *string

	mux sync.Mutex
}

var serviceStatus = &status{
	Healthy: true,
}

var Module = fx.Invoke(
	useHealthEndpoint,
	useHealthLogger,
)

func useHealthEndpoint(app *fiber.App) {
	app.Get("/health", func(c *fiber.Ctx) error {
		if !serviceStatus.Healthy {
			return c.Status(503).SendString(
				fmt.Sprintf("Service unavailable: %s", *serviceStatus.Reason),
			)
		}
		return c.Status(200).SendString("Healthy")
	})
}

func useHealthLogger(lifecycle fx.Lifecycle, logrus *logrus.Logger) {
	ctx, cancel := context.WithCancel(context.Background())
	healthy := true

	lifecycle.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			cancel()
			return nil
		},
	})

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				if healthy != serviceStatus.Healthy {
					if healthy {
						logrus.Info("Healthy")
					} else {
						logrus.WithField("reason", *serviceStatus.Reason).
							Warn("Service unavailable")
					}
				}
				healthy = serviceStatus.Healthy
			}
		}
	}()
}

func IsHealthy() bool {
	return serviceStatus.Healthy
}

func StatusHealthy() {
	serviceStatus.mux.Lock()
	defer serviceStatus.mux.Unlock()

	serviceStatus.Healthy = true
	serviceStatus.Reason = nil
}

func StatusNotHealthy(reason string) {
	serviceStatus.mux.Lock()
	defer serviceStatus.mux.Unlock()

	serviceStatus.Healthy = false
	serviceStatus.Reason = &reason
}
