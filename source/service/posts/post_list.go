package posts

import (
	"context"

	"github.com/Bagussurya12/discuss-forum/source/model/posts"
	"github.com/rs/zerolog/log"
)

func (s *service) GetAllPost(ctx context.Context, pageSize, pageIndex int) (posts.GetAllPostResponse, error) {
	limit := pageSize
	offset := pageSize * (pageIndex - 1)
	response, err := s.postRepo.GetAllPost(ctx, limit, offset)

	if err != nil {
		log.Error().Err(err).Msg("Failed Get All Posts")
		return response, err
	}
	return response, nil
}
