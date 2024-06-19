package domain

import (
	"github.com/NetSinx/yconnect-shop/server/seller/model/entity"
	cartModel "github.com/NetSinx/yconnect-shop/server/cart/model"
)

type MessageResp struct {
	Message string `json:"message"`
}

type RespData struct {
	Data   interface{} `json:"data"`
}

type PreloadSeller struct {
	Data   entity.Seller	  `json:"data"`
}

type PreloadCarts struct {
	Data   []cartModel.Cart	`json:"data"`
}