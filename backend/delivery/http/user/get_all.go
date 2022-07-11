package user

import (
	"github.com/labstack/echo/v4"

	"github.com/namnguyen/backend/utils"
)

func (r *Route) getAll(c echo.Context) error {
	ctx := &utils.CustomEchoContext{Context: c}

	resp, err := r.UseCase.User.GetAll(ctx)
	if err != nil {
		return utils.Response.Error(ctx, err.(utils.TeqError))
	}

	return utils.Response.Success(ctx, resp)
}
