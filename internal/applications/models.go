package applications

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Application struct {
	Id            uuid.UUID  `json:"id"`
	UserId        uuid.UUID  `json:"userId"`
	CompanyName   string     `json:"companyName"`
	PositionTitle string     `json:"positionTitle"`
	JobUrl        *string    `json:"jobUrl"`
	SalaryRange   *string    `json:"salaryRange"`
	Location      *string    `json:"location"`
	Status        *string    `json:"status"`
	Notes         *string    `json:"notes"`
	AppliedDate   *time.Time `json:"appliedDate"`
	CreatedAt     time.Time  `json:"createdAt"`
	UpdatedAt     *time.Time `json:"updatedAt"`
	DeletedAt     *time.Time `json:"deletedAt"`
}

type ApplicationOptions struct {
	StatusOption []string `json:"statusOption"`
}

type ApplicationRepository struct {
	db *pgxpool.Pool
}
type ApplicationService struct {
	repo *ApplicationRepository
}

func NewApplicationService(repo *ApplicationRepository) *ApplicationService {
	return &ApplicationService{
		repo: repo,
	}
}

func NewApplicationRepository(db *pgxpool.Pool) *ApplicationRepository {
	return &ApplicationRepository{
		db: db,
	}
}
