package repository

import "github.com/NetSinx/yconnect-shop/server/order/model/entity"

type OrderRepository interface {
	GetOrder(order []entity.Order, user_id string) ([]entity.Order, error)
	AddOrder(order entity.Order) error
	DeleteOrder(order entity.Order, username, id string) error
}