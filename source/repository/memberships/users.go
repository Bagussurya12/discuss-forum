package memberships

import (
	"context"
	"database/sql"

	"github.com/Bagussurya12/discuss-forum/source/model/memberships"
)

func (r *repository) GetUser(ctx context.Context, email, username, phone_number string) (*memberships.UserModel, error) {
	q := `SELECT id, email, password, username, phone_number, created_at, updated_at, created_by, updated_by FROM users WHERE email = ? OR username = ?`

	row := r.db.QueryRowContext(ctx, q, email, username)

	var response memberships.UserModel
	err := row.Scan(&response.ID, &response.Email, &response.Password, &response.Username, &response.PhoneNumber, &response.CreatedAt, &response.UpdatedAt, &response.CreatedBy, &response.UpdatedBy)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &response, nil
}

func (r *repository) CreateUser(ctx context.Context, model memberships.UserModel) error {
	q := `INSERT INTO users (email, password, username, phone_number, created_at, updated_at, created_by, updated_by) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := r.db.ExecContext(ctx, q, model.Email, model.Password, model.Username, model.PhoneNumber, model.CreatedAt, model.UpdatedAt, model.CreatedBy, model.UpdatedBy)

	if err != nil {
		return err
	}

	return nil
}
