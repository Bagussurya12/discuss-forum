package memberships

import (
	"context"

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
