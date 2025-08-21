package middleware

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"sync"
	"time"

	"github.com/NetSinx/yconnect-shop/server/authentication/handler/dto"
	"github.com/labstack/echo/v4"
)

type CSRFManager struct {
	tokens map[string]time.Time
	mu     sync.Mutex
	ttl    time.Duration
}

func NewCSRFManager(ttl time.Duration) *CSRFManager {
	return &CSRFManager{
		tokens: make(map[string]time.Time),
		ttl:    ttl,
	}
}

func (m *CSRFManager) GenerateToken() string {
	b := make([]byte, 32)
	_, _ = rand.Read(b)
	token := hex.EncodeToString(b)

	m.mu.Lock()
	defer m.mu.Unlock()
	m.tokens[token] = time.Now().Add(m.ttl)
	return token
}

func (m *CSRFManager) ValidateToken(token string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	exp, exists := m.tokens[token]
	if !exists {
		return false
	}

	if time.Now().After(exp) {
		delete(m.tokens, token)
		return false
	}

	delete(m.tokens, token)
	return true
}

func CSRFMiddleware(m *CSRFManager) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if c.Request().Method == http.MethodPost ||
				c.Request().Method == http.MethodPut ||
				c.Request().Method == http.MethodDelete {

				token, err := c.Request().Cookie("csrf_token")
				if !m.ValidateToken(token.Value) || err != nil {
					return c.JSON(http.StatusBadRequest, map[string]string{
						"error": "Invalid or expired CSRF token",
					})
				}
			}

			return next(c)
		}
	}
}

func GetCSRFTokenHandler(m *CSRFManager) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := m.GenerateToken()

		csrfTokenCookie := http.Cookie{
			Name: "csrf_token",
			Path: "/",
			Value: token,
			SameSite: http.SameSiteLaxMode,
			HttpOnly: true,
			Secure: true,
		}

		c.SetCookie(&csrfTokenCookie)

		return c.JSON(http.StatusOK, dto.MessageResp{
			Message: "CSRF token berhasil di-generate",
		})
	}
}