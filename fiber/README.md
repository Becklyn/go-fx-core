# Fiber (webserver)

Setting up the preconfigured webserver:

```go
app := fx.New(
	env.Module,
	logrus.Module,
	metrics.Module,
	fiber.Module,
)
app.Run()
```

It will integrate logging, metrics and basic error handling into your webserver app.

## Error handling

```go
func useFiber(app *fiber.App) {
	app.Get("/path", func(c *fiber.Ctx) error {
		// whenever you return a `fiber.Error`:
		// - the status code will be set automatically
		// - the error will be logged
		// - and the error string wil be returned in the response body
		return fiber.NewFiberError(errors.New("your error"), fiber.StatusBadRequest)
	})
}
```
