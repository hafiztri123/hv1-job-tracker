package router

import (
	"hafiztri123/hv1-job-tracker/internal/auth"
	"hafiztri123/hv1-job-tracker/internal/config"
	appError "hafiztri123/hv1-job-tracker/internal/error"
	"hafiztri123/hv1-job-tracker/internal/handler"
	"hafiztri123/hv1-job-tracker/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func NewRouter(h *handler.Handler, cfg fiber.Config, isDev bool) *fiber.App {
	app := fiber.New(cfg)

	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${method} ${path} (${latency})\n",
	}))

	app.Use(recover.New(config.NewRecoverConfig(isDev)))

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

	api.Post("/auth/register", h.RegisterUserHandler)
	api.Post("/auth/login", h.LoginUserHandler)
	api.Get("/health", h.HealthHandler)

	api.Get("/auth/verify", auth.AuthMiddleware, func(c *fiber.Ctx) error {
		userId, ok := c.Locals("userId").(string)
		if !ok {
			return appError.ErrUnauthorized
		}

		return utils.NewResponse(
			c,
			utils.WithData(userId),
		)
	})

	applications := api.Group("/applications")
	applications.Use(auth.AuthMiddleware)
	applications.Get("/", h.GetApplicationsHandler)
	applications.Post("/", h.CreateApplicationHandler)
	applications.Delete("/:id", h.DeleteApplicationHandler)
	applications.Put("/:id", h.UpdateApplicationHandler)
	applications.Get("/options", h.GetApplicationOptionsHandler)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"error": "Route not found",
		})
	})
}
