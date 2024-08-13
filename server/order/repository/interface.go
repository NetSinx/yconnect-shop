package repository

import "github.com/NetSinx/yconnect-shop/server/order/model/entity"

type OrderRepository interface {
	ListOrder(order []entity.Order) []entity.Order
	AddOrder(order entity.Order) error
	GetOrder(order entity.Order, id uint) (entity.Order, error)
	DeleteOrder(order entity.Order, id uint) error
}