# Validator

Initialization and usage:

```go
app := fx.New(
	validation.Module,
	fx.Invoke(func(validator *validator.Validate) {
		// use validator here
	}),
)
app.Run()
```
