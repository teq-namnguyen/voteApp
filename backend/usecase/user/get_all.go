package user

import (
	"context"

	"github.com/namnguyen/backend/utils"
)

func (u *UseCase) GetAll(ctx context.Context) (*ResponseUsersWrapper, error) {
	users, err := u.UserRepo.GetAll(ctx)
	if err != nil {
		return nil, utils.ErrServerInternals(err.Error())
	}

	return &ResponseUsersWrapper{
		User: &users,
	}, nil
}
