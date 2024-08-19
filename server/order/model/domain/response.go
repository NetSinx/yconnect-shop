package domain

import "github.com/NetSinx/yconnect-shop/server/product/model/entity"

type MessageResp struct {
	Message string `json:"message"`
}

type DataResp struct {
	Data interface{} `json:"data"`
}

type DataProduct struct {
	Data entity.Product `json:"data"`
}