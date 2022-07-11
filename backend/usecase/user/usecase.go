package user

import (
	"github.com/namnguyen/backend/repository"
	"github.com/namnguyen/backend/repository/user"
)

type UseCase struct {
	UserRepo user.Repository
}

func New(repo *repository.Repository) IUseCase {
	return &UseCase{
		UserRepo: repo.User,
	}
}
