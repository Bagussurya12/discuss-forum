package memberships

import (
	"context"
	"errors"
	"time"

	"github.com/Bagussurya12/discuss-forum/pkg/jwt"
	tokenUtil "github.com/Bagussurya12/discuss-forum/pkg/token"
	"github.com/Bagussurya12/discuss-forum/source/model/memberships"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Login(ctx context.Context, req memberships.LoginRequest) (string, string, error) {

	user, err := s.membershipRepo.GetUser(ctx, req.Email, "", "")

	if err != nil {
		log.Error().Err(err).Msg("Failed To Get User")
		return "", "", err
	}

	if user == nil {
		return "", "", errors.New("email not exist")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))

	if err != nil {
		return "", "", errors.New("email or password invalid")
	}

	token, err := jwt.CreateToken(user.ID, user.Username, s.cfg.Service.SecretJWT)

	if err != nil {
		return "", "", err
	}

	existingRefreshToken, err := s.membershipRepo.GetRefreshToken(ctx, user.ID, time.Now())

	if err != nil {
		log.Error().Err(err).Msg("Failed Get Refresh Token")
		return "", "", err
	}

	if existingRefreshToken != nil {
		return token, existingRefreshToken.RefreshToken, nil
	}
	refreshToken := tokenUtil.GenerateRefreshToken()

	if refreshToken == "" {
		return token, "", errors.New("Failed Generate Refresh Token")
	}

	err = s.membershipRepo.InsertRefreshToken(ctx, memberships.RefreshTokenModel{
		UserID:       user.ID,
		RefreshToken: refreshToken,
		ExpiredAt:    time.Now().Add(3 * 24 * time.Hour),
	})
	if err != nil {
		log.Error().Err(err).Msg("error insert refresh token to db")
		return token, "", err
	}
	return token, refreshToken, nil
}
