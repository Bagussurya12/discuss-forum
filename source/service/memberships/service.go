package memberships

import (
	"context"
	"time"

	"github.com/Bagussurya12/discuss-forum/source/configs"
	"github.com/Bagussurya12/discuss-forum/source/model/memberships"
)

type membershipRepository interface {
	GetUser(ctx context.Context, email, username, phone_number string, userID int64) (*memberships.UserModel, error)
	CreateUser(ctx context.Context, model memberships.UserModel) error
	GetRefreshToken(ctx context.Context, userID int64, now time.Time) (*memberships.RefreshTokenModel, error)
	InsertRefreshToken(ctx context.Context, model memberships.RefreshTokenModel) error
}

type service struct {
	cfg            *configs.Config
	membershipRepo membershipRepository
}

func NewService(cfg *configs.Config, membershipRepo membershipRepository) *service {
	return &service{
		cfg:            cfg,
		membershipRepo: membershipRepo,
	}
}
