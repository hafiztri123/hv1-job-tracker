package handler

import (
	"hafiztri123/hv1-job-tracker/internal/applications"
	"hafiztri123/hv1-job-tracker/internal/user"
)

type Handler struct {
	UserService        *user.UserService
	ApplicationService *applications.ApplicationService
}
