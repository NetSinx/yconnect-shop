package middleware

import (
	"errors"
	"net/http"
	"time"
	"github.com/NetSinx/yconnect-shop/cart/utils"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type CustomClaims struct {
	Username  string  `json:"username"`
	Admin     bool    `json:"admin"`
	jwt.RegisteredClaims
}

func AuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			apiToken := c.Request().Header.Get("Authorization")

			if apiToken == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, utils.ErrServer{
					Code: http.StatusUnauthorized,
					Status: http.StatusText(http.StatusUnauthorized),
					Message: "You are not logged in!",
				})
			}

			claims := &jwt.RegisteredClaims{
				Issuer: "this is a jwt",
				IssuedAt: jwt.NewNumericDate(time.Now()),
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			}

			token, err := jwt.ParseWithClaims(apiToken, claims, func(t *jwt.Token) (interface{}, error) {
				return []byte("yasinganteng15"), nil
			})

			if errors.Is(err, jwt.ErrSignatureInvalid) {
				claims := &CustomClaims{
					"netsinx_15",
					true,
					jwt.RegisteredClaims{
						Issuer: "this is a jwt",
						IssuedAt: jwt.NewNumericDate(time.Now()),
						ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
					},
				}

				token, _ := jwt.ParseWithClaims(apiToken, claims, func(t *jwt.Token) (interface{}, error) {
					return []byte("netsinxadmin"), nil
				})

				if !token.Valid {
					return echo.NewHTTPError(http.StatusUnauthorized, utils.ErrServer{
						Code: http.StatusUnauthorized,
						Status: http.StatusText(http.StatusUnauthorized),
						Message: "You are not logged in!",
					})
				}

				return next(c)
			}

			if !token.Valid {
				return echo.NewHTTPError(http.StatusUnauthorized, utils.ErrServer{
					Code: http.StatusUnauthorized,
					Status: http.StatusText(http.StatusUnauthorized),
					Message: "You are not logged in!",
				})
			}

			return next(c)
		}
	}
}