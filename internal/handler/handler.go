package handler

import (
	"hafiztri123/hv1-job-tracker/internal/applications"
	"hafiztri123/hv1-job-tracker/internal/config"
	"hafiztri123/hv1-job-tracker/internal/user"
	"hafiztri123/hv1-job-tracker/internal/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func NewHandler(services *config.Services) *Handler {
	return &Handler{
		UserService:        services.UserService,
		ApplicationService: services.ApplicationService,
	}
}

func (h *Handler) HealthHandler(c *fiber.Ctx) error {
	return c.SendString("OK")
}

func (h *Handler) RegisterUserHandler(c *fiber.Ctx) error {
	dto := new(user.RegisterUserDto)

	if err := c.BodyParser(dto); err != nil {
		return utils.NewResponse(
			c,
			utils.WithMessage("Bad Request"),
			utils.WithStatus(http.StatusBadRequest),
		)
	}

	if errors := utils.ValidateStruct(dto); errors != nil {
		return utils.NewResponse(
			c,
			utils.WithStatus(http.StatusBadRequest),
			utils.WithMessage("invalid request body"),
			utils.WithError(errors),
		)
	}

	err := h.UserService.RegisterUser(dto)

	if err != nil {
		return err
	}

	return utils.NewResponse(
		c,
		utils.WithMessage("User created"),
		utils.WithStatus(http.StatusCreated),
	)
}

func (h *Handler) LoginUserHandler(c *fiber.Ctx) error {
	dto := new(user.LoginUserDto)

	if err := c.BodyParser(dto); err != nil {
		return utils.NewResponse(
			c,
			utils.WithMessage("Bad Request"),
			utils.WithStatus(http.StatusBadRequest),
		)
	}

	if errors := utils.ValidateStruct(dto); errors != nil {
		return utils.NewResponse(
			c,
			utils.WithStatus(http.StatusBadRequest),
			utils.WithMessage("invalid request body"),
			utils.WithError(errors),
		)
	}

	token, err := h.UserService.LoginUser(dto)
	if err != nil {
		return err
	}

	return utils.NewResponse(
		c,
		utils.WithMessage("Login success"),
		utils.WithData(token),
	)
}

func (h *Handler) CreateApplicationHandler(c *fiber.Ctx) error {
	dto := new(applications.CreateApplicationDto)

	if err := c.BodyParser(dto); err != nil {
		return utils.NewResponse(
			c,
			utils.WithMessage("Bad Request"),
			utils.WithStatus(http.StatusBadRequest),
		)
	}

	if errors := utils.ValidateStruct(dto); errors != nil {
		return utils.NewResponse(
			c,
			utils.WithStatus(http.StatusBadRequest),
			utils.WithMessage("invalid request body"),
			utils.WithError(errors),
		)
	}

	userId, ok := c.Locals("userId").(string)
	if !ok {
		return utils.NewResponse(
			c,
			utils.WithMessage("Unauthorized"),
			utils.WithStatus(http.StatusUnauthorized),
		)
	}

	err := h.ApplicationService.CreateApplication(dto, userId)
	if err != nil {
		return err
	}

	return utils.NewResponse(
		c,
		utils.WithMessage("Application created"),
		utils.WithStatus(http.StatusCreated),
	)
}

func (h *Handler) GetApplicationsHandler(c *fiber.Ctx) error {

	userId, ok := c.Locals("userId").(string)
	if !ok {
		return utils.NewResponse(
			c,
			utils.WithMessage("Unauthorized"),
			utils.WithStatus(http.StatusUnauthorized),
		)
	}

	applications, err := h.ApplicationService.GetApplications(userId)
	if err != nil {
		return err
	}

	return utils.NewResponse(
		c,
		utils.WithMessage("Application created"),
		utils.WithStatus(http.StatusCreated),
		utils.WithData(applications),
	)
}
