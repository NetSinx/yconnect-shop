package middleware

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"sync"
	"time"
	"github.com/labstack/echo/v4"
)

var (
	tokenStore = make(map[string]time.Time)
	mu         sync.Mutex
)

func generateToken() string {
	b := make([]byte, 32)
	_, _ = rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

func CSRFMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Method == http.MethodPost || 
			c.Request().Method == http.MethodPut || 
			c.Request().Method == http.MethodDelete {

			getToken, err := c.Request().Cookie("csrf_token")
			if err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"error": "CSRF token missing"})
			} else if getToken.Value == "" {
				return c.JSON(http.StatusBadRequest, map[string]string{"error": "CSRF token missing"})
			}

			mu.Lock()
			exp, ok := tokenStore[getToken.Value]
			if !ok || time.Now().After(exp) {
				mu.Unlock()
				return c.JSON(http.StatusForbidden, map[string]string{"error": "Invalid or expired CSRF token"})
			}

			delete(tokenStore, getToken.Value)
			mu.Unlock()
		}

		if c.Request().Method == http.MethodGet {
			token := generateToken()
			mu.Lock()
			tokenStore[token] = time.Now().Add(1 * time.Minute)
			mu.Unlock()

			csrfTokenCookie := http.Cookie{
				Name: "csrf_token",
				Path: "/",
				Value: token,
				HttpOnly: true,
				SameSite: http.SameSiteLaxMode,
				Secure: true,
			}

			c.SetCookie(&csrfTokenCookie)
		}

		return next(c)
	}
}
