package test

import (
	"net/http"
	"strconv"
	"time"
	"github.com/NetSinx/yconnect-shop/server/order/model/domain"
	"github.com/NetSinx/yconnect-shop/server/order/model/entity"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type DataTest struct {
	Data entity.Order
}

var modelDB = map[string]*DataTest{
	"netsinx_15": {
		Data: entity.Order{
			Id: 1,
			ProductID: 1,
			Kuantitas: 5,
			Status: "Dalam pengiriman",
			Estimasi: time.Now().AddDate(0, 0, 3),
		},
	},
}

func AddOrder(c echo.Context) error {
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

	return c.JSON(http.StatusOK, domain.MessageResp{
		Message: "Pesanan sedang diproses oleh penjual",
	})
}

func ListOrder(c echo.Context) error {
	username := c.Param("username")

	orders, ok := modelDB[username]
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, domain.MessageResp{
			Message: "Pesanan masih kosong.",
		})
	}

	return c.JSON(http.StatusOK, orders)
}

func DeleteOrder(c echo.Context) error {
	username := c.Param("username")
	paramId := c.Param("id")
	id, _ := strconv.Atoi(paramId)

	getOrder, ok := modelDB[username]
	if !ok {
		return echo.NewHTTPError(http.StatusNotFound, domain.MessageResp{
			Message: "Pesanan tidak ditemukan",
		})
	}
	
	if getOrder.Data.Id == uint(id) {
		getOrder = nil
		
		return c.JSON(http.StatusOK, domain.MessageResp{
			Message: "Pesanan berhasil dibatalkan",
		})
	}

	return echo.NewHTTPError(http.StatusNotFound, domain.MessageResp{
		Message: "Pesanan tidak ditemukan",
	})
}