package user

import (
	"context"

	"github.com/namnguyen/backend/model"
)

type IUseCase interface {
	GetByUsername(ctx context.Context, username string) (*ResponseUserWrapper, error)
	Create(ctx context.Context, req *CreateRequest) (*ResponseUserWrapper, error)
	GetAll(ctx context.Context) (*ResponseUsersWrapper, error)
}

type ResponseUserWrapper struct {
	User *model.User `json:"user"`
}

type ResponseUsersWrapper struct {
	User *[]model.User `json:"user"`
}
