package domain

import "github.com/NetSinx/yconnect-shop/server/product/model/entity"

type MessageResp struct {
	Message string `json:"message"`
}

type RespData struct {
	Data interface{} `json:"data"`
}

type PreloadProducts struct {
	Data entity.Product `json:"data"`
}