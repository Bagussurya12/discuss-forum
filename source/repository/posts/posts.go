package posts

import (
	"context"
	"log"

	"github.com/Bagussurya12/discuss-forum/source/model/posts"
)

func (r *repository) CreatePost(ctx context.Context, model posts.PostModel) error {
	q := `INSERT INTO posts (user_id, post_title, post_content, post_hastags, created_at, updated_at, created_by, updated_by) 
	VALUES 
	(?,?,?,?,?,?,?,?)
	`
	_, err := r.db.ExecContext(ctx, q, model.UserID, model.PostTitle, model.PostContent, model.PostHashtags, model.CreatedAt, model.UpdatedAt, model.CreatedBy, model.UpdatedBy)

	if err != nil {
		log.Fatal("Something When Wrong", err)
		return err
	}

	return nil
}
