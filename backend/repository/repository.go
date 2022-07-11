package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/namnguyen/backend/repository/option"
	"github.com/namnguyen/backend/repository/poll"
	"github.com/namnguyen/backend/repository/user"
	"github.com/namnguyen/backend/repository/userpoll"
)

type Repository struct {
	User     user.Repository
	Poll     poll.Repository
	Option   option.Repository
	UserPoll userpoll.Repository
}

func New(getClient func(ctx context.Context) *gorm.DB) *Repository {
	return &Repository{
		User:     user.NewPG(getClient),
		Poll:     poll.NewPG(getClient),
		Option:   option.NewPG(getClient),
		UserPoll: userpoll.NewPG(getClient),
	}
}
