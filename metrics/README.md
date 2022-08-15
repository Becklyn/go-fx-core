# Metrics (prometheus)

The metrics module is designed to be used in combination with the `fiber.Module`:

```go
app := fx.New(
	env.Module,
	logrus.Module,
	metrics.Module,
	fiber.Module,
)
app.Run()
```

It wil provide a `/metrics` http route that provides metrics for use with prometeus.
