package domain

import (
	modelProd "github.com/NetSinx/yconnect-shop/server/product/app/model"
	modelUser "github.com/NetSinx/yconnect-shop/server/user/app/model"
)

type Response struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type GetProductResponse struct {
	Code   int    					`json:"code"`
	Status string 					`json:"status"`
	Data   []modelProd.Product	`json:"data"`
}

type GetUserResponse struct {
	Code   int    					`json:"code"`
	Status string 					`json:"status"`
	Data   modelUser.User	`json:"data"`
}