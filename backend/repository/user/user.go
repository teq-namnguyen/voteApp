package user

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

func (p *pgRepository) GetByID(ctx context.Context, id int) (*model.User, error) {
	var data model.User

	err := p.getDB(ctx).Where("id = ?", id).First(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (p *pgRepository) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	var data model.User

	err := p.getDB(ctx).
		Where("username = ?", username).
		First(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (p *pgRepository) Create(ctx context.Context, data *model.User) error {
	return p.getDB(ctx).Create(data).Error
}

func (p *pgRepository) GetAll(ctx context.Context) ([]model.User, error) {
	var user []model.User

	err := p.getDB(ctx).Find(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (p *pgRepository) GetByListID(ctx context.Context, id []int) (*[]model.User, error) {
	var user []model.User

	err := p.getDB(ctx).Where("id in (?)", id).Find(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
