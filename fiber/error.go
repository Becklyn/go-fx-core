package fiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type Error struct {
	err        error
	statusCode int
}

func NewError(err error, statusCode int) *Error {
	return &Error{
		err:        err,
		statusCode: statusCode,
	}
}

func (e *Error) Error() string {
	return e.err.Error()
}

func (e *Error) StatusCode() int {
	return e.statusCode
}

func newErrorMiddleware(logger *logrus.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Next()
		if err == nil {
			return nil
		}

		fiberError, ok := err.(*Error)
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
