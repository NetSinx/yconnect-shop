package service

import "github.com/NetSinx/yconnect-shop/server/seller/model/entity"

type SellerServ interface {
	ListSeller() ([]entity.Seller, error)
	RegisterSeller(username string) (entity.Seller, error)
	UpdateSeller(username string) (entity.Seller, error)
	DeleteSeller(username string) error
	GetSeller(username string) (entity.Seller, error)
}