package controller

import (
	"net/http"
	"time"
	"github.com/NetSinx/yconnect-shop/server/order/model/domain"
	"github.com/NetSinx/yconnect-shop/server/order/model/entity"
	"github.com/NetSinx/yconnect-shop/server/order/service"
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

func (oc *orderController) ListOrder(c echo.Context) error {
	var order []entity.Order
	username := c.Param("username")

	orders, err := oc.orderService.ListOrder(order, username)
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
	var order entity.Order
	order.Estimasi = time.Now().Add(72 * time.Hour)

	if err := c.Bind(&order); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
			Message: err.Error(),
		})
	}

	err := oc.orderService.AddOrder(order)
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
		Message: "Pesanan berhasil ditambahkan. Siap untuk dikirim!",
	})
}

func (oc *orderController) DeleteOrder(c echo.Context) error {
	var order entity.Order

	username := c.Param("username")

	err := oc.orderService.DeleteOrder(order, username)
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