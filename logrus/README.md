# Logrus (logging)

```go
fx.Invoke(func(logger *logrus.Logger) {
	// use logger here
}),
```

The log level can be configured using the `LOG_LEVEL` env variable. Default level is `info`.

Possible values:

- debug
- info (this is the default / fallback level)
- warn
- error
- fatal
