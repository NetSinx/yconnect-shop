package repository

import "github.com/NetSinx/yconnect-shop/server/order/model/entity"

type OrderRepository interface {
	ListOrder(order []entity.Order, username string) ([]entity.Order, error)
	AddOrder(order entity.Order) error
	DeleteOrder(order entity.Order, username string) error
}