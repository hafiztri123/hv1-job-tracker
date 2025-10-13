package middleware

import (
	"errors"
	appError "hafiztri123/hv1-job-tracker/internal/error"
	"hafiztri123/hv1-job-tracker/internal/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler() fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error {

		var appError *appError.AppError
		if errors.As(err, &appError) {
			errorMsg := ""
			if appError.Err != nil {
				errorMsg = appError.Err.Error()
			}
			return utils.NewResponse(
				c,
				utils.WithStatus(appError.StatusCode),
				utils.WithMessage(appError.Message),
				utils.WithError(errorMsg),
			)
		}

		var fiberErr *fiber.Error
		if errors.As(err, &fiberErr) {
			return utils.NewResponse(
				c,
				utils.WithStatus(fiberErr.Code),
				utils.WithMessage(fiberErr.Message),
			)
		}

		return utils.NewResponse(
			c,
			utils.WithStatus(http.StatusInternalServerError),
			utils.WithError(err.Error()),
			utils.WithMessage("Internal Server Error"),
		)

	}
}
