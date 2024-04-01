package domain

import "github.com/NetSinx/yconnect-shop/server/product/app/model"

type ErrorResponse struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type FindAllResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type GetProductResponse struct {
	Code   int    					`json:"code"`
	Status string 					`json:"status"`
	Data   []model.Product	`json:"data"`
}