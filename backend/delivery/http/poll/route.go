package poll

import (
	"github.com/labstack/echo/v4"

	"github.com/namnguyen/backend/usecase"
	"github.com/namnguyen/backend/utils"
)

type Route struct {
	useCase *usecase.UseCase
}

func Init(group *echo.Group, useCase *usecase.UseCase) {
	r := &Route{useCase: useCase}
	group.GET("", r.getList, utils.JWTAuth())
	group.POST("", r.create, utils.JWTAuth())
	group.GET("/:id", r.getByID, utils.JWTAuth())
	group.POST("/:id/vote", r.vote, utils.JWTAuth())
	group.POST("/:id/share", r.share, utils.JWTAuth())
	group.DELETE("/:id", r.delete, utils.JWTAuth())
	group.POST("/:id/options", r.createOption, utils.JWTAuth())
}
