package utils

import (
	"github.com/NetSinx/yconnect-shop/product/app/model"
)

type SuccessGetData struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type PreloadCategory struct {
	Code   int              `json:"code"`
	Status string           `json:"status"`
	Data   []model.Category `json:"data"`
}

type PreloadUser struct {
	Code   int          `json:"code"`
	Status string       `json:"status"`
	Data   []model.User	`json:"data"`
}

type SuccessCUD struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ErrServer struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}