package utils

import "github.com/NetSinx/yconnect-shop/product/app/model"

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