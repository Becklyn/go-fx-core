# GraphQL

The `GRPC_ADDR` env variable is used to configure the address of the server. Default: `tcp://0.0.0.0:9000`.

```go
app := fx.New(
	env.Module,
	logrus.Module,
	grpc.Module,
	fx.Invoke(func(server *grpc.Server) {
		// use server here
	}),
)
app.Run()
```
