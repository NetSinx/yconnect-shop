package controller

import "github.com/labstack/echo/v4"

type UserContr interface {
	RegisterUser(c echo.Context) error
	LoginUser(c echo.Context) error
	ListUsers(c echo.Context) error
	UpdateUser(c echo.Context) error
	GetUser(c echo.Context) error
	DeleteUser(c echo.Context) error
	SendOTP(c echo.Context) error
	VerifyEmail(c echo.Context) error
	Verify(c echo.Context) error
	SetTimezone(c echo.Context) error
	UserLogout(c echo.Context) error
}