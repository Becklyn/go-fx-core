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
	app.Get("/path", func(ctx *fiber.Ctx) error {
		// whenever you return a `fiber.Error`:
		// - the status code will be set automatically
		// - the error will be logged
		// - and the error string wil be returned in the response body
		return fiber.NewFiberError(errors.New("your error"), fiber.StatusBadRequest)
	})
}
```

## Middleware usage

If you want to add middleware to your fiber app, you can do this like in the following:

```go
func newMiddlewareMap() *MiddlewareHandlerMap {
	return &MiddlewareHandlerMap{
		// fiber supports middleware restriction to paths. Use an empty string for global middleware.
		// all middleware will be initialized in the same order that it is defined.
		"/path": []fiber.Handler{
			func(ctx *fiber.Ctx) error {
				// your middleware code here
				return ctx.Next()
			},
		},
	}
}
```

### Predefined middleware

- `middleware.NoCors()` middleware that disables CORS for development environments
