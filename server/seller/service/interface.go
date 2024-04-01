package service

import "github.com/NetSinx/yconnect-shop/server/seller/model/entity"

type SellerServ interface {
	ListSeller() ([]entity.Seller, error)
}