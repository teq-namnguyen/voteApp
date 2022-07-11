package poll

import (
	"context"

	"github.com/namnguyen/backend/model"
)

type IUseCase interface {
	GetByID(ctx context.Context, id int) (*ResponsePollWrapper, error)
	Create(ctx context.Context, req *CreateRequest) (*ResponsePollWrapper, error)
	GetList(ctx context.Context) (*ResponseListPollOptionsWrapper, error)
	Share(ctx context.Context, req *SharePollRequest) error
	Vote(ctx context.Context, req *VoteRequest) error
	Delete(ctx context.Context, pollID int) error
}

type ResponsePollWrapper struct {
	Poll    *model.Poll    `json:"poll"`
	Options []model.Option `json:"options"`
}

type ResponseListPollOptionsWrapper struct {
	Poll []model.OptionInPoll `json:"poll"`
}
