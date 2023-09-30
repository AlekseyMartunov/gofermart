package router

import "github.com/labstack/echo/v4"

type handler interface {
	Register(c echo.Context) error
	Login(c echo.Context) error
	GetOrders(c echo.Context) error
	SaveOrder(c echo.Context) error
	Balance(c echo.Context) error
	Withdraw(c echo.Context) error
	Withdrawals(c echo.Context) error
}

type Router struct {
	handler handler
}

func NewRouter(h handler) *Router {
	return &Router{handler: h}
}

func (r *Router) Route() *echo.Echo {
	e := echo.New()

	e.POST("/api/user/register", r.handler.Register)
	e.POST("/api/user/login", r.handler.Login)
	e.POST("/api/user/orders", r.handler.SaveOrder)
	e.POST("/api/user/balance/withdraw", r.handler.Withdraw)

	e.GET("/api/user/orders", r.handler.GetOrders)
	e.GET("/api/user/balance", r.handler.Balance)
	e.GET("/api/user/withdrawals", r.handler.Withdrawals)

	return e
}
