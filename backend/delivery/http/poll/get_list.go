package poll

import (
	"github.com/labstack/echo/v4"

	"github.com/namnguyen/backend/utils"
)

func (r *Route) getList(c echo.Context) error {
	ctx := &utils.CustomEchoContext{Context: c}

	resp, err := r.useCase.Poll.GetList(ctx)
	if err != nil {
		return utils.Response.Error(ctx, err.(utils.TeqError))
	}

	return utils.Response.Success(c, resp)
}
