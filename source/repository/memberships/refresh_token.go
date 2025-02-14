package memberships

import (
	"context"
	"database/sql"
	"time"

	"github.com/Bagussurya12/discuss-forum/source/model/memberships"
)

func (r *repository) InsertRefreshToken(ctx context.Context, model memberships.RefreshTokenModel) error {
	query := `INSERT INTO refresh_token (user_id,refresh_token, expired_at, created_at, updated_at) VALUES (?,?,?,?,?,?,?)`

	_, err := r.db.ExecContext(ctx, query, model.UserID, model.RefreshToken, model.ExpiredAt, model.CreatedAt, model.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetRefreshToken(ctx context.Context, userID int64, now time.Time) (*memberships.RefreshTokenModel, error) {
	query := `SELECT id, user_id, refresh_token, expired_at, created_at, updated_at FROM refresh_token WHERE user_id = ? AND expired_at >= ?`

	var response memberships.RefreshTokenModel
	row := r.db.QueryRowContext(ctx, query, userID, now)

	err := row.Scan(&response.ID, &response.UserID, &response.RefreshToken, &response.ExpiredAt, &response.CreatedAt, &response.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &response, nil
}
