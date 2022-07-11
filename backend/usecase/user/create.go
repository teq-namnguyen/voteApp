package user

import (
	"context"
	"strings"

	"github.com/namnguyen/backend/model"
	"github.com/namnguyen/backend/utils"
)

type CreateRequest struct {
	Username *string `json:"username" `
	Password *string `json:"password"`
}

func (u *UseCase) validateCreate(req *CreateRequest) error {
	*req.Username = strings.TrimSpace(*req.Username)
	if req.Username == nil || *req.Username == "" {
		return utils.ErrExampleInvalidParam("Username")
	}

	if req.Password == nil || *req.Password == "" {
		return utils.ErrExampleInvalidParam("Password")
	}

	return nil
}

func (u *UseCase) Create(ctx context.Context, req *CreateRequest) (*ResponseUserWrapper, error) {
	if err := u.validateCreate(req); err != nil {
		return nil, err
	}

	data := &model.User{
		Username: *req.Username,
		Password: *req.Password,
	}

	err := u.UserRepo.Create(ctx, data)
	if err != nil {
		return nil, utils.ErrExampleCreate(err)
	}

	return &ResponseUserWrapper{User: data}, nil
}
