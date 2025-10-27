package middleware

import (
	"errors"
	appError "hafiztri123/hv1-job-tracker/internal/error"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(isDev bool) fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError
		message := "Internal Server Error"

		var appError *appError.AppError
		if errors.As(err, &appError) {
			code = appError.StatusCode
			message = appError.Message
		}

		var fiberErr *fiber.Error
		if errors.As(err, &fiberErr) {
			code = fiberErr.Code
			message = fiberErr.Message
		}

		response := fiber.Map{
			"status":  code,
			"message": message,
		}

		if isDev {
			response["error"] = err.Error()
			response["path"] = c.Path()
		}

		return c.Status(code).JSON(response)
	}
}
