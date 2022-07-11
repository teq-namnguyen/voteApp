package auth

import (
	"context"
)

type IUseCase interface {
	Login(ctx context.Context, input *LoginRequest) (*ResponseTokenWrapper, error)
}

type ResponseTokenWrapper struct {
	Token string `json:"token"`
}
