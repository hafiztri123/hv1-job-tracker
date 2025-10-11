package handler

import (
	"hafiztri123/hv1-job-tracker/internal/config"

	"github.com/gofiber/fiber/v2"
)

func NewHandler(services *config.Services) *Handler {
	return &Handler{
		UserService: services.UserService,
	}
}

func (h *Handler) HealthHandler(c *fiber.Ctx) error {
	return c.SendString("OK")
}
