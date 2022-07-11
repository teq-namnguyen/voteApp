package userpoll

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/namnguyen/backend/model"
)

func NewPG(getDB func(ctx context.Context) *gorm.DB) Repository {
	return &pgRepository{getDB}
}

type pgRepository struct {
	getDB func(ctx context.Context) *gorm.DB
}

func (p *pgRepository) Create(ctx context.Context, userPoll *model.UserPoll) error {
	return p.getDB(ctx).Create(userPoll).Error
}

func (p *pgRepository) GetList(ctx context.Context, userID int) (*[]model.UserPoll, error) {
	var resp *[]model.UserPoll

	err := p.getDB(ctx).Where("user_id = ?", userID).Find(&resp).Error
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *pgRepository) CreateList(ctx context.Context, userPoll *[]model.UserPoll) error {
	return p.getDB(ctx).CreateInBatches(userPoll, 500).Error
}

func (p *pgRepository) CreateOrUpdateList(ctx context.Context, userPoll *[]model.UserPoll) error {
	return p.getDB(ctx).
		Clauses(clause.OnConflict{
			DoNothing: true,
		}).
		Create(userPoll).
		Error
}

func (p *pgRepository) Get(ctx context.Context, userID int, pollID int) (*model.UserPoll, error) {
	var resp model.UserPoll

	err := p.getDB(ctx).Where("user_id = ? AND poll_id = ?", userID, pollID).First(&resp).Error
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (p *pgRepository) DeleteByPollID(ctx context.Context, id int) error {
	return p.getDB(ctx).Where("poll_id = ?", id).Delete(&model.UserPoll{}).Error
}
