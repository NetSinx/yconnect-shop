package repository

import "github.com/NetSinx/yconnect-shop/server/seller/model/entity"

type SellerRepo interface {
	ListSeller() ([]entity.Seller, error)
}