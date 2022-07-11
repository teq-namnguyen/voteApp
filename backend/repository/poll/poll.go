package poll

import (
	"context"

	"gorm.io/gorm"

	"github.com/namnguyen/backend/model"
)

func NewPG(getDB func(ctx context.Context) *gorm.DB) Repository {
	return &pgRepository{getDB}
}

type pgRepository struct {
	GetDB func(ctx context.Context) *gorm.DB
}

func (p *pgRepository) GetByID(ctx context.Context, id int) (*model.Poll, error) {
	var poll model.Poll

	err := p.GetDB(ctx).Where("id = ?", id).First(&poll).Error
	if err != nil {
		return nil, err
	}

	return &poll, nil
}

func (p *pgRepository) Create(ctx context.Context, poll *model.Poll) error {
	return p.GetDB(ctx).Create(poll).Error
}

func (p *pgRepository) Get(ctx context.Context, poll model.Poll) (*[]model.Poll, error) {
	var pollResult []model.Poll

	err := p.GetDB(ctx).Where(&poll).Find(&pollResult).Error
	if err != nil {
		return nil, err
	}

	return &pollResult, nil
}

func (p *pgRepository) GetListByUserID(ctx context.Context, userID *int) ([]model.Poll, error) {
	var pollObj []model.Poll

	subQuery := p.GetDB(ctx).Table("user_polls").Where("user_id = ?", *userID).Select("poll_id")

	err := p.GetDB(ctx).Where("id IN (?)", subQuery).Find(&pollObj).Error
	if err != nil {
		return nil, err
	}

	return pollObj, nil
}

func (p *pgRepository) Delete(ctx context.Context, id int) error {
	return p.GetDB(ctx).Delete(&model.Poll{}, id).Error
}

func (p *pgRepository) DeleteByPollID(tx *gorm.DB, id int) error {
	return tx.Where("poll_id = ?", id).Delete(&model.Option{}).Error
}

func (p *pgRepository) DeleteByPollID2(tx *gorm.DB, id int) error {
	return tx.Where("poll_Mid = ?", id).Delete(&model.UserPoll{}).Error
}

func (p *pgRepository) TestDelete(ctx context.Context, id int, test ...func(tx *gorm.DB, id int) error) error {
	return p.GetDB(ctx).Transaction(func(tx *gorm.DB) error {
		for i := range test {
			err := test[i](tx, id)
			if err != nil {
				return err
			}
		}

		err := tx.Delete(&model.Poll{}, id).Error
		if err != nil {
			return err
		}

		return nil
	})
}
