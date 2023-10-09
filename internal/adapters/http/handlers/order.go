package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) SaveOrder(c echo.Context) error {
	//userUUID := c.Get("userUUID").(string)
	//number, err := io.ReadAll(c.Request().Body)
	//if err != nil {
	//	return c.String(http.StatusBadRequest, incorrectReq)
	//}

	//err := h.orderService.Create(string(number), userUUID)
	//if err != nil {
	//	ownerUUID, err := h.orderService.GetOwner(string(number))
	//}

	return c.String(http.StatusOK, "OK")
}
