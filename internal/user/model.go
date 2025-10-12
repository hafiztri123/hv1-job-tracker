package user

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	ID           uuid.UUID `json:"id"`
	Email        string    `json:"email"`
	FirstName    string    `json:"firstName"`
	LastName     string    `json:"lastName"`
	PasswordHash string    `json:"-"`
	CreatedAt    int64     `json:"createdAt"`
	UpdatedAt    int64     `json:"updatedAt"`
	DeletedAt    int64     `json:"deletedAt"`
}

type UserRepository struct {
	Db *pgxpool.Pool
}

type UserService struct {
	Repo *UserRepository
}

func NewUserService(repo *UserRepository) *UserService {
	return &UserService{
		Repo: repo,
	}
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		Db: db,
	}
}
