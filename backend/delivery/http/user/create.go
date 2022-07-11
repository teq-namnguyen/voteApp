package user

import (
	"github.com/labstack/echo/v4"

	"github.com/namnguyen/backend/usecase/user"
	"github.com/namnguyen/backend/utils"
)

func (r *Route) create(c echo.Context) error {
	var (
		ctx = &utils.CustomEchoContext{Context: c}
		req = user.CreateRequest{}
	)

	if err := c.Bind(&req); err != nil {
		return utils.Response.Error(ctx, utils.ErrExampleInvalidParam(err.Error()))
	}

	resp, err := r.UseCase.User.Create(ctx, &req)
	if err != nil {
		return utils.Response.Error(ctx, utils.ErrExampleInvalidParam(err.Error()))
	}

	return utils.Response.Success(c, resp)
}
