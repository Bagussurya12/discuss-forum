package memberships

import (
	"context"
	"errors"
	"time"

	"github.com/Bagussurya12/discuss-forum/pkg/jwt"
	"github.com/Bagussurya12/discuss-forum/source/model/memberships"
	"github.com/rs/zerolog/log"
)

func (s *service) ValidateRefreshToken(ctx context.Context, userID int64, request memberships.RefreshTokenRequest) (string, error) {
	existingRefreshToken, err := s.membershipRepo.GetRefreshToken(ctx, userID, time.Now())

	if err != nil {
		log.Error().Err(err).Msg("error get refresh token")
		return "", err
	}

	if existingRefreshToken == nil {
		return "", errors.New("refresh token has expired")
	}
	if existingRefreshToken.RefreshToken != request.Token {
		return "", errors.New("refresh token is invalid")
	}
	user, err := s.membershipRepo.GetUser(ctx, "", "", "", userID)

	if err != nil {
		log.Error().Err(err).Msg("Failed To Get User")
		return "", err
	}

	if user == nil {
		return "", errors.New("user not exist")
	}

	token, err := jwt.CreateToken(user.ID, user.Username, s.cfg.Service.SecretJWT)

	if err != nil {
		return "", err
	}

	return token, nil
}
