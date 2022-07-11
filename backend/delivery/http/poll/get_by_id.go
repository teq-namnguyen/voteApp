package poll

import (
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/namnguyen/backend/utils"
)

func (r *Route) getByID(c echo.Context) error {
	var (
		ctx    = &utils.CustomEchoContext{Context: c}
		pollID = c.Param("id")
	)

	pollIDInt, err := strconv.Atoi(pollID)
	if err != nil {
		return utils.Response.Error(ctx, utils.ErrServerInternals("convert str to int"))
	}

	resp, err := r.useCase.Poll.GetByID(ctx, pollIDInt)
	if err != nil {
		return utils.Response.Error(ctx, err.(utils.TeqError))
	}

	return utils.Response.Success(c, resp)
}
