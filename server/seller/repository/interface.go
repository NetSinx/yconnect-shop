package repository

import "github.com/NetSinx/yconnect-shop/server/seller/model/entity"

type SellerRepo interface {
	ListSeller() ([]entity.Seller, error)
	RegisterSeller(seller entity.Seller) (entity.Seller, error)
	UpdateSeller(username string, seller entity.Seller) (entity.Seller, error)
	DeleteSeller(username string) error
	GetSeller(username string) (entity.Seller, error)
}