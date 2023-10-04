package handlers

import (
	"AlekseyMartunov/internal/adapters/db/users/postgres"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Register(c echo.Context) error {
	defer c.Request().Body.Close()

	b, err := io.ReadAll(c.Request().Body)
	if err != nil {
		h.logger.Error("request body read error")
		return c.String(http.StatusBadRequest, incorrectReq)
	}

	user := userDTO{}

	err = json.Unmarshal(b, &user)
	if err != nil {
		h.logger.Error("unmarshal error")
		return c.String(http.StatusBadRequest, incorrectReq)
	}

	if user.Login == "" || user.Password == "" {
		return c.String(http.StatusBadRequest, incorrectReq)
	}

	err = h.userService.Register(context.Background(), user.Login, user.Password)
	if err != nil {
		if errors.Is(err, postgres.LoginAlreadyUsedErr) {
			return c.String(http.StatusConflict, loginAlreadyExist)
		}
		return c.String(http.StatusInternalServerError, internalErr)
	}
	return c.String(http.StatusOK, ok)
}
