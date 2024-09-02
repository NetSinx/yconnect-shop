package service

import "github.com/NetSinx/yconnect-shop/server/order/model/entity"

type OrderService interface {
	ListOrder(order []entity.Order, username string) ([]entity.Order, error)
	AddOrder(order entity.Order) error
	DeleteOrder(order entity.Order, username, id string) error
}