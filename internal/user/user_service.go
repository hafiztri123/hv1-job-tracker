package user

import (
	"hafiztri123/hv1-job-tracker/internal/auth"
	appError "hafiztri123/hv1-job-tracker/internal/error"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func (u *UserService) RegisterUser(req *RegisterUserDto) error {
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

	if err := u.Repo.CreateUser(user); err != nil {
		return err
	}

	return nil
}

func (u *UserService) LoginUser(req *LoginUserDto) (string, error) {
	user, err := u.Repo.FindUserByEmail(req.Email)

	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return "", appError.New(
			err,
			"Invalid credentials",
			http.StatusBadRequest,
		)
	}

	token, err := auth.GenerateToken(user.ID.String(), user.Email)
	if err != nil {
		return "", err
	}

	return token, err
}
