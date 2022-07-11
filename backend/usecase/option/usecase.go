package option

import (
	"github.com/namnguyen/backend/repository"
	"github.com/namnguyen/backend/repository/option"
	"github.com/namnguyen/backend/repository/userpoll"
)

type UseCase struct {
	OptionRepo   option.Repository
	UserPollRepo userpoll.Repository
}

func New(repo *repository.Repository) IUseCase {
	return &UseCase{
		OptionRepo:   repo.Option,
		UserPollRepo: repo.UserPoll,
	}
}
