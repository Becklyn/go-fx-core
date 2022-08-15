package fiber_test

import (
	"errors"
	"testing"

	"github.com/Becklyn/go-fx-core/fiber"
	gofiber "github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestFiberError(t *testing.T) {
	expectedErr := errors.New("test")
	expectedStatus := gofiber.StatusInternalServerError

	err := fiber.NewError(expectedErr, expectedStatus)

	assert.Equal(t, expectedErr.Error(), err.Error())
	assert.Equal(t, expectedStatus, err.StatusCode())
}
