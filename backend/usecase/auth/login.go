package auth

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/namnguyen/backend/utils"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *UseCase) Login(ctx context.Context, input *LoginRequest) (*ResponseTokenWrapper, error) {
	user, err := u.UserRepo.GetByUsername(ctx, input.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.ErrExampleNotFound()
		}

		return nil, utils.ErrExampleGet(err)
	}

	if input.Password != user.Password {
		return nil, utils.ErrGetWithString()
	}

	token, err := utils.CreateToken(user.Username, user.ID)
	if err != nil {
		return nil, utils.ErrGetWithString()
	}

	return &ResponseTokenWrapper{Token: *token}, nil
}
