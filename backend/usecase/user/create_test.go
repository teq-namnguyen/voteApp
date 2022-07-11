package user_test

import (
	"github.com/namnguyen/backend/model"
	"github.com/namnguyen/backend/usecase/user"
)

func (suite *TestSuite) TestCreate_Success() {
	username := "nam123"
	password := "1234"
	req := user.CreateRequest{
		Username: &username,
		Password: &password,
	}

	mockUser := model.User{
		Username: *req.Username,
		Password: *req.Password,
	}

	suite.mockUserRepo.On("Create", suite.ctx, &mockUser).Return(nil)

	_, err := suite.useCase.Create(suite.ctx, &req)

	suite.Nil(err)
}
