package option

import (
	"context"
	"github.com/namnguyen/backend/model"
	"github.com/namnguyen/backend/utils"
	"strings"
)

type CreateOptionRequest struct {
	Name   string `json:"name"`
	PollID int    `json:"-"`
}

func (u *UseCase) Create(ctx context.Context, req *CreateOptionRequest) (*ResponseOptionWrapper, error) {
	userID := ctx.Value(model.UserIDClaim).(float64)

	req.Name = strings.TrimSpace(req.Name)
	if req.Name == "" {
		return nil, utils.ErrExampleInvalidParam("Name")
	}

	_, err := u.UserPollRepo.Get(ctx, int(userID), req.PollID)
	if err != nil {
		return nil, utils.ErrServerInternals("Not have Permission")
	}

	option := model.Option{Name: req.Name, PollID: req.PollID}

	err = u.OptionRepo.Create(ctx, &option)
	if err != nil {
		return nil, utils.ErrServerInternals("Create Error")
	}

	return &ResponseOptionWrapper{
		Options: &option,
	}, nil
}
