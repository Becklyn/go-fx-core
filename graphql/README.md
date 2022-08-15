# GraphQL

The preconfigured `graphql.Schema` provides root objects for `Query`, `Mutation` and `Subscription`.

```go
app := fx.New(
	env.Module,
	logrus.Module,
	graphql.Module,
	fx.Invoke(func(schema *graphql.Schema) {
		// use schema here
	}),
)
app.Run()
```
