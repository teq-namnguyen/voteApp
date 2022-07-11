package poll

import (
	"github.com/labstack/echo/v4"
	"github.com/namnguyen/backend/usecase/option"
	"github.com/namnguyen/backend/utils"
)

func (r *Route) createOption(c echo.Context) error {
	ctx := &utils.CustomEchoContext{Context: c}
	pollID, err := utils.StrToInt(c.Param("id"))
	if err != nil {
		return utils.Response.Error(ctx, utils.ErrServerInternals(err.Error()))
	}

	var req option.CreateOptionRequest

	err = c.Bind(&req)
	if err != nil {
		return utils.Response.Error(ctx, utils.ErrServerInternals(err.Error()))
	}

	req.PollID = *pollID
	resp, err := r.useCase.Option.Create(ctx, &req)
	if err != nil {
		return utils.Response.Error(ctx, err.(utils.TeqError))
	}

	return utils.Response.Success(ctx, resp)
}
