package posts

import (
	"context"
	"log"
	"strings"

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

func (r *repository) GetAllPost(ctx context.Context, limit, offset int) (posts.GetAllPostResponse, error) {
	query := `SELECT p.id, p.user_id, u.username, p.post_title, p.post_content, p.post_hastags FROM posts p JOIN users u ON p.user_id = u.id  ORDER BY p.updated_at DESC LIMIT ? OFFSET ?`

	response := posts.GetAllPostResponse{}
	rows, err := r.db.QueryContext(ctx, query, limit, offset)

	if err != nil {
		return posts.GetAllPostResponse{}, err
	}

	defer rows.Close()

	data := make([]posts.Post, 0)
	for rows.Next() {

		var (
			model    posts.PostModel
			username string
		)
		err = rows.Scan(&model.ID, &model.UserID, &username, &model.PostTitle, &model.PostContent, &model.PostHashtags)

		if err != nil {
			return response, err
		}

		data = append(data, posts.Post{
			ID:           model.ID,
			UserID:       model.UserID,
			Username:     username,
			PostTitle:    model.PostTitle,
			PostContent:  model.PostContent,
			PostHashtags: strings.Split(model.PostHashtags, ","),
		})
	}
	response.Data = data
	response.Pagination = posts.Pagination{
		Limit:  limit,
		Offset: offset,
	}
	return response, nil
}

func (r *repository) GetPostById(ctx context.Context, id int64) (*posts.Post, error) {
	query := `SELECT p.id, p.user_id, u.username, p.post_title, p.post_content, p.post_hastags ua.is_liked FROM posts p JOIN users u ON p.user_id = u.id JOIN user_activities ua ON ua.post_id = p.id WHERE p.id = ?`

	var (
		model    posts.PostModel
		username string
		isLiked  bool
	)
	row := r.db.QueryRowContext(ctx, query, id)

	err := row.Scan(&model.ID, &model.UserID, &username, &model.PostTitle, &model.PostContent, &model.PostHashtags, &isLiked)

	if err != nil {
		return nil, err
	}
	return &posts.Post{
		ID:           model.ID,
		UserID:       model.UserID,
		Username:     username,
		PostTitle:    model.PostTitle,
		PostContent:  model.PostContent,
		PostHashtags: strings.Split(model.PostHashtags, ","),
		IsLiked:      isLiked,
	}, nil
}
