package userhandlers

import (
	"AlekseyMartunov/internal/adapters/db/orders/postgres"
	"AlekseyMartunov/internal/orders"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *UserHandlers) Withdraw(c echo.Context) error {
	defer c.Request().Body.Close()
	userID := c.Get("userID").(int)
	ctx := c.Request().Context()

	b, err := io.ReadAll(c.Request().Body)
	if err != nil {
		h.logger.Error("request body read error")
		return c.String(http.StatusInternalServerError, internalErr)
	}

	o := orderDTO{UserID: userID}
	err = json.Unmarshal(b, &o)
	if err != nil {
		h.logger.Error("Unmarshal error")
		return c.String(http.StatusBadRequest, incorrectReq)
	}

	err = h.orderService.AddDiscount(ctx, o.toEntity())
	if err != nil {
		if errors.Is(err, orders.ErrNotValidNumber) {
			return c.JSON(http.StatusUnprocessableEntity, notValidOrderNumber)
		}

		if errors.Is(err, postgres.ErrNotEnoughMoney) {
			return c.JSON(http.StatusPaymentRequired, notEnoughMoney)
		}
		h.logger.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, internalErr)
	}

	return c.String(http.StatusOK, ok)
}
