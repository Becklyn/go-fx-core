# CQRS

A extremely simple set of interfaces for cqrs usage:

- `cqrs.CommandHandler`
- `cqrs.QueryHandler`
- `cqrs.EventHandler`

## CommandHandler

```go
type ExampleCommand struct {
	Value string
}

type ExampleHandler struct {
	repo ExampleRepo
}

func NewExampleHandler(
	repo ExampleRepo,
) cqrs.CommandHandler[*ExampleCommand] {
	return &ExampleHandler{
		repo:    repo,
	}
}

func (h *ExampleHandler) Handle(command *ExampleCommand) error {
	return repo.DoSomething(command.Value)
}
```

## QueryHandler

```go
type ExampleQuery struct {
	Id string
}

type ExampleHandler struct {
	repo ExampleRepo
}

func NewExampleHandler(
	repo ExampleRepo,
) cqrs.QueryHandler[*ExampleQuery, *Entity] {
	return &ExampleHandler{
		repo:    repo,
	}
}

func (h *ExampleHandler) Handle(query *ExampleQuery) (*Entity, error) {
	return repo.GetExampleEntity(query.Id)
}
```

## EventHandler

```go
type ExampleEvent struct {
	Value string
}

type ExampleHandler struct {
	service ExampleService
}

func NewExampleHandler(
	service ExampleService,
) cqrs.EventHandler[*ExampleEvent] {
	return &ExampleHandler{
		service:    service,
	}
}

func (h *ExampleHandler) Handle(event *ExampleEvent) error {
	return service.Handle(event.Value)
}
```
