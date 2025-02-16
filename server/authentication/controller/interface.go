package controller

import "github.com/labstack/echo/v4"

type AuthController interface {
	LoginUser(ctx echo.Context) error
	LogoutUser(ctx echo.Context) error
	RefreshToken(ctx echo.Context) error
}