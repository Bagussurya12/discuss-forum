package posts

import (
	"context"

	"github.com/Bagussurya12/discuss-forum/source/model/posts"
	"github.com/rs/zerolog/log"
)

func (s *service) GetPostById(ctx context.Context, postId int64) (*posts.GetPostResponse, error) {
	postDetail, err := s.postRepo.GetPostById(ctx, postId)

	if err != nil {
		log.Error().Err(err).Msg("error get post by id")
		return nil, err
	}

	likeCount, err := s.postRepo.CountLikePostId(ctx, postId)
	if err != nil {
		log.Error().Err(err).Msg("error count like")
		return nil, err
	}

	comments, err := s.postRepo.GetCommentByPostId(ctx, postId)
	if err != nil {
		log.Error().Err(err).Msg("error get comment")
		return nil, err
	}

	return &posts.GetPostResponse{
		PostDetail: posts.Post{
			ID:           postDetail.ID,
			UserID:       postDetail.UserID,
			Username:     postDetail.Username,
			PostTitle:    postDetail.PostTitle,
			PostContent:  postDetail.PostContent,
			PostHashtags: postDetail.PostHashtags,
			IsLiked:      postDetail.IsLiked,
		},
		LikeCount: likeCount,
		Comments:  comments,
	}, nil
}
