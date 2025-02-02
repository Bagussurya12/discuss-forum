package memberships

import (
	"context"

	"github.com/Bagussurya12/discuss-forum/source/model/memberships"
)

type membershipRepository interface {
	GetUser(ctx context.Context, email, username, phone_number string) (*memberships.UserModel, error)
	CreateUser(ctx context.Context, model memberships.UserModel) error
}

type service struct {
	membershipRepo membershipRepository
}

func NewService(membershipRepo membershipRepository) *service {
	return &service{
		membershipRepo: membershipRepo,
	}
}
