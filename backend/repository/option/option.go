package option

import (
	"context"

	"gorm.io/gorm"

	"github.com/namnguyen/backend/model"
)

func NewPG(getDB func(ctx context.Context) *gorm.DB) Repository {
	return &pgRepository{getDB}
}

type pgRepository struct {
	getDB func(ctx context.Context) *gorm.DB
}

func (p *pgRepository) GetByID(ctx context.Context, id int) (*model.Option, error) {
	var option model.Option

	err := p.getDB(ctx).Where("id = ?", id).First(&option).Error
	if err != nil {
		return nil, err
	}

	return &option, nil
}

func (p *pgRepository) GetByPollID(ctx context.Context, id int) (*[]model.Option, error) {
	var options []model.Option

	err := p.getDB(ctx).Where("poll_id = ?", id).Find(&options).Error
	if err != nil {
		return nil, err
	}

	return &options, nil
}

func (p *pgRepository) Create(ctx context.Context, data *model.Option) error {
	return p.getDB(ctx).Create(data).Error
}

func (p *pgRepository) CreateList(ctx context.Context, data *[]model.Option) error {
	return p.getDB(ctx).CreateInBatches(data, 700).Error
}

func (p *pgRepository) GetListByPollsID(ctx context.Context, listPollID []int) ([]model.Option, error) {
	var option []model.Option

	err := p.getDB(ctx).Where("poll_id IN ?", listPollID).Find(&option).Error
	if err != nil {
		return nil, err
	}

	return option, nil
}

func (p *pgRepository) GetListByIDs(ctx context.Context, optionsID []int) ([]model.Option, error) {
	var option []model.Option

	err := p.getDB(ctx).Where("id IN ?", optionsID).Find(&option).Error
	if err != nil {
		return nil, err
	}

	return option, nil
}

func (p *pgRepository) UpdateOptions(ctx context.Context, options []int) error {
	return p.getDB(ctx).
		Model(&model.Option{}).
		Where("id IN (?)", options).
		UpdateColumn("vote", gorm.Expr(" vote + 1")).Error
}

func (p *pgRepository) DeleteByPollID(ctx context.Context, id int) error {
	return p.getDB(ctx).Where("poll_id = ?", id).Delete(&model.Option{}).Error
}
