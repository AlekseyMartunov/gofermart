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

type loggerMiddleware interface {
	Logging(next echo.HandlerFunc) echo.HandlerFunc
}

type Router struct {
	userHandler    userHandler
	orderHandler   orderHandler
	loginHandler   loginHandler
	authMiddleware authMiddleware
	log            loggerMiddleware
}

func NewRouter(uh userHandler, oh orderHandler, lh loginHandler, am authMiddleware, lm loggerMiddleware) *Router {
	return &Router{
		userHandler:    uh,
		orderHandler:   oh,
		loginHandler:   lh,
		authMiddleware: am,
		log:            lm,
	}
}

func (r *Router) Route() *echo.Echo {
	e := echo.New()

	e.POST("/api/user/register", r.loginHandler.Register, r.log.Logging)
	e.POST("/api/user/login", r.loginHandler.Login, r.log.Logging)

	e.POST("/api/user/orders", r.orderHandler.SaveOrder, r.authMiddleware.CheckAuth, r.log.Logging)
	e.GET("/api/user/orders", r.orderHandler.GetOrders, r.authMiddleware.CheckAuth, r.log.Logging)

	e.POST("/api/user/balance/withdraw", r.userHandler.Withdraw, r.authMiddleware.CheckAuth, r.log.Logging)
	e.GET("/api/user/balance", r.userHandler.Balance, r.authMiddleware.CheckAuth, r.log.Logging)
	e.GET("/api/user/withdrawals", r.userHandler.Withdrawals, r.authMiddleware.CheckAuth, r.log.Logging)

	return e
}
