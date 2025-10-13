package router

import (
	"hafiztri123/hv1-job-tracker/internal/auth"
	"hafiztri123/hv1-job-tracker/internal/handler"
	"hafiztri123/hv1-job-tracker/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func NewRouter(h *handler.Handler, cfg fiber.Config) *fiber.App {
	app := fiber.New(cfg)

	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${method} ${path} (${latency})\n",
	}))

	app.Use(recover.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins: utils.GetEnv("CORS_ORIGIN", "*"),
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))

	setupRoutes(app, h)

	return app

}

func setupRoutes(app *fiber.App, h *handler.Handler) {
	api := app.Group("/api/v1")

	api.Post("/user/register", h.RegisterUserHandler)
	api.Post("/user/login", h.LoginUserHandler)
	api.Get("/health", h.HealthHandler)

	applications := api.Group("/applications")
	applications.Use(auth.AuthMiddleware)
	applications.Get("/", h.GetApplicationsHandler)
	applications.Post("/", h.CreateApplicationHandler)
	applications.Put("/:id", h.DeleteApplicationHandler)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"error": "Route not found",
		})
	})
}
