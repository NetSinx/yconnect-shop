package utils

import (
	"github.com/NetSinx/yconnect-shop/server/product/app/model"
	cartModel "github.com/NetSinx/yconnect-shop/server/cart/model"
)

type ErrServer struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type SuccessCUD struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type SuccessGet struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type PreloadProducts struct {
	Code   int							`json:"code"`
	Status string						`json:"status"`
	Data   []model.Product	`json:"data"`
}

type PreloadCarts struct {
	Code   int							`json:"code"`
	Status string						`json:"status"`
	Data   []cartModel.Cart	`json:"data"`
}