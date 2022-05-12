package core

import (
	"github.com/go-playground/validator/v10"
	"go.uber.org/fx"
)

var ValidatorModule = fx.Provide(
	newValidator,
)

func newValidator() *validator.Validate {
	return validator.New()
}
