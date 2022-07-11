package poll

import (
	"context"

	"github.com/namnguyen/backend/model"
	"github.com/namnguyen/backend/utils"
)

type VoteRequest struct {
	Options []int `json:"options"`
	PollID  int   `json:"-"`
}

func getOptionsID(option []model.Option) []int {
	result := make([]int, len(option))
	for i := range option {
		result[i] = option[i].ID
	}

	return result
}

func (u *UseCase) validateVote(ctx context.Context, req *VoteRequest) (*[]model.Option, error) {
	userID := ctx.Value(model.UserIDClaim).(float64)

	if len(req.Options) == 0 {
		return nil, utils.ErrExampleInvalidParam("Param err")
	}

	_, err := u.UserPollRepo.Get(ctx, int(userID), req.PollID)
	if err != nil {
		return nil, utils.ErrExampleInvalidParam("pollID or userID")
	}

	option, err := u.OptionRepo.GetListByIDs(ctx, req.Options)
	if err != nil {
		return nil, utils.ErrServerInternals("get DB err")
	}

	if len(option) == 0 {
		return nil, utils.ErrExampleNotFound()
	}

	req.Options = getOptionsID(option)

	return &option, nil
}

func (u *UseCase) Vote(ctx context.Context, req *VoteRequest) error {
	_, err := u.validateVote(ctx, req)
	if err != nil {
		return err
	}

	err = u.OptionRepo.UpdateOptions(ctx, req.Options)
	if err != nil {
		return utils.ErrServerInternals("update DB err")
	}

	return nil
}
