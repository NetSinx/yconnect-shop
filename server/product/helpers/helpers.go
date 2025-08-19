package helpers

import (
	"crypto/rand"
	"encoding/base64"
	"strings"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"github.com/NetSinx/yconnect-shop/server/authentication/utils"
	"github.com/NetSinx/yconnect-shop/server/product/handler/dto"
)

func CheckAdminRole(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*utils.CustomClaims)

		if claims.Role != "admin" {
			return echo.NewHTTPError(http.StatusForbidden, dto.MessageResp{
				Message: "Forbidden Access.",
			})
		}

		return next(c)
	}
}

func GenerateSlugByName(name string) string {
	name = strings.ToLower(name)
	name = strings.ReplaceAll(name, " ", "-")

	b := make([]byte, 8)
	rand.Read(b)
	uniqueId := base64.URLEncoding.EncodeToString(b)

	return name + "-" + uniqueId
}