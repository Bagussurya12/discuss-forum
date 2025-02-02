package memberships

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/Bagussurya12/discuss-forum/source/model/memberships"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) SignUp(ctx context.Context, req memberships.SignUpRequest) error {

	user, err := s.membershipRepo.GetUser(ctx, req.Email, req.Username)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	if user != nil {
		return errors.New("UserName Already Exist")
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	now := time.Now()
	model := memberships.UserModel{
		Email:       req.Email,
		Password:    string(pass),
		Username:    req.Username,
		CreatedAt:   now,
		UpdatedAt:   now,
		CreatedBy:   req.Username,
		UpdatedBy:   req.Username,
		PhoneNumber: req.PhoneNumber,
	}

	return s.membershipRepo.CreateUser(ctx, model)

}
