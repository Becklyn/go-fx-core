package core

import (
	"github.com/graphql-go/graphql"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

var GraphQlModule = fx.Options(
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

func newGraphQlQuery() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name:        "Query",
		Description: "The application's root query object",
		Fields:      graphql.Fields{},
	})
}

func newGraphQlMutation() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name:        "Mutation",
		Description: "The application's root mutation object",
		Fields:      graphql.Fields{},
	})
}

func newGraphQlSubscribtion() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name:        "Subscription",
		Description: "The application's root subscription object",
		Fields:      graphql.Fields{},
	})
}

type graphqlDependencies struct {
	fx.In

	Query        *graphql.Object `name:"query"`
	Mutation     *graphql.Object `name:"mutation"`
	Subscription *graphql.Object `name:"subscription"`
}

func newGraphQlSchema(dependencies graphqlDependencies, logger *logrus.Logger) *graphql.Schema {
	query := dependencies.Query
	mutation := dependencies.Mutation
	subscription := dependencies.Subscription

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: func() *graphql.Object {
			if len(query.Fields()) == 0 {
				return nil
			}
			return query
		}(),
		Mutation: func() *graphql.Object {
			if len(mutation.Fields()) == 0 {
				return nil
			}
			return mutation
		}(),
		Subscription: func() *graphql.Object {
			if len(subscription.Fields()) == 0 {
				return nil
			}
			return subscription
		}(),
	})
	if err != nil {
		logger.Error(err)
	}

	return &schema
}
