package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"net/http"
)

type CSRFConfig struct {
	RedisClient *redis.Client
}

func CSRFMiddleware(csrfConfig *CSRFConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()

			if req.Method == http.MethodPost || req.Method == http.MethodPut || req.Method == http.MethodDelete {
				reqToken := req.Header.Get("CSRF-Token")
				if reqToken == "" {
					return echo.NewHTTPError(http.StatusBadRequest, "missing csrf token")
				}

				result, err := csrfConfig.RedisClient.GetDel(c.Request().Context(), "csrf:"+reqToken).Result()
				if err != nil {
					return echo.NewHTTPError(http.StatusInternalServerError, "failed to getting csrf token")
				}
				
				if result != "valid" {
					return echo.NewHTTPError(http.StatusForbidden, "invalid csrf token")
				}
			}

			return next(c)
		}
	}
}
