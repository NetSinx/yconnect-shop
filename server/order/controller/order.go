package controller

import (
	"net/http"
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

	orders := oc.orderService.ListOrder(order)

	return c.JSON(http.StatusOK, domain.DataResp{
		Data: orders,
	})
}