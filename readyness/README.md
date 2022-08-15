# Readyness

This module provides a readyness http endpoint (`/ready`) that can be controlled by the use of a `readyness.ServiceReadyness`.

## Initialization

```go
app := fx.New(
	env.Module,
	logrus.Module,
	metrics.Module,
	fiber.Module,
	health.Module,
	func(service *readyness.ServiceReadyness) {
		// use service here
	},
)
app.Run()
```

## Service functions

- `IsReady` checks a list of components if they are ready. Returns false as soon as one of those components is not ready.
  If no component is specified, all available components will be checked.
- `SetReady` defines the ready status of a given component as "ready".
- `Register` registers a new component as not ready.
