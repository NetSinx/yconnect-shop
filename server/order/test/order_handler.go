package test

import (
	"net/http"
	"github.com/NetSinx/yconnect-shop/server/order/model/domain"
	"github.com/NetSinx/yconnect-shop/server/order/model/entity"
	"github.com/labstack/echo/v4"
)

type orderHandler struct {
	db map[string]*entity.Order
}

func (h *orderHandler) AddOrder(c echo.Context) error {
	var orderModel entity.Order

	if err := c.Bind(&orderModel); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, domain.MessageResp{
		Message: "Pesanan berhasil dibuat",
	})
}

func (h *orderHandler) ListOrder(c echo.Context) error {
	username := c.Param("username")
	listOrder := h.db[username]

	if listOrder == nil {
		return echo.NewHTTPError(http.StatusNotFound, domain.MessageResp{
			Message: "Pesanan masih kosong",
		})
	}

	return c.JSON(http.StatusOK, domain.DataResp{
		Data: listOrder,
	})
}