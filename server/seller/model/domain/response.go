package domain

import (
	modelProd "github.com/NetSinx/yconnect-shop/server/product/app/model"
	modelUser "github.com/NetSinx/yconnect-shop/server/user/app/model"
)

type MessageResp struct {
	Message string `json:"message"`
}

type RespData struct {
	Data   interface{} `json:"data"`
}

type GetProductResponse struct {
	Data   []modelProd.Product	`json:"data"`
}

type GetUserResponse struct {
	Data   modelUser.User	`json:"data"`
}