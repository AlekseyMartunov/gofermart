package userhandlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *UserHandlers) Balance(c echo.Context) error {
	ctx := c.Request().Context()
	userID := c.Get("userID").(int)

	u, err := h.userService.Balance(ctx, userID)
	if err != nil {
		h.logger.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "some erorororor")
	}

	user := userDTO{}
	user.fromEntity(u)

	return c.JSON(http.StatusOK, user)
}
