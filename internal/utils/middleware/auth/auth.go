package auth

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"

	"AlekseyMartunov/internal/adapters/db/users/postgres"
)

const invalidToken = "Please provide valid credentials"
const emptyUser = "User does not exist"
const internalServerError = "Internal server error"

type tokenController interface {
	GetUserUUID(tokenString string) (string, error)
}

type userService interface {
	CheckUserUUID(ctx context.Context, userUUID string) error
}

type Auth struct {
	tokenController
	userService
}

func New(us userService, tk tokenController) *Auth {
	return &Auth{
		tokenController: tk,
		userService:     us,
	}
}

func (a *Auth) CheckAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		if token == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, invalidToken)
		}
		token = strings.Split(token, " ")[1]

		userUUID, err := a.tokenController.GetUserUUID(token)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, invalidToken)
		}

		err = a.userService.CheckUserUUID(c.Request().Context(), userUUID)
		if err != nil {
			if errors.Is(err, postgres.ErrUserDoseNotExist) {
				return echo.NewHTTPError(http.StatusUnauthorized, emptyUser)
			}
			return echo.NewHTTPError(http.StatusInternalServerError, internalServerError)
		}
		return next(c)
	}
}
