package loginhandlers

import (
	"AlekseyMartunov/internal/adapters/db/users/postgres"
	"encoding/json"
	"errors"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
)

func (lh *LoginHandler) Register(c echo.Context) error {
	defer c.Request().Body.Close()
	c.Response().Header().Set("Content-Type", "application/json")

	b, err := io.ReadAll(c.Request().Body)
	if err != nil {
		lh.logger.Error("request body read error")
		return c.String(http.StatusBadRequest, incorrectReq)
	}

	user := userDTO{}

	err = json.Unmarshal(b, &user)
	if err != nil {
		lh.logger.Error("unmarshal error")
		return c.String(http.StatusBadRequest, incorrectReq)
	}

	if user.Login == "" || user.Password == "" {
		return c.String(http.StatusBadRequest, incorrectReq)
	}

	userUUID, err := lh.userService.Create(c.Request().Context(), user.toEntity())
	if err != nil {
		if errors.Is(err, postgres.ErrLoginAlreadyUsed) {
			return c.String(http.StatusConflict, loginAlreadyExist)
		}
		return c.String(http.StatusInternalServerError, internalErr)
	}

	token, err := lh.tokenManager.CreateToken(userUUID)
	if err != nil {
		return c.String(http.StatusInternalServerError, internalErr)
	}

	c.Response().Header().Add("Authorization", "Bearer "+token)
	return c.String(http.StatusOK, ok)
}
