package helpers

import (
	"crypto/rand"
	"encoding/base64"
	"strings"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"github.com/NetSinx/yconnect-shop/server/authentication/helpers"
	"github.com/NetSinx/yconnect-shop/server/product/handler/dto"
)

func CheckAdminRole(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*helpers.CustomClaims)

		if claims.Role != "admin" {
			return echo.NewHTTPError(http.StatusForbidden, dto.MessageResp{
				Message: "Forbidden Access.",
			})
		}

		return next(c)
	}
}

func GenerateSlugByName(name string) (string, error) {
	trimmed := strings.TrimSpace(name)
	words := strings.Fields(trimmed)
	name = strings.Join(words, " ")
	name = strings.ToLower(name)
	name = strings.ReplaceAll(name, " ", "-")

	b := make([]byte, 8)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	uniqueId := base64.URLEncoding.EncodeToString(b)

	return name + "-" + uniqueId, nil
}

func ReplaceProductSlug(slug string, namaProduct string) string {
	splitSlug := strings.Split(slug, "-")
	uid := splitSlug[len(splitSlug) - 1]

	trimmed := strings.TrimSpace(namaProduct)
	words := strings.Fields(trimmed)
	namaProduct = strings.Join(words, " ")
	namaProduct = strings.ToLower(namaProduct)
	namaProduct = strings.ReplaceAll(namaProduct, " ", "-")
	newSlug := namaProduct + "-" + uid

	return newSlug
}