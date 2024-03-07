package service

import "github.com/NetSinx/yconnect-shop/server/cart/model"

type CartServ interface {
	ListCart(carts []model.Cart) ([]model.Cart, error)
	AddToCart(cart model.Cart, id int) (model.Cart, error)
	UpdateCart(cart model.Cart, id uint) (model.Cart, error)
	DeleteProductInCart(cart model.Cart, id uint) error
	GetCart(cart model.Cart, id uint) (model.Cart, error)
	GetCartByUser(cart []model.Cart, id uint) ([]model.Cart, error)
}