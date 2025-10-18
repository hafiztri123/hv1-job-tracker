package user

import (
	"context"
	"database/sql"
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

func (r *UserRepository) FindUserByEmail(email string) (*User, error) {
	fetchQuery := `select id, email, first_name, last_name, password_hash from users
		where email = $1 and deleted_at is null
	`

	user := new(User)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := r.Db.QueryRow(ctx, fetchQuery, email).Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.PasswordHash,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, appError.ErrNotFound
		}

		return nil, err

	}

	return user, nil

}

func (r *UserRepository) FindUserById(id string) (*User, error) {
	fetchQuery := `select id, first_name, last_name from users where id = $1 and deleted_at is null`

	user := new(User)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := r.Db.QueryRow(ctx, fetchQuery, id).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, appError.ErrNotFound
		}
		return nil, err
	}

	return user, nil
}
