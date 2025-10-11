package user

import (
	"context"

	"golang.org/x/crypto/bcrypt"
)

func NewUserService(repo *UserRepository) *UserService {
	return &UserService{
		Repo: repo,
	}
}

func (u *UserService) RegisterUser(req *RegisterUserDto, ctx context.Context) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 12)
	if err != nil {
		return err
	}

	user := &User{
		Email:        req.Email,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		PasswordHash: string(hashedPassword),
	}

	if err := u.Repo.CreateUser(user, ctx); err != nil {
		return err
	}

	return nil
}
