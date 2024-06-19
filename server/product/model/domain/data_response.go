package domain

import (
	"github.com/NetSinx/yconnect-shop/server/product/model/entity"
)

type RespData struct {
	Data   interface{} `json:"data"`
}

type PreloadCategory struct {
	Data   entity.Category `json:"data"`
}

type PreloadUser struct {
	Data   entity.Seller	`json:"data"`
}

type MessageResp struct {
	Message string `json:"message"`
}

type ResponseCSRF struct {
	CSRFToken string `json:"csrf_token"`
}