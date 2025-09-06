package middleware

import (
	"net/http"
	"strings"
	"github.com/NetSinx/yconnect-shop/server/user/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func JWTAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		auth_token, err := c.Cookie("user_session")
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, map[string]string{
				"message": "missing jwt token authentication",
			})
		}

		token, err := jwt.ParseWithClaims(auth_token.Value, &utils.CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte("yasinnetsinx15"), nil
		})
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, map[string]string{
				"message": err.Error(),
			})
		} else if !token.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, map[string]string{
				"message": "your token is invalid",
			})
		}
		
		claims := token.Claims.(*utils.CustomClaims)
		if claims.Role != "admin" && strings.Split(c.Path(), "/")[1] == "admin" {
			return echo.NewHTTPError(http.StatusUnauthorized, map[string]string{
				"message": "your claims is invalid",
			})
		} else if claims.Username != "" && claims.Role != "" {
			return echo.NewHTTPError(http.StatusUnauthorized, map[string]string{
				"message": "your claims is invalid",
			})
		}

		return next(c)
	}
}