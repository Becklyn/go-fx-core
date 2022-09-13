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

func errorMiddleware(logger *logrus.Logger) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		err := ctx.Next()
		if err == nil {
			return nil
		}

		fiberError, ok := err.(*Error)
		if !ok {
			return err
		}

		logger.Error(fiberError)

		ctx.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)

		if err := ctx.SendString(err.Error()); err != nil {
			return err
		}
		return ctx.SendStatus(fiberError.statusCode)
	}
}
