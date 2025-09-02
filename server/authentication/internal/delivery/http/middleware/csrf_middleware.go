package middleware

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"time"

	"github.com/NetSinx/yconnect-shop/server/authentication/internal/model"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
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

func (cc *CSRFConfig) CSRFMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()

		var token string
		resultToken, err := cc.RedisClient.Get(c.Request().Context(), "csrf_token").Result()
		if err != nil {
			token = generateCSRFToken()
			if err := cc.RedisClient.Set(c.Request().Context(), "csrf:"+token, "valid", 5*time.Minute).Err(); err != nil {
				return c.JSON(http.StatusInternalServerError, &model.MessageResponse{
					Message: "failed to set csrf token",
				})
			}
		} else {
			token = resultToken
		}

		if req.Method == http.MethodPost || req.Method == http.MethodPut || req.Method == http.MethodDelete {
			reqToken := req.Header.Get("X-CSRF-Token")
			if reqToken == "" || reqToken != token {
				return echo.NewHTTPError(http.StatusBadRequest, "invalid csrf token")
			}
		}

		c.Set("csrf_token", token)

		return next(c)
	}
}
