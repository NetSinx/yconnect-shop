package domain

import (
	cartModel "github.com/NetSinx/yconnect-shop/server/cart/model/entity"
	orderModel "github.com/NetSinx/yconnect-shop/server/order/model/entity"
)

type MessageResp struct {
	Message string `json:"message"`
}

type RespData struct {
	Data   interface{} `json:"data"`
}

type PreloadCarts struct {
	Data   []cartModel.Cart	`json:"data"`
}

type PreloadOrders struct {
	Data   []orderModel.Order	`json:"data"`
}