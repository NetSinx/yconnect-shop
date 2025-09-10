package helpers

import (
	"net/http"
	"time"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func FatalError(log *logrus.Logger, err error, msg string) {
  if err != nil {
    log.Fatalf("%s: %s", msg, err)
  }
}

func PanicError(log *logrus.Logger, err error, msg string) {
  if err != nil {
    log.Panicf("%s: %s", msg, err)
  }
}

func SetCookie(ctx echo.Context, name string, value string, expire time.Time) {
  ctx.SetCookie(&http.Cookie{
		Name: name,
		Path: "/",
		Value: value,
		Secure: true,
		HttpOnly: true,
		Expires: expire,
		SameSite: http.SameSiteStrictMode,
	})
}