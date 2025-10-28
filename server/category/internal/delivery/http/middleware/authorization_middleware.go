package middleware

import "github.com/labstack/echo/v4"

func AuthorizationMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		role := c.Request().Header.Get("X-User-Role")
		if role != "admin" {
			return echo.ErrUnauthorized
		}

		return next(c)
	}
}
