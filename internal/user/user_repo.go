package user

import (
	"context"
	"errors"
	"fmt"
	appError "hafiztri123/hv1-job-tracker/internal/error"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
)

func (r *UserRepository) CreateUser(user *User) error {
	createQuery := `insert into users (
	email, 
	first_name, 
	last_name, 
	password_hash
	) values ($1, $2, $3, $4)`

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := r.Db.Exec(ctx, createQuery, user.Email, user.FirstName, user.LastName, user.PasswordHash)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				switch pgErr.ConstraintName {
				case "idx_users_email_lower":
					return appError.ErrDuplicateEmail
				default:
					return appError.New(
						fmt.Errorf("duplicated data in: %s", pgErr.ConstraintName),
						"duplicated data",
						http.StatusConflict,
					)
				}
			}
		}

		return appError.New(
			fmt.Errorf("database error: %s", err),
			"internal server error",
			http.StatusInternalServerError,
		)
	}

	return nil
}
