package orderhandlers

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"

	"AlekseyMartunov/internal/adapters/db/orders/postgres"
)

func (h *OrderHandler) GetOrders(c echo.Context) error {
	userID := c.Get("userID").(int)
	res, err := h.orderService.GetOrders(c.Request().Context(), userID)

	if errors.Is(err, postgres.ErrEmptyResult) {
		return c.JSON(http.StatusNoContent, noContent)
	}

	if err != nil {
		h.logger.Error(err.Error())
		return c.String(http.StatusInternalServerError, internalErr)
	}

	return c.JSON(http.StatusOK, fromEntity(res))
}
