package controller

import (
	"net/http"
	"github.com/NetSinx/yconnect-shop/server/order/model/domain"
	"github.com/NetSinx/yconnect-shop/server/order/model/entity"
	"github.com/NetSinx/yconnect-shop/server/order/service"
	"github.com/NetSinx/yconnect-shop/server/user/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type orderController struct {
	orderService service.OrderService
}

func OrderContrllr(orderService service.OrderService) *orderController {
	return &orderController{
		orderService: orderService,
	}
}

func (oc *orderController) GetOrder(c echo.Context) error {
	var order []entity.Order

	user_session, err := c.Cookie("user_session")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
			Message: "session token not available",
		})
	}

	token, err := jwt.ParseWithClaims(user_session.Value, &utils.CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte("yasinnetsinx15"), nil
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, domain.MessageResp{
			Message: err.Error(),
		})
	}
	if !token.Valid {
		return echo.NewHTTPError(http.StatusUnauthorized, domain.MessageResp{
			Message: "your token is invalid",
		})
	}

	username := token.Claims.(*utils.CustomClaims).Username
	orders, err := oc.orderService.GetOrder(order, username)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, domain.MessageResp{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, domain.DataResp{
		Data: orders,
	})
}

func (oc *orderController) AddOrder(c echo.Context) error {
	var reqOrder domain.OrderRequest

	if err := c.Bind(&reqOrder); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
			Message: err.Error(),
		})
	}

	err := oc.orderService.AddOrder(reqOrder)
	if err != nil && err == echo.ErrBadRequest{
		return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
			Message: err.Error(),
		})
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, domain.MessageResp{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, domain.MessageResp{
		Message: "Pesanan berhasil ditambahkan",
	})
}

func (oc *orderController) DeleteOrder(c echo.Context) error {
	var order entity.Order

	user_session, err := c.Cookie("user_session")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
			Message: "session token not available",
		})
	}

	token, err := jwt.ParseWithClaims(user_session.Value, &utils.CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte("yasinnetsinx15"), nil
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, domain.MessageResp{
			Message: err.Error(),
		})
	}
	if !token.Valid {
		return echo.NewHTTPError(http.StatusUnauthorized, domain.MessageResp{
			Message: "your token is invalid",
		})
	}

	username := token.Claims.(*utils.CustomClaims).Username
	err = oc.orderService.DeleteOrder(order, username)
	if err != nil && err.Error() == "pesanan tidak ditemukan" {
		return echo.NewHTTPError(http.StatusNotFound, domain.MessageResp{
			Message: "Pesanan tidak ditemukan",
		})
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, domain.MessageResp{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, domain.MessageResp{
		Message: "Pesanan berhasil dibatalkan",
	})
}