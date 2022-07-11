package user

import (
	"github.com/labstack/echo/v4"

	"github.com/namnguyen/backend/utils"
)

func (r *Route) getByUsername(c echo.Context) error {
	var (
		ctx      = &utils.CustomEchoContext{Context: c}
		userName = c.Param("username")
	)

	resp, err := r.UseCase.User.GetByUsername(ctx, userName)
	if err != nil {
		return utils.Response.Error(ctx, err.(utils.TeqError))
	}

	return utils.Response.Success(c, resp)
}
