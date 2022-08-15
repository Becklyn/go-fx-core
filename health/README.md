# Health

This module provides a health http endpoint (`/health`) that can be controlled by the use of a `health.Service`.

## Initialization

```go
app := fx.New(
	env.Module,
	logrus.Module,
	metrics.Module,
	fiber.Module,
	health.Module,
	func(service *health.ServiceHealth) {
		// use service here
	},
)
app.Run()
```

## Service functions

- `IsHealthy` checks a list of components if they are healthy. Returns false as soon as one of those components is unhealthy.
  If no component is specified, all available components will be checked.
- `SetHealthy` defines the health status of a given component as "healthy".
- `SetUnhealthy` defines the health status of a given component as "not healthy" and requires a reason.
