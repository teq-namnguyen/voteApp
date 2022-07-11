package poll

import (
	"context"
	"strings"

	"github.com/namnguyen/backend/model"
	"github.com/namnguyen/backend/utils"
)

type CreateRequest struct {
	Question string   `json:"question"`
	Options  []string `json:"options"`
}

func (u *UseCase) validateCreateRequest(req *CreateRequest) error {
	options := make([]string, 0, len(req.Options))

	req.Question = strings.TrimSpace(req.Question)
	if req.Question == "" {
		return utils.ErrExampleInvalidParam("Question")
	}

	for i := range req.Options {
		req.Options[i] = strings.TrimSpace(req.Options[i])
		if req.Options[i] == "" || utils.StrArrContains(options, req.Options[i]) {
			continue
		}

		options = append(options, req.Options[i])
	}

	req.Options = options

	return nil
}

func (u *UseCase) Create(ctx context.Context, req *CreateRequest) (*ResponsePollWrapper, error) {
	userID := ctx.Value(model.UserIDClaim).(float64)

	// options
	err := u.validateCreateRequest(req)
	if err != nil {
		return nil, err
	}

	pollObj := &model.Poll{
		Question: req.Question,
		UserID:   int(userID),
	}

	err = u.PollRepo.Create(ctx, pollObj)
	if err != nil {
		return nil, utils.ErrExampleCreate(err)
	}

	optionsObj := make([]model.Option, len(req.Options))
	for i := range req.Options {
		optionsObj[i] = model.Option{Name: req.Options[i], PollID: pollObj.ID}
	}

	err = u.OptionRepo.CreateList(ctx, &optionsObj)
	if err != nil {
		return nil, utils.ErrExampleCreate(err)
	}

	pollUserObj := &model.UserPoll{PollID: pollObj.ID, UserID: pollObj.UserID}

	err = u.UserPollRepo.Create(ctx, pollUserObj)
	if err != nil {
		return nil, utils.ErrExampleCreate(err)
	}

	return &ResponsePollWrapper{
		Poll:    pollObj,
		Options: optionsObj,
	}, nil
}
