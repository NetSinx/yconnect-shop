package utils

import (
	"github.com/NetSinx/yconnect-shop/cart/model"
	prodModel "github.com/NetSinx/yconnect-shop/product/app/model"
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

type PreloadCategory struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   model.Category `json:"data"`
}

type PreloadUser struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   model.User `json:"data"`
}

type PreloadProduct struct {
	Code   int               `json:"code"`
	Status string            `json:"status"`
	Data   prodModel.Product `json:"data"`
}