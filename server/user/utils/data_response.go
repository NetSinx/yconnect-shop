package utils

import (
	"github.com/NetSinx/yconnect-shop/server/seller/model/entity"
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

type PreloadSeller struct {
	Code   int							`json:"code"`
	Status string						`json:"status"`
	Data   entity.Seller	`json:"data"`
}

type PreloadCarts struct {
	Code   int							`json:"code"`
	Status string						`json:"status"`
	Data   []cartModel.Cart	`json:"data"`
}