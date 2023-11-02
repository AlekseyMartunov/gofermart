package orderhandlers

import (
	"AlekseyMartunov/internal/orders"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *OrderHandler) GetOrders(c echo.Context) error {
	userID := c.Get("userID").(int)
	res, err := h.orderService.GetOrders(c.Request().Context(), userID)

	if errors.Is(err, orders.ErrEmptyResult) {
		return c.JSON(http.StatusNoContent, noContent)
	}

	if err != nil {
		h.logger.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, internalErr)
	}

	return c.JSON(http.StatusOK, fromEntity(res))
}
