package usecase

import (
	"github.com/namnguyen/backend/repository"
	"github.com/namnguyen/backend/usecase/auth"
	"github.com/namnguyen/backend/usecase/option"
	"github.com/namnguyen/backend/usecase/poll"
	"github.com/namnguyen/backend/usecase/user"
)

type UseCase struct {
	User   user.IUseCase
	Poll   poll.IUseCase
	Auth   auth.IUseCase
	Option option.IUseCase
}

func New(repo *repository.Repository) *UseCase {
	return &UseCase{
		User:   user.New(repo),
		Poll:   poll.New(repo),
		Auth:   auth.New(repo),
		Option: option.New(repo),
	}
}
