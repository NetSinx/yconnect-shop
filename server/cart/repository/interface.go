package repository

import "github.com/NetSinx/yconnect-shop/server/cart/model/entity"

type CartRepo interface {
	ListCart(cart []entity.Cart) ([]entity.Cart, error)
	AddToCart(cart entity.Cart) (entity.Cart, error)
	UpdateCart(cart entity.Cart, id uint) (entity.Cart, error)
	DeleteProductInCart(cart entity.Cart, id uint) error
	GetCart(cart entity.Cart, id uint) (entity.Cart, error)
	GetCartByUser(cart []entity.Cart, id uint) ([]entity.Cart, error)
}