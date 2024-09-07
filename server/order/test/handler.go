package test

import (
	"net/http"
	"strconv"
	"github.com/NetSinx/yconnect-shop/server/order/model/domain"
	"github.com/NetSinx/yconnect-shop/server/order/model/entity"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type DataTest struct {
	Data entity.Order
}

type OrderHandler struct {
	db map[string]*DataTest
}

func (h *OrderHandler) AddOrder(c echo.Context) error {
	var orderModel entity.Order

	if err := c.Bind(&orderModel); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
			Message: err.Error(),
		})
	}

	if err := validator.New().Struct(orderModel); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, orderModel)
}

func (h *OrderHandler) ListOrder(c echo.Context) error {
	username := c.Param("username")
	listOrder := h.db[username]

	if listOrder == nil {
		return echo.NewHTTPError(http.StatusNotFound, domain.MessageResp{
			Message: "Pesanan masih kosong",
		})
	}

	return c.JSON(http.StatusOK, listOrder)
}

func (h *OrderHandler) DeleteOrder(c echo.Context) error {
	username := c.Param("username")
	paramId := c.Param("id")
	id, _ := strconv.Atoi(paramId)

	getOrder := h.db[username]
	if getOrder.Data.Id == uint(id) {
		getOrder = nil

		if getOrder != nil {
			return echo.NewHTTPError(http.StatusNotFound, domain.MessageResp{
				Message: "Pesanan tidak ditemukan",
			})
		}
	}
	
	return c.JSON(http.StatusOK, domain.MessageResp{
		Message: "Pesanan berhasil dibatalkan",
	})
}