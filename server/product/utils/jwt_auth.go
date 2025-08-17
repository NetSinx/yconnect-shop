package utils

import (
	"net/http"
	"github.com/NetSinx/yconnect-shop/server/authentication/utils"
	"github.com/NetSinx/yconnect-shop/server/product/model/domain"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func CheckAdminRole(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*utils.CustomClaims)

		if claims.Role != "admin" {
			return echo.NewHTTPError(http.StatusForbidden, domain.MessageResp{
				Message: "Forbidden Access.",
			})
		}

		return next(c)
	}
}