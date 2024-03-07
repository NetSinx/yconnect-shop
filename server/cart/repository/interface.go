package repository

import "github.com/NetSinx/yconnect-shop/server/cart/model"

type CartRepo interface {
	ListCart(cart []model.Cart) ([]model.Cart, error)
	AddToCart(cart model.Cart) (model.Cart, error)
	UpdateCart(cart model.Cart, id uint) (model.Cart, error)
	DeleteProductInCart(cart model.Cart, id uint) error
	GetCart(cart model.Cart, id uint) (model.Cart, error)
	GetCartByUser(cart []model.Cart, id uint) ([]model.Cart, error)
}