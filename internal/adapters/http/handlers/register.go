package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"

	"AlekseyMartunov/internal/adapters/db/users/postgres"
)

func (h *Handler) Register(c echo.Context) error {
	defer c.Request().Body.Close()
	c.Response().Header().Set("Content-Type", "application/json")

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

	userUUID, err := h.userService.Create(c.Request().Context(), user.Login, user.Password)
	if err != nil {
		if errors.Is(err, postgres.ErrLoginAlreadyUsed) {
			return c.String(http.StatusConflict, loginAlreadyExist)
		}
		return c.String(http.StatusInternalServerError, internalErr)
	}

	token, err := h.tokenManager.CreateToken(userUUID)
	if err != nil {
		return c.String(http.StatusInternalServerError, internalErr)
	}

	c.Response().Header().Add("Authorization", "Bearer "+token)
	return c.String(http.StatusOK, ok)
}
