package poll

import (
	"context"
	"github.com/namnguyen/backend/model"
	"github.com/namnguyen/backend/utils"
)

type SharePollRequest struct {
	PollID int   `json:"poll_id"`
	Users  []int `json:"users"`
}

func unique(intSlice []int) []int {
	keys := make(map[int]bool)

	var list []int

	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true

			list = append(list, entry)
		}
	}

	return list
}

func getUserIDsFromUsers(users *[]model.User) *[]int {
	result := make([]int, len(*users))
	for i, data := range *users {
		result[i] = data.ID
	}
	return &result
}

func (u *UseCase) validateShare(ctx context.Context, req *SharePollRequest) error {
	userID := ctx.Value(model.UserIDClaim).(float64)

	_, err := u.UserPollRepo.Get(ctx, int(userID), req.PollID)
	if err != nil {
		return utils.ErrExampleInvalidParam("pollID or userID")
	}

	if req.Users == nil {
		return utils.ErrExampleInvalidParam("userID")
	}

	req.Users = unique(req.Users)

	users, err := u.UserRepo.GetByListID(ctx, req.Users)
	if err != nil {
		return utils.ErrExampleInvalidParam("userID")
	}

	req.Users = *getUserIDsFromUsers(users)

	return nil
}

func (u *UseCase) Share(ctx context.Context, req *SharePollRequest) error {
	err := u.validateShare(ctx, req)
	if err != nil {
		return err
	}

	userPollObj := make([]model.UserPoll, len(req.Users))
	for i := range req.Users {
		userPollObj[i] = model.UserPoll{PollID: req.PollID, UserID: req.Users[i]}
	}

	err = u.UserPollRepo.CreateOrUpdateList(ctx, &userPollObj)
	if err != nil {
		return utils.ErrServerInternals(err.Error())
	}

	return nil
}
