package middleware

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"net/http"
	"time"
)

type CSRFConfig struct {
	RedisClient *redis.Client
}

func generateCSRFToken() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	return base64.RawURLEncoding.EncodeToString(b)
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

				if err := csrfConfig.RedisClient.Exists(c.Request().Context(), "csrf:"+reqToken).Err(); err != nil {
					return echo.NewHTTPError(http.StatusBadRequest, "invalid csrf token")
				}

				csrfConfig.RedisClient.Del(c.Request().Context(), "csrf:"+reqToken)
			}

			token := generateCSRFToken()
			csrfConfig.RedisClient.Set(c.Request().Context(), "csrf:"+token, "valid", 5*time.Minute)

			c.Set("csrf_token", token)

			return next(c)
		}
	}
}
