package cqrs

type CommandHandler[TCommand any] interface {
	Handle(command TCommand) error
}

type QueryHandler[TQuery any, TResult any] interface {
	Handle(query TQuery) (TResult, error)
}

type EventHandler[TEvent any] interface {
	Handle(event TEvent) error
}
