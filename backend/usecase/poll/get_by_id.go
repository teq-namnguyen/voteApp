package poll

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/namnguyen/backend/utils"
)

func (u *UseCase) GetByID(ctx context.Context, id int) (*ResponsePollWrapper, error) {
	poll, err := u.PollRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.ErrExampleNotFound()
		}

		return nil, utils.ErrExampleGet(err)
	}

	return &ResponsePollWrapper{Poll: poll}, nil
}
