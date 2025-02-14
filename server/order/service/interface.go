package service

import (
	"github.com/NetSinx/yconnect-shop/server/order/model/domain"
	"github.com/NetSinx/yconnect-shop/server/order/model/entity"
)

type OrderService interface {
	GetOrder(order []entity.Order, username string) ([]entity.Order, error)
	AddOrder(reqOrder domain.OrderRequest) error
	DeleteOrder(order entity.Order, username string) error
}