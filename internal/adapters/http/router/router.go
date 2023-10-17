package router

import (
	"github.com/labstack/echo/v4"
)

type userHandler interface {
	Balance(c echo.Context) error
	Withdraw(c echo.Context) error
	Withdrawals(c echo.Context) error
}

type orderHandler interface {
	GetOrders(c echo.Context) error
	SaveOrder(c echo.Context) error
}

type loginHandler interface {
	Register(c echo.Context) error
	Login(c echo.Context) error
}

type authMiddleware interface {
	CheckAuth(next echo.HandlerFunc) echo.HandlerFunc
}

type Router struct {
	userHandler    userHandler
	orderHandler   orderHandler
	loginHandler   loginHandler
	authMiddleware authMiddleware
}

func NewRouter(uh userHandler, oh orderHandler, lh loginHandler, m authMiddleware) *Router {
	return &Router{
		userHandler:    uh,
		orderHandler:   oh,
		loginHandler:   lh,
		authMiddleware: m,
	}
}

func (r *Router) Route() *echo.Echo {
	e := echo.New()

	e.POST("/api/user/register", r.loginHandler.Register)
	e.POST("/api/user/login", r.loginHandler.Login)

	e.POST("/api/user/orders", r.orderHandler.SaveOrder, r.authMiddleware.CheckAuth)
	e.GET("/api/user/orders", r.orderHandler.GetOrders, r.authMiddleware.CheckAuth)

	e.POST("/api/user/balance/withdraw", r.userHandler.Withdraw, r.authMiddleware.CheckAuth)
	e.GET("/api/user/balance", r.userHandler.Balance, r.authMiddleware.CheckAuth)
	e.GET("/api/user/withdrawals", r.userHandler.Withdrawals, r.authMiddleware.CheckAuth)

	return e
}
