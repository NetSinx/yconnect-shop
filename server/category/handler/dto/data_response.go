package dto

import "github.com/NetSinx/yconnect-shop/server/product/model"

var (
	CreateResponse = "Kategori berhasil ditambahkan."
	UpdateResponse = "Kategori berhasil diubah."
	DeleteResponse = "Kategori berhasil dihapus."
)

type MessageResp struct {
	Message string `json:"message"`
}

type RespData struct {
	Data any `json:"data"`
}

type RespProduct struct {
	Data []model.Product `json:"data"`
}