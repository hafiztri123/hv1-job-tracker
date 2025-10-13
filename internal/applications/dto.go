package applications

import "time"

type CreateApplicationDto struct {
	//Required
	CompanyName   string `json:"companyName" validate:"required,min=2,max=255"`
	PositionTitle string `json:"positionTitle" validate:"required,min=2,max=255"`

	//Optional
	JobUrl      *string    `json:"jobUrl" validate:"omitempty,url"`
	SalaryRange *string    `json:"salaryRange" validate:"omitempty,min=2,max=100"`
	Location    *string    `json:"location" validate:"omitempty,min=2,max=100"`
	Status      *string    `json:"status" validate:"omitempty,min=2,max=50"`
	Notes       *string    `json:"notes" validate:"omitempty"`
	AppliedDate *time.Time `json:"appliedDate" validate:"omitempty"`
}

type ApplicationOptionQueryParams struct {
	StatusOption bool `json:"statusOption"`
}
