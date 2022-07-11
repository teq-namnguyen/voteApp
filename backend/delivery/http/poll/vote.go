package poll

import (
	"github.com/labstack/echo/v4"

	"github.com/namnguyen/backend/usecase/poll"
	"github.com/namnguyen/backend/utils"
)

func (r *Route) vote(c echo.Context) error {
	ctx := &utils.CustomEchoContext{Context: c}
	req := poll.VoteRequest{}

	pollID, err := utils.StrToInt(c.Param("id"))
	if err != nil {
		return utils.Response.Error(c, utils.ErrServerInternals("convert err"))
	}

	if err = c.Bind(&req); err != nil {
		return utils.Response.Error(ctx, utils.ErrServerInternals("parse err"))
	}

	req.PollID = *pollID

	err = r.useCase.Poll.Vote(ctx, &req)
	if err != nil {
		return utils.Response.Error(c, err.(utils.TeqError))
	}

	return utils.Response.Success(c, nil)
}
