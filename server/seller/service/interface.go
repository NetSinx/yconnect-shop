package service

import (
	"github.com/NetSinx/yconnect-shop/server/seller/model/entity"
	"github.com/NetSinx/yconnect-shop/server/seller/model/domain"
)

type SellerServ interface {
	ListSeller() ([]entity.Seller, error)
	RegisterSeller(username string, sellerValidity domain.Seller) (entity.Seller, error)
	UpdateSeller(username string) (entity.Seller, error)
	DeleteSeller(username string) error
	GetSeller(username string) (entity.Seller, error)
}