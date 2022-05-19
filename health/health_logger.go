package health

import (
	"context"

	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

type HealthLogger struct {
	fx.In

	Logrus *logrus.Logger
	Health *ServiceHealth
}

func useHealthLogger(lifecycle fx.Lifecycle, logger HealthLogger) {
	ctx, cancel := context.WithCancel(context.Background())

	lifecycle.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			cancel()
			return nil
		},
	})

	go logger.Execute(ctx)
}

func (l *HealthLogger) Execute(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case component := <-l.Health.componentChanged:
			if component.health.healthy {
				logrus.WithField("component", component.name).
					Info("Component is healthy (again)")
			} else {
				logrus.WithField("component", component.name).
					Error("Component became unhealthy")
			}
		}
	}
}
