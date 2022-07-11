package poll

import (
	"context"
	"github.com/namnguyen/backend/client/postgresql"
	"github.com/namnguyen/backend/model"
	"github.com/namnguyen/backend/utils"
)

func (u *UseCase) validateDeleteRequest(userID int, poll *model.Poll) error {
	if userID != poll.UserID {
		return utils.ErrExampleInvalidParam("User have no permission")
	}

	return nil
}

func (u *UseCase) Delete(ctx context.Context, pollID int) error {
	userID := ctx.Value(model.UserIDClaim).(float64)

	poll, err := u.PollRepo.GetByID(ctx, pollID)
	if err != nil {
		return utils.ErrServerInternals("Server error")
	}

	err = u.validateDeleteRequest(int(userID), poll)
	if err != nil {
		return err
	}

	ctx = utils.TxBegin(ctx, postgresql.GetClient)
	_, err = utils.TxEnd(ctx, func(ctx context.Context) error {
		err := u.OptionRepo.DeleteByPollID(ctx, pollID)
		if err != nil {
			return utils.ErrServerInternals(err.Error())
		}

		err = u.UserPollRepo.DeleteByPollID(ctx, pollID)
		if err != nil {
			return utils.ErrServerInternals(err.Error())
		}

		err = u.PollRepo.Delete(ctx, pollID)
		if err != nil {
			return utils.ErrServerInternals(err.Error())
		}

		return nil
	})
	if err != nil {
		return utils.ErrServerInternals(err.Error())
	}

	return nil
}
