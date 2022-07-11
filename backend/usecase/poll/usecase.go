package poll

import (
	"github.com/namnguyen/backend/repository"
	"github.com/namnguyen/backend/repository/option"
	"github.com/namnguyen/backend/repository/poll"
	"github.com/namnguyen/backend/repository/user"
	"github.com/namnguyen/backend/repository/userpoll"
)

type UseCase struct {
	PollRepo     poll.Repository
	OptionRepo   option.Repository
	UserPollRepo userpoll.Repository
	UserRepo     user.Repository
}

func New(repo *repository.Repository) IUseCase {
	return &UseCase{
		PollRepo:     repo.Poll,
		OptionRepo:   repo.Option,
		UserPollRepo: repo.UserPoll,
		UserRepo:     repo.User,
	}
}
