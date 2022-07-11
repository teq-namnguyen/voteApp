package poll

import (
	"github.com/labstack/echo/v4"

	"github.com/namnguyen/backend/utils"
)

func (r *Route) delete(c echo.Context) error {
	ctx := &utils.CustomEchoContext{Context: c}

	pollID, err := utils.StrToInt(c.Param("id"))
	if err != nil {
		return utils.Response.Error(ctx, utils.ErrServerInternals("parse error"))
	}

	err = r.useCase.Poll.Delete(ctx, *pollID)
	if err != nil {
		return utils.Response.Error(ctx, err.(utils.TeqError))
	}

	return utils.Response.Success(ctx, nil)
}
