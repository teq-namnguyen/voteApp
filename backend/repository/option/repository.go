package option

import (
	"context"

	"github.com/namnguyen/backend/model"
)

type Repository interface {
	GetByID(ctx context.Context, id int) (*model.Option, error)
	GetByPollID(ctx context.Context, pollID int) (*[]model.Option, error)
	Create(ctx context.Context, data *model.Option) error
	CreateList(ctx context.Context, data *[]model.Option) error
	GetListByPollsID(ctx context.Context, listPollID []int) ([]model.Option, error)
	UpdateOptions(ctx context.Context, options []int) error
	GetListByIDs(ctx context.Context, optionsID []int) ([]model.Option, error)
	DeleteByPollID(ctx context.Context, id int) error
}
