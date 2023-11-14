package login

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

type logger interface {
	Info(msg string)
	Warn(msg string)
	Error(msg string)
}

type LoggerMiddleware struct {
	logger logger
}

func NewLoggerMiddleware(l logger) *LoggerMiddleware {
	return &LoggerMiddleware{
		logger: l,
	}
}

func (l *LoggerMiddleware) Logging(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		userID := c.Get("userID")

		var id int
		if userID != nil {
			id = userID.(int)
		}

		l.logger.Info(fmt.Sprintf("METHOD: %s, URL: %s, userID: %d",
			c.Request().Method,
			c.Request().URL,
			id,
		))

		return next(c)
	}
}
