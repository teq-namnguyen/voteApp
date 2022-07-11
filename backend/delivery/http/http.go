package http

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/namnguyen/backend/delivery/http/auth"
	"github.com/namnguyen/backend/delivery/http/poll"
	"github.com/namnguyen/backend/delivery/http/user"
	"github.com/namnguyen/backend/usecase"
)

func NewHTTPHandler(useCase *usecase.UseCase) *echo.Echo {
	e := echo.New()
	api := e.Group("/api")
	api.Use(middleware.Logger())
	user.Init(api.Group("/users"), useCase)
	auth.Init(api.Group("/auth"), useCase)
	poll.Init(api.Group("/polls"), useCase)

	return e
}
