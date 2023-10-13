package router

import (
	"github.com/labstack/echo/v4"
)

type handler interface {
	GetOrders(c echo.Context) error
	SaveOrder(c echo.Context) error
	Balance(c echo.Context) error
	Withdraw(c echo.Context) error
	Withdrawals(c echo.Context) error
}

type loginHandler interface {
	Register(c echo.Context) error
	Login(c echo.Context) error
}

type authMiddleware interface {
	CheckAuth(next echo.HandlerFunc) echo.HandlerFunc
}

type Router struct {
	handler        handler
	authMiddleware authMiddleware
	loginHandler   loginHandler
}

func NewRouter(h handler, lh loginHandler, a authMiddleware) *Router {
	return &Router{
		handler:        h,
		authMiddleware: a,
		loginHandler:   lh,
	}
}

func (r *Router) Route() *echo.Echo {
	e := echo.New()

	e.POST("/api/user/register", r.loginHandler.Register)
	e.POST("/api/user/login", r.loginHandler.Login)

	e.POST("/api/user/orders", r.handler.SaveOrder, r.authMiddleware.CheckAuth)
	e.POST("/api/user/balance/withdraw", r.handler.Withdraw, r.authMiddleware.CheckAuth)

	e.GET("/api/user/orders", r.handler.GetOrders, r.authMiddleware.CheckAuth)
	e.GET("/api/user/balance", r.handler.Balance, r.authMiddleware.CheckAuth)
	e.GET("/api/user/withdrawals", r.handler.Withdrawals, r.authMiddleware.CheckAuth)

	return e
}
