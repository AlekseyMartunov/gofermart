package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"

	"AlekseyMartunov/internal/adapters/db/users/postgres"
)

const invalidToken = "Please provide valid credentials"
const emptyUser = "User does not exist"
const internalServerError = "Internal server error"

type logger interface {
	Info(msg string)
	Warn(msg string)
	Error(msg string)
}

type tokenController interface {
	GetUserUUID(tokenString string) (string, error)
}

type userService interface {
	GetIDByUUID(ctx context.Context, uuid string) (int, error)
}

type Auth struct {
	tokenController tokenController
	userService     userService
	logger          logger
}

func New(us userService, tk tokenController, l logger) *Auth {
	return &Auth{
		tokenController: tk,
		userService:     us,
		logger:          l,
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
			a.logger.Error(invalidToken)
			return echo.NewHTTPError(http.StatusInternalServerError, invalidToken)
		}

		userID, err := a.userService.GetIDByUUID(c.Request().Context(), userUUID)
		if err != nil {
			a.logger.Error("не нашлость пользователя с таким uuid")
			if errors.Is(err, postgres.ErrUserDoseNotExist) {
				return echo.NewHTTPError(http.StatusUnauthorized, emptyUser)
			}
			return echo.NewHTTPError(http.StatusInternalServerError, internalServerError)
		}

		c.Set("userID", userID)

		a.logger.Info(fmt.Sprintf("Пришел пользователь с ID: %d", userID))
		return next(c)
	}
}
