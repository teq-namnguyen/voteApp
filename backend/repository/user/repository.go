package user

import (
	"context"

	"github.com/namnguyen/backend/model"
)

type Repository interface {
	GetByID(ctx context.Context, id int) (*model.User, error)
	GetByUsername(ctx context.Context, username string) (*model.User, error)
	Create(ctx context.Context, data *model.User) error
	GetAll(ctx context.Context) ([]model.User, error)
	GetByListID(ctx context.Context, id []int) (*[]model.User, error)
}
