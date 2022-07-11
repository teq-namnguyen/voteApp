package user

import (
	"github.com/labstack/echo/v4"

	"github.com/namnguyen/backend/usecase"
	"github.com/namnguyen/backend/utils"
)

type Route struct {
	UseCase *usecase.UseCase
}

func Init(group *echo.Group, useCase *usecase.UseCase) {
	r := &Route{useCase}

	group.GET("", r.getAll, utils.JWTAuth())
	group.POST("/create", r.create)
	group.GET("/:username", r.getByUsername, utils.JWTAuth())
}
