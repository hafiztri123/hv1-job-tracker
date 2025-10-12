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

type RegisterUserDto struct {
	Email     string `json:"email" validate:"required,email"`
	FirstName string `json:"firstName" validate:"required,min=2,max=50"`
	LastName  string `json:"lastName" validate:"required,min=2,max=50"`
	Password  string `json:"password" validate:"required,min=8,max=64"`
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
