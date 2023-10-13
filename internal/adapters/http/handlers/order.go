package handlers

import (
	"AlekseyMartunov/internal/adapters/db/orders/postgres"
	"AlekseyMartunov/internal/orders"
	"errors"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
)

func (h *Handler) SaveOrder(c echo.Context) error {
	number, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.String(http.StatusBadRequest, incorrectReq)
	}

	order := orderDTO{
		number: string(number),
		UserID: c.Get("userID").(int),
	}

	ctx := c.Request().Context()

	err = h.orderService.Create(ctx, order.toEntity())
	if err != nil {
		if errors.Is(err, orders.ErrNotValidNumber) {
			return c.String(http.StatusUnprocessableEntity, notValidOrderNumber)
		}
		if errors.Is(err, postgres.ErrOrderAlreadyCreated) {
			id, err := h.orderService.GetUserID(ctx, order.number)
			if err != nil {
				return c.String(http.StatusInternalServerError, internalErr)
			}
			if id == order.UserID {
				return c.String(http.StatusOK, ok)
			} else {
				return c.String(http.StatusConflict, orderRegisteredByAnotherUser)
			}
		}
	}

	return c.String(http.StatusAccepted, ok)
}
