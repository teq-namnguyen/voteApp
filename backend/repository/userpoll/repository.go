package userpoll

import (
	"context"

	"github.com/namnguyen/backend/model"
)

type Repository interface {
	Create(ctx context.Context, userPoll *model.UserPoll) error
	GetList(ctx context.Context, userID int) (*[]model.UserPoll, error)
	CreateList(ctx context.Context, userPolls *[]model.UserPoll) error
	CreateOrUpdateList(ctx context.Context, userPolls *[]model.UserPoll) error
	Get(ctx context.Context, userID int, pollID int) (*model.UserPoll, error)
	DeleteByPollID(ctx context.Context, id int) error
}
