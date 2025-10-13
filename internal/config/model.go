package config

import (
	"hafiztri123/hv1-job-tracker/internal/applications"
	"hafiztri123/hv1-job-tracker/internal/user"
)

type Config struct {
	DbAddr     string
	DbMaxConns int32
}

type Services struct {
	UserService        *user.UserService
	ApplicationService *applications.ApplicationService
}

type Repositories struct {
	UserRepository        *user.UserRepository
	ApplicationRepository *applications.ApplicationRepository
}
