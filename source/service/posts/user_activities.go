package posts

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/Bagussurya12/discuss-forum/source/model/posts"
	"github.com/rs/zerolog/log"
)

func (s *service) UserActivity(ctx context.Context, postID, userID int64, request posts.UserActivityRequest) error {

	now := time.Now()
	model := posts.UserActivityModel{
		PostID:    postID,
		UserID:    userID,
		IsLiked:   request.IsLiked,
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: strconv.FormatInt(userID, 10),
		UpdatedBy: strconv.FormatInt(userID, 10),
	}
	userActivity, err := s.postRepo.GetUserActivity(ctx, model)

	if err != nil {
		log.Error().Err(err).Msg("Error User Activity From DB")
		return err
	}

	if userActivity == nil {
		if !request.IsLiked {
			return errors.New("you haven't liked before")
		}

		err = s.postRepo.CreateUserActivity(ctx, model)
	} else {

		err = s.postRepo.UpdateUserActivity(ctx, model)
	}

	if err != nil {
		log.Error().Err(err).Msg("error create or update activity to DB")
		return err
	}

	return nil
}
