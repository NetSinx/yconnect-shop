package utils

import (
	"net/http"
	"time"
	"github.com/labstack/echo/v4"
)

func SetCookies(name string, value string) {
	var c echo.Context
	cookie := http.Cookie{
		Name: name,
		Value: value,
		Expires: time.Now().Add(30 * time.Minute),
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Secure: true,
	}
	
	c.SetCookie(&cookie)
}