package utils

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"

	"github.com/namnguyen/backend/config"
	"github.com/namnguyen/backend/model"
)

func JWTAuth() echo.MiddlewareFunc {
	config := config.GetConfig()
	c := middleware.DefaultJWTConfig

	c.SigningKey = []byte(config.SercretKey)
	c.ErrorHandlerWithContext = func(err error, c echo.Context) error {
		httpErr, ok := err.(*echo.HTTPError)
		if ok {
			if httpErr.Code == http.StatusBadRequest || httpErr.Code == 401 {
				return Response.Error(c, ErrToken(err.Error()))
			}
		}

		return Response.Error(c, ErrToken(err.Error()))
	}
	// c.TokenLookup = "query:token"
	c.ParseTokenFunc = func(auth string, c echo.Context) (interface{}, error) {
		keyFunc := func(t *jwt.Token) (interface{}, error) {
			if t.Method.Alg() != "HS256" {
				return nil, errors.New("Error token")
			}

			return []byte(config.SercretKey), nil
		}

		// claims are of type `jwt.MapClaims` when token is created with `jwt.Parse`
		token, err := jwt.Parse(auth, keyFunc)
		if err != nil {
			return nil, err
		}

		if !token.Valid {
			return nil, errors.New("Invalid token")
		}

		// get user

		c.Set(model.UsernameClaim, token.Claims.(jwt.MapClaims)["name"])
		c.Set(model.UserIDClaim, token.Claims.(jwt.MapClaims)["user_id"])

		return token, nil
	}

	return middleware.JWTWithConfig(c)
}

type jwtCustomClaims struct {
	jwt.StandardClaims
	Name   string `json:"name"`
	UserID int    `json:"user_id"`
}

func CreateToken(username string, userID int) (*string, error) {
	cfg := config.GetConfig()
	claims := &jwtCustomClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
		username,
		userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(cfg.SercretKey))
	if err != nil {
		return nil, err
	}

	return &t, nil
}
