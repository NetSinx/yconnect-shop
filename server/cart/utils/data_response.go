package utils

import (
	"github.com/NetSinx/yconnect-shop/server/product/app/model"
)

type ErrServer struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type SuccessGet struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type SuccessDelete struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type PreloadProduct struct {
	Code   int           `json:"code"`
	Status string        `json:"status"`
	Data   model.Product `json:"data"`
}