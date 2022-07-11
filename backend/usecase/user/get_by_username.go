package user

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/namnguyen/backend/utils"
)

func (u *UseCase) GetByUsername(ctx context.Context, username string) (*ResponseUserWrapper, error) {
	user, err := u.UserRepo.GetByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.ErrExampleNotFound()
		}

		return nil, utils.ErrExampleGet(err)
	}

	return &ResponseUserWrapper{User: user}, nil
}
