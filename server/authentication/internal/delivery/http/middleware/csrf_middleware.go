package middleware

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"github.com/labstack/echo/v4"
)

func generateCSRFToken() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	return base64.RawURLEncoding.EncodeToString(b)
}

func CSRFMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		res := c.Response()

		var token string
		cookie, err := c.Cookie("csrf_token")
		if err != nil {
			token = generateCSRFToken()
		} else {
			token = cookie.Value
		}

		if req.Method == http.MethodPost || req.Method == http.MethodPut || req.Method == http.MethodDelete {
			clientToken := req.Header.Get("X-CSRF-Token")
			if clientToken == "" || clientToken != token {
				return echo.NewHTTPError(http.StatusForbidden, "invalid csrf token")
			}
		}

		newToken := generateCSRFToken()
		http.SetCookie(res, &http.Cookie{
			Name:     "csrf_token",
			Value:    newToken,
			Path:     "/",
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteStrictMode,
		})

		c.Set("csrf_token", newToken)

		return next(c)
	}
}