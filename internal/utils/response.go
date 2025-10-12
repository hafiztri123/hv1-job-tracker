package utils

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Error   any    `json:"error,omitempty"`
}

type ResponseOption func(*Response)

func WithStatus(status int) ResponseOption {
	return func(r *Response) {
		r.Status = status
	}
}

func WithMessage(msg string) ResponseOption {
	return func(r *Response) {
		r.Message = msg
	}
}

func WithData(data any) ResponseOption {
	return func(r *Response) {
		r.Data = data
	}
}

func WithError(err any) ResponseOption {
	return func(r *Response) {
		r.Error = err
	}
}

func NewResponse(c *fiber.Ctx, opts ...ResponseOption) error {
	response := &Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    nil,
		Error:   nil,
	}

	for _, opt := range opts {
		opt(response)
	}

	return c.Status(response.Status).JSON(response)
}
