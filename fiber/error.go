package fiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type FiberError struct {
	err        error
	statusCode int
}

func NewFiberError(err error, statusCode int) *FiberError {
	return &FiberError{
		err:        err,
		statusCode: statusCode,
	}
}

func (e *FiberError) Error() string {
	return e.err.Error()
}

func (e *FiberError) StatusCode() int {
	return e.statusCode
}

func newErrorMiddleware(logger *logrus.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Next()
		if err == nil {
			return nil
		}

		fiberError, ok := err.(*FiberError)
		if !ok {
			return err
		}

		logger.Error(fiberError)

		if err := c.SendString(err.Error()); err != nil {
			return err
		}
		return c.SendStatus(fiberError.statusCode)
	}
}
