package poll

import (
	"github.com/labstack/echo/v4"

	"github.com/namnguyen/backend/usecase/poll"
	"github.com/namnguyen/backend/utils"
)

func (r *Route) create(c echo.Context) error {
	ctx := &utils.CustomEchoContext{Context: c}
	req := &poll.CreateRequest{}

	if err := c.Bind(req); err != nil {
		return utils.Response.Error(ctx, utils.ErrExampleInvalidParam(err.Error()))
	}

	resp, err := r.useCase.Poll.Create(ctx, req)
	if err != nil {
		return utils.Response.Error(ctx, err.(utils.TeqError))
	}

	return utils.Response.Success(ctx, resp)
}
