package user

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	ID           uuid.UUID `json:"id"`
	Email        string    `json:"email"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	PasswordHash string    `json:"-"`
	CreatedAt    int64     `json:"created_at"`
	UpdatedAt    int64     `json:"updated_at"`
	DeletedAt    int64     `json:"deleted_at"`
}

type UserRepository struct {
	Db *pgxpool.Pool
}

type UserService struct {
	Repo *UserRepository
}

type RegisterUserDto struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
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
