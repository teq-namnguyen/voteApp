package poll

import (
	"context"

	"github.com/namnguyen/backend/model"
	"github.com/namnguyen/backend/utils"
)

func (u *UseCase) GetList(ctx context.Context) (*ResponseListPollOptionsWrapper, error) {
	userID := int(ctx.Value(model.UserIDClaim).(float64))

	pollObj, err := u.PollRepo.GetListByUserID(ctx, &userID)
	if err != nil {
		return nil, utils.ErrServerInternals("getListByUserID")
	}

	if pollObj == nil {
		return nil, utils.ErrServerInternals("No polls created")
	}

	pollIDs := make([]int, len(pollObj))

	for i := range pollObj {
		pollIDs[i] = pollObj[i].ID
	}

	optionsObj, err := u.OptionRepo.GetListByPollsID(ctx, pollIDs)
	if err != nil {
		return nil, utils.ErrServerInternals("get options Err")
	}

	resp := createResponse(pollObj, optionsObj)

	return &ResponseListPollOptionsWrapper{
		Poll: resp,
	}, nil
}

func createResponse(polls []model.Poll, options []model.Option) []model.OptionInPoll {
	groupOptions := make(map[int][]model.Option)

	for i := range options {
		groupOptions[options[i].PollID] = append(groupOptions[options[i].PollID], options[i])
	}

	resp := make([]model.OptionInPoll, len(polls))

	for i := range polls {
		opts, ok := groupOptions[polls[i].ID]
		if !ok {
			opts = []model.Option{}
		}

		resp[i] = model.OptionInPoll{Poll: polls[i], Options: opts}
	}

	return resp
}
