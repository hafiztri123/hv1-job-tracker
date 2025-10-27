package handler

import (
	"hafiztri123/hv1-job-tracker/internal/applications"
	"hafiztri123/hv1-job-tracker/internal/config"
	appError "hafiztri123/hv1-job-tracker/internal/error"
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
			utils.WithError(err.Error()),
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

	var queryParams applications.ApplicationQueryParams

	if err := c.QueryParser(&queryParams); err != nil {
		return appError.NewBadRequestError(err.Error())
	}

	userId, ok := c.Locals("userId").(string)
	if !ok {
		return utils.NewResponse(
			c,
			utils.WithMessage("Unauthorized"),
			utils.WithStatus(http.StatusUnauthorized),
		)
	}

	applications, err := h.ApplicationService.GetApplications(userId, queryParams)
	if err != nil {
		return err
	}

	return utils.NewResponse(
		c,
		utils.WithMessage("Successfully get applications"),
		utils.WithData(applications),
	)
}

func (h *Handler) DeleteApplicationHandler(c *fiber.Ctx) error {

	id := c.Params("id")
	userId, ok := c.Locals("userId").(string)

	if !ok {
		return appError.ErrUnauthorized
	}

	if id == "" {
		return appError.ErrInvalidInput
	}

	err := h.ApplicationService.DeleteApplications(userId, id)

	if err != nil {
		return err
	}

	return utils.NewResponse(
		c,
		utils.WithMessage("Successfully deleted applications"),
	)
}

func (h *Handler) GetApplicationOptionsHandler(c *fiber.Ctx) error {
	var queryParams applications.ApplicationOptionQueryParams

	if err := c.QueryParser(&queryParams); err != nil {
		return appError.NewBadRequestError(err.Error())
	}

	options := h.ApplicationService.GetApplicationOptions(queryParams)

	return utils.NewResponse(
		c,
		utils.WithMessage("Successfully get application options"),
		utils.WithData(options),
	)

}

func (h *Handler) UpdateApplicationHandler(c *fiber.Ctx) error {
	var body applications.UpdateApplicationDto

	if err := c.BodyParser(&body); err != nil {
		return appError.NewBadRequestError(err.Error())
	}

	if errors := utils.ValidateStruct(body); errors != nil {
		return utils.NewResponse(
			c,
			utils.WithMessage("Bad Request"),
			utils.WithStatus(http.StatusBadRequest),
			utils.WithError(errors),
		)
	}

	userId, ok := c.Locals("userId").(string)
	if !ok {
		return appError.ErrUnauthorized
	}

	applicationId := c.Params("id")
	if applicationId == "" {
		return appError.NewBadRequestError("Application id is missing")
	}

	err := h.ApplicationService.UpdateApplication(body, userId, applicationId)
	if err != nil {
		return err
	}

	return utils.NewResponse(
		c,
		utils.WithMessage("Application updated"),
	)

}

func (h *Handler) BatchDeleteApplicationHandler(c *fiber.Ctx) error {
	var dto applications.BatchDeleteDto

	if err := c.BodyParser(&dto); err != nil {
		return appError.NewBadRequestError(err.Error())
	}

	if errors := utils.ValidateStruct(dto); errors != nil {
		return utils.NewResponse(
			c,
			utils.WithMessage("Bad Request"),
			utils.WithStatus(http.StatusBadRequest),
			utils.WithError(errors),
		)
	}

	userId, ok := c.Locals("userId").(string)
	if !ok {
		return appError.ErrUnauthorized
	}

	err := h.ApplicationService.BatchDeleteApplications(userId, &dto)
	if err != nil {
		return err
	}

	return utils.NewResponse(
		c,
		utils.WithMessage("Applications deleted successfully"),
	)
}

func (h *Handler) BatchUpdateStatusApplicationHandler(c *fiber.Ctx) error {
	var dto applications.BatchUpdateStatusDto

	if err := c.BodyParser(&dto); err != nil {
		return appError.NewBadRequestError(err.Error())
	}

	if errors := utils.ValidateStruct(dto); errors != nil {
		return utils.NewResponse(
			c,
			utils.WithMessage("Bad Request"),
			utils.WithStatus(http.StatusBadRequest),
			utils.WithError(errors),
		)
	}

	userId, ok := c.Locals("userId").(string)
	if !ok {
		return appError.ErrUnauthorized
	}

	err := h.ApplicationService.BatchUpdateStatusApplications(userId, &dto)
	if err != nil {
		return err
	}

	return utils.NewResponse(
		c,
		utils.WithMessage("Application statuses updated successfully"),
	)
}

func (h *Handler) VerifyTokenHandler(c *fiber.Ctx) error {
	userId, ok := c.Locals("userId").(string)
	if !ok {
		return appError.ErrUnauthorized
	}

	fullUser, err := h.UserService.Repo.FindUserById(userId)

	if err != nil {
		return err
	}

	resp := new(user.GetUserDetailResponse)

	resp.ID = fullUser.ID
	resp.FirstName = fullUser.FirstName
	resp.LastName = fullUser.LastName

	return utils.NewResponse(
		c,
		utils.WithData(resp),
	)
}

func (h *Handler) LogoutHandler(c *fiber.Ctx) error {
	return utils.NewResponse(
		c,
		utils.WithMessage("Logged out successfully"),
	)
}
