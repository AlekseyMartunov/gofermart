package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type handler struct {
}

func New() *handler {
	return &handler{}
}

func (h *handler) Register(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func (h *handler) Login(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func (h *handler) GetOrders(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func (h *handler) SaveOrder(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func (h *handler) Balance(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func (h *handler) Withdraw(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func (h *handler) Withdrawals(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
