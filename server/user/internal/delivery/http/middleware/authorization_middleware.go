package middleware

import "github.com/labstack/echo/v4"

func AuthorizationMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Request().Header.Get("X-User-ID")
		if id != c.Param("id") {
			return echo.ErrUnauthorized
		}

		return next(c)
	}
}