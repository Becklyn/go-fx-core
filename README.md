# golang fx core module

## Installation

Adding _go-fx-core_ to a Go project is as easy as calling this command

```shell
go get -u github.com/Becklyn/go-fx-core
```

## Modules

We provide modules for common use cases:

- `core`: Collection of common modules used for basic webserver applications: webserver (fiber), env, logging (logrus), metrics, readyness, health.
- `cqrs`: Basic interfaces for apps that do cqrs
- `env`: Environment variables processing
- `fiber`: Webserver
- `graphql`: GraphQL
- `grpc`: Basis of a gRPC server
- `health`: Health indication
- `logrus`: Logging
- `metrics`: Automatic webserver metrics
- `readyness`: Readyness indication
- `validator`: Basic validation service

You can see detailed documentation in the `README.md` based at the root of the packages.

## Using the core module

The project must be based on the [uber-go/fx](https://github.com/uber-go/fx) application framework

```go
package main

import "go.uber.org/fx"

func main() {
    // Creates a new fx application
    fx.New(
        // Add the core module to the container
        core.Module,
		// Add more modules that your application needs (not all provided modules are pert of the core.Module)
    ).Run()
}
```

## Integrated 3rd party modules

The list of 3rd party libraries that we provide as uber/fx modules by this package:

- Fiber webserver - https://github.com/gofiber/fiber
- GraphQL - https://github.com/graphql-go/graphql
- gRPC - https://github.com/grpc/grpc-go
- Logrus - https://github.com/sirupsen/logrus
- Prometheus metrics - https://github.com/prometheus/client_golang
- Go-Playground validator - https://github.com/go-playground/validator
