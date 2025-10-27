package applications

import (
	"context"
	"fmt"
	appError "hafiztri123/hv1-job-tracker/internal/error"
	"time"
)

func (r *ApplicationRepository) InsertApplication(req *CreateApplicationDto, userId string) error {
	createQuery := `
		insert into applications (
			user_id,
			company_name, 
			position_title,
			job_url,
			salary_range,
			location,
			status,
			notes,
			applied_date
		) values ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := r.db.Exec(ctx, createQuery, userId, req.CompanyName, req.PositionTitle, req.JobUrl, req.SalaryRange, req.Location, req.Status, req.Notes, req.AppliedDate)

	if err != nil {
		return err
	}

	return nil

}

func (r *ApplicationRepository) FindApplicationsById(userId string, queryParams ApplicationQueryParams) ([]Application, error) {
	fetchQuery := `
		select 
		id, 
		user_id, 
		company_name, 
		position_title, 
		job_url, 
		salary_range, 
		location, 
		status, 
		notes, 
		applied_date, 
		created_at, 
		updated_at, 
		deleted_at
		from applications where user_id = $1 and deleted_at is null
	`

	args := []any{userId}
	paramCount := 1

	if queryParams.Status != nil {
		paramCount++
		fetchQuery += fmt.Sprintf(" and status = $%d", paramCount)
		args = append(args, *queryParams.Status)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	rows, err := r.db.Query(ctx, fetchQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	applications := []Application{}

	for rows.Next() {
		app := new(Application)

		err := rows.Scan(
			&app.Id,
			&app.UserId,
			&app.CompanyName,
			&app.PositionTitle,
			&app.JobUrl,
			&app.SalaryRange,
			&app.Location,
			&app.Status,
			&app.Notes,
			&app.AppliedDate,
			&app.CreatedAt,
			&app.UpdatedAt,
			&app.DeletedAt,
		)

		if err != nil {
			return nil, err
		}

		applications = append(applications, *app)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return applications, err
}

func (r *ApplicationRepository) DeleteApplications(userId, applicationId string) error {
	updateQuery := `
		update applications
		set deleted_at = now()
		where id = $1 and user_id = $2 and deleted_at is null;
	`

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	result, err := r.db.Exec(ctx, updateQuery, applicationId, userId)
	if err != nil {
		return appError.NewInternalServerError(err.Error())
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return appError.NewNotFoundErr("Application not found")
	}

	return nil
}

func (r *ApplicationRepository) UpdateApplications(userId, applicationId string, body *UpdateApplicationDto) error {

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return appError.NewInternalServerError(err.Error())
	}
	defer func() {
		err = tx.Rollback(ctx)
		if err != nil {
			return
		}
	}()

	query := "update applications set updated_at = now()"
	args := []any{}
	paramCount := 0

	if body.CompanyName != nil {
		paramCount++
		query += fmt.Sprintf(" , company_name = $%d", paramCount)
		args = append(args, *body.CompanyName)
	}

	if body.PositionTitle != nil {
		paramCount++
		query += fmt.Sprintf(" , position_title = $%d", paramCount)
		args = append(args, *body.PositionTitle)
	}

	if body.JobUrl != nil {
		paramCount++
		query += fmt.Sprintf(" , job_url = $%d", paramCount)
		args = append(args, *body.JobUrl)
	}

	if body.SalaryRange != nil {
		paramCount++
		query += fmt.Sprintf(" , salary_range = $%d", paramCount)
		args = append(args, *body.SalaryRange)
	}

	if body.Location != nil {
		paramCount++
		query += fmt.Sprintf(" , location = $%d", paramCount)
		args = append(args, *body.Location)
	}

	if body.Status != nil {
		paramCount++
		query += fmt.Sprintf(" , status = $%d", paramCount)
		args = append(args, *body.Status)
	}

	if body.Notes != nil {
		paramCount++
		query += fmt.Sprintf(" , notes = $%d", paramCount)
		args = append(args, *body.Notes)
	}

	if body.AppliedDate != nil {
		paramCount++
		query += fmt.Sprintf(" , applied_date = $%d", paramCount)
		args = append(args, *body.AppliedDate)
	}

	query += fmt.Sprintf(" where id = $%d and user_id = $%d and deleted_at is null", paramCount+1, paramCount+2)
	args = append(args, applicationId, userId)

	result, err := tx.Exec(ctx, query, args...)
	if err != nil {
		return appError.NewInternalServerError(err.Error())
	}

	if result.RowsAffected() == 0 {
		return appError.NewNotFoundErr("Application not found")
	}

	if err := tx.Commit(ctx); err != nil {
		return appError.NewInternalServerError(err.Error())
	}

	return nil

}

func (r *ApplicationRepository) BatchDeleteApplications(userId string, applicationIds []string) error {
	if len(applicationIds) == 0 {
		return appError.NewBadRequestError("No application IDs provided")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	query := `
		update applications
		set deleted_at = now()
		where id = ANY($1) and user_id = $2 and deleted_at is null
	`

	result, err := r.db.Exec(ctx, query, applicationIds, userId)
	if err != nil {
		return appError.NewInternalServerError(err.Error())
	}

	if result.RowsAffected() == 0 {
		return appError.NewNotFoundErr("No applications found to delete")
	}

	return nil
}

func (r *ApplicationRepository) BatchUpdateStatusApplications(userId string, applicationIds []string, status string) error {
	if len(applicationIds) == 0 {
		return appError.NewBadRequestError("No application IDs provided")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	query := `
		update applications
		set status = $1, updated_at = now()
		where id = ANY($2) and user_id = $3 and deleted_at is null
	`

	result, err := r.db.Exec(ctx, query, status, applicationIds, userId)
	if err != nil {
		return appError.NewInternalServerError(err.Error())
	}

	if result.RowsAffected() == 0 {
		return appError.NewNotFoundErr("No applications found to update")
	}

	return nil
}
