package middleware

import (
	"net/http"
	"time"
	"github.com/NetSinx/yconnect-shop/product/utils"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		apiToken := c.Request().Header.Get("Authorization")

		if apiToken == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, utils.ErrServer{
				Code:    http.StatusUnauthorized,
				Status:  "Unauthorized",
				Message: "You are not logged in!",
			})
		}

		claims := &jwt.RegisteredClaims{
				Issuer:    "this is jwt",
				IssuedAt: jwt.NewNumericDate(time.Now()),
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		}

		token, _ := jwt.ParseWithClaims(apiToken, claims, func(t *jwt.Token) (interface{}, error) {
				return []byte("yasinganteng15"), nil
		})
	
		if !token.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, utils.ErrServer{
				Code:    http.StatusUnauthorized,
				Status:  "Unauthorized",
				Message: "You are not logged in!",
			})
		}

		return next(c)
	}
}