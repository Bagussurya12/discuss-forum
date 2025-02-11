package posts

import (
	"context"
	"database/sql"

	"github.com/Bagussurya12/discuss-forum/source/model/posts"
)

func (r *repository) GetUserActivity(ctx context.Context, model posts.UserActivityModel) (*posts.UserActivityModel, error) {
	query := `SELECT id, post_id, user_id, is_liked, created_at, updated_at, created_by, updated_by FROM user_activities WHERE post_id = ? AND user_id = ?`

	var res posts.UserActivityModel

	row := r.db.QueryRowContext(ctx, query, model.PostID, model.UserID)

	err := row.Scan(&res.ID, &res.PostID, &res.UserID, &res.IsLiked, &res.CreatedAt, &res.UpdatedAt, &res.CreatedBy, &res.UpdatedBy)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &res, nil
}
