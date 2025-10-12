package applications

import (
	"context"
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
			applied_date,
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

func (r *ApplicationRepository) FindApplicationsById(userId string) ([]Application, error) {
	fetchQuery := `
		select id, user_id, company_name, position_title, job_url, salary_range, location, status, notes, applied_date, created_at, updated_at, deleted_at
		from applications where user_id = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	rows, err := r.db.Query(ctx, fetchQuery, userId)
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
