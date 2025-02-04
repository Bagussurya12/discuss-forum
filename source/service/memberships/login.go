package memberships

import (
	"context"
	"errors"

	"github.com/Bagussurya12/discuss-forum/pkg/jwt"
	"github.com/Bagussurya12/discuss-forum/source/model/memberships"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Login(ctx context.Context, req memberships.LoginRequest) (string, error) {

	user, err := s.membershipRepo.GetUser(ctx, req.Email, "", "")

	if err != nil {
		log.Error().Err(err).Msg("Failed To Get User")
		return "", err
	}

	if user == nil {
		return "", errors.New("Email Not Exist")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))

	if err != nil {
		return "", errors.New("Email Or Password Invalid")
	}

	token, err := jwt.CreateToken(user.ID, user.Username, s.cfg.Service.SecretJWT)

	if err != nil {
		return "", err
	}

	return token, nil
}
