package middleware

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"github.com/labstack/echo/v4"
)

type CSRFMiddleware struct {
	Base64Encoding *base64.Encoding
	Token        string
}

func (cm *CSRFMiddleware) generateCSRFToken() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	return cm.Base64Encoding.EncodeToString(b)
}

func (cm *CSRFMiddleware) NewCSRFMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		res := c.Response()

		cookie, err := c.Cookie("csrf_token")
		if err != nil {
			cm.Token = cm.generateCSRFToken()
		} else {
			cm.Token = cookie.Value
		}

		if req.Method == http.MethodPost || req.Method == http.MethodPut || req.Method == http.MethodDelete {
			clientToken := req.Header.Get("X-CSRF-Token")
			if clientToken == "" || clientToken != cm.Token {
				return echo.NewHTTPError(http.StatusForbidden, "invalid csrf token")
			}
		}

		newToken := cm.generateCSRFToken()
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