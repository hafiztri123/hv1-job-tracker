package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
)

func (r *UserRepository) CreateUser(user *User, ctx context.Context) error {
	createQuery := `insert into users (
	email, 
	first_name, 
	last_name, 
	password_hash
	) values ($1, $2, $3, $4)`

	_, err := r.Db.Exec(ctx, createQuery, user.Email, user.FirstName, user.LastName, user.PasswordHash)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				switch pgErr.ConstraintName {
				case "idx_users_email_lower":
					return ErrDuplicateEmail
				default:
					return fmt.Errorf("duplicate value in constraint: %s", pgErr.ConstraintName)
				}
			}
		}

		return ErrDatabase
	}

	return nil
}
