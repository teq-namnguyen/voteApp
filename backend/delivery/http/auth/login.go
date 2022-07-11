package auth

import (
	"github.com/labstack/echo/v4"

	"github.com/namnguyen/backend/usecase/auth"
	"github.com/namnguyen/backend/utils"
)

func (r *Route) login(c echo.Context) error {
	var (
		ctx = &utils.CustomEchoContext{Context: c}
		req = &auth.LoginRequest{}
	)

	if err := c.Bind(&req); err != nil {
		return utils.Response.Error(ctx, utils.ErrExampleInvalidParam(err.Error()))
	}

	resp, err := r.useCase.Auth.Login(ctx, req)
	if err != nil {
		return utils.Response.Error(ctx, err.(utils.TeqError))
	}

	return utils.Response.Success(c, resp)
}
