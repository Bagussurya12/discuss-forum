package posts

import (
	"context"

	"github.com/Bagussurya12/discuss-forum/source/model/posts"
)

func (r *repository) CreateComment(ctx context.Context, model posts.CommentModel) error {

	query := `INSERT INTO comments(post_id, user_id, comment_content, created_at, updated_at, created_by, updated_by) VALUES (?,?,?,?,?,?,?)`

	_, err := r.db.ExecContext(ctx, query, model.PostID, model.UserID, model.CommentContent, model.CreatedAt, model.UpdatedAt, model.CreatedBy, model.UpdatedBy)

	if err != nil {
		return err
	}

	return nil
}
