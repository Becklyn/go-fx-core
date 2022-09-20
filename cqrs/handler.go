package cqrs

import "context"

type CommandHandler[TCommand any] interface {
	Handle(ctx context.Context, command TCommand) error
}

type QueryHandler[TQuery any, TResult any] interface {
	Handle(ctx context.Context, query TQuery) (TResult, error)
}

type EventHandler[TEvent any] interface {
	Handle(ctx context.Context, event TEvent) error
}
