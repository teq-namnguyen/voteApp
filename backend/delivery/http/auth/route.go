package auth

import (
	"github.com/labstack/echo/v4"

	"github.com/namnguyen/backend/usecase"
)

type Route struct {
	useCase *usecase.UseCase
}

func Init(group *echo.Group, useCase *usecase.UseCase) {
	r := &Route{useCase}
	group.POST("/login", r.login)
}
