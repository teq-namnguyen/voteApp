package poll

import (
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/namnguyen/backend/usecase/poll"
	"github.com/namnguyen/backend/utils"
)

func (r *Route) share(c echo.Context) error {
	ctx := &utils.CustomEchoContext{Context: c}
	req := poll.SharePollRequest{}

	err := c.Bind(&req)
	if err != nil {
		return utils.Response.Error(ctx, utils.ErrServerInternals("bind json err"))
	}

	pollID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return utils.Response.Error(ctx, utils.ErrServerInternals("convert str to int"))
	}

	req.PollID = pollID

	err = r.useCase.Poll.Share(ctx, &req)
	if err != nil {
		return utils.Response.Error(ctx, err.(utils.TeqError))
	}

	return utils.Response.Success(ctx, nil)
}
