package utils

import (
	"net/http"
	"os"
	"github.com/NetSinx/yconnect-shop/server/product/model/domain"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

var (
	AdminJwtKey = os.Getenv("JWT_KEY_ADMIN")
	CustomerJwtKey = os.Getenv("JWT_KEY_CUSTOMER")
)


type CustomClaims struct {
	Username  string     `json:"username"`
	Role      string     `json:"role"`
	jwt.RegisteredClaims
}

func CheckAdminRole(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*CustomClaims)

		if claims.Role != "admin" {
			return echo.NewHTTPError(http.StatusForbidden, domain.MessageResp{
				Message: "Forbidden Access.",
			})
		}

		return next(c)
	}
}