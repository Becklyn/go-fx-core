# golang fx core module

## Installation

Adding _go-fx-core_ to a Go project is as easy as calling this command

```shell
go get -u github.com/Becklyn/go-fx-core
```

## Using the module

The project must be based on the [uber-go/fx](https://github.com/uber-go/fx) application framework

```go
package main

import "go.uber.org/fx"

func main() {
    // Creates a new fx application
    fx.New(
        // Add the core module to the container
        core.Module,
    ).Run()
}
```

## Contained sub-modules

The list of uber/fx modules that are currently available:

- Logrus - https://github.com/sirupsen/logrus
- Fiber webserver - https://github.com/gofiber/fiber
- Go-Playground validator - https://github.com/go-playground/validator
