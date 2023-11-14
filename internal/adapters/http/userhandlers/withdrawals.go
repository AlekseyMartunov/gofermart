package userhandlers

import (
	"AlekseyMartunov/internal/adapters/db/users/postgres"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *UserHandlers) Withdrawals(c echo.Context) error {
	userID := c.Get("userID").(int)
	res, err := h.userService.GetHistory(c.Request().Context(), userID)

	if err != nil {
		if errors.Is(err, postgres.ErrEmptyHistory) {
			return c.JSON(http.StatusNoContent, emptyHistory)
		}
		h.logger.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, internalErr)
	}
	return c.JSON(http.StatusOK, res)
}
