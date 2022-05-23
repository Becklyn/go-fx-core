package graphql

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(
		fx.Annotated{
			Name:   "query",
			Target: newGraphQlQuery,
		},
		fx.Annotated{
			Name:   "mutation",
			Target: newGraphQlMutation,
		},
		fx.Annotated{
			Name:   "subscription",
			Target: newGraphQlSubscribtion,
		},
		newGraphQlSchema,
	),
)
