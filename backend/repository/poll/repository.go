package poll

import (
	"context"
	"gorm.io/gorm"

	"github.com/namnguyen/backend/model"
)

type Repository interface {
	GetByID(ctx context.Context, id int) (*model.Poll, error)
	Create(ctx context.Context, poll *model.Poll) error
	Get(ctx context.Context, poll model.Poll) (*[]model.Poll, error)
	GetListByUserID(ctx context.Context, userID *int) ([]model.Poll, error)
	Delete(ctx context.Context, id int) error

	TestDelete(ctx context.Context, id int, test ...func(tx *gorm.DB, id int) error) error
	DeleteByPollID(tx *gorm.DB, id int) error
	DeleteByPollID2(tx *gorm.DB, id int) error
}
