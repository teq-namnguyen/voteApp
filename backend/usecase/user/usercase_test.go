package user_test

import (
	"context"
	userRepo "github.com/namnguyen/backend/repository/user"
	userUC "github.com/namnguyen/backend/usecase/user"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TestSuite struct {
	suite.Suite

	ctx     context.Context
	useCase *userUC.UseCase

	mockUserRepo *userRepo.MockRepository
}

func (suite *TestSuite) SetupTest() {
	suite.ctx = context.Background()
	suite.mockUserRepo = &userRepo.MockRepository{}
	suite.useCase = &userUC.UseCase{
		UserRepo: suite.mockUserRepo,
	}
}

func TestUseCaseUser(t *testing.T) {
	t.Parallel()
	suite.Run(t, &TestSuite{})
}
