package option

import (
	"context"
	"github.com/namnguyen/backend/model"
)

type IUseCase interface {
	Create(ctx context.Context, options *CreateOptionRequest) (*ResponseOptionWrapper, error)
}

type ResponseOptionWrapper struct {
	Options *model.Option `json:"options"`
}
