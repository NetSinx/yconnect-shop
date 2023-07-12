package repository

import "github.com/NetSinx/yconnect-shop/cart/model"

type CartRepo interface {
	ListCart(cart []model.Cart) ([]model.Cart, error)
	CreateCart(cart model.Cart) (model.Cart, error)
	UpdateCart(cart model.Cart, id string) (model.Cart, error)
	DeleteCart(cart model.Cart, id string) error
	GetCartById(cart model.Cart, id string) (model.Cart, error)
	GetCartBySlug(cart model.Cart, slug string) (model.Cart, error)
}