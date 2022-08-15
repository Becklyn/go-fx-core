# env

This package tries to locate a `.env` or `.env.local` file in your app directory (or the parent directory). `.env.local` is prefered over `.env` and variables that are already existing will not be overridden.

## Initialization

```go
app := fx.New(
	env.Module,
)
app.Run()
```

## Functions

```go
env.GetEnvironment()
```

Will return the `APP_ENV` variables string value.

```go
env.IsDevelopment()
```

Will return true if the `APP_ENV` variable is set to `"dev"`. Useful to check if you are in development mode.

```go
env.String("YOUR_ENV_VAR_NAME")
```

Returns the `string` value of your env variable.

```go
env.StringWithDefault("YOUR_ENV_VAR_NAME", "your default")
```

Returns the `string` value of your env variable. If empty it will return the default string (in this case `"your default"`).

```go
env.Int("YOUR_ENV_VAR_NAME")
```

Returns the `int` value of your env variable.

```go
env.IntWithDefault("YOUR_ENV_VAR_NAME", 123)
```

Returns the `int` value of your env variable. If empty it will return the default int (in this case `123`).

```go
env.Bool("YOUR_ENV_VAR_NAME")
```

Returns the `bool` value of your env variable (`1` and `"true"` are equal to `true`).

```go
env.BoolWithDefault("YOUR_ENV_VAR_NAME", true)
```

Returns the `bool` value of your env variable (`1` and `"true"` are equal to `true`). If unset it will return the default bool (in this case `true`).
