package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"github.com/NetSinx/yconnect-shop/server/cart/model"
	"github.com/NetSinx/yconnect-shop/server/cart/repository"
	"github.com/NetSinx/yconnect-shop/server/cart/utils"
	"github.com/go-playground/validator/v10"
)

type cartService struct {
	cartRepo repository.CartRepo
}

func CartService(cr repository.CartRepo) cartService {
	return cartService{
		cartRepo: cr,
	}
}

func (cr cartService) ListCart(carts []model.Cart) ([]model.Cart, error) {
	listCart, err := cr.cartRepo.ListCart(carts)
	if err != nil {
		return nil, err
	}

	for i := range listCart {
		var preloadProduct utils.PreloadProduct

		resp, err := http.Get(fmt.Sprintf("http://product-service:8081/product/%d", listCart[i].ProductID))
		if err != nil {
			return carts, err
		}

		json.NewDecoder(resp.Body).Decode(&preloadProduct)

		listCart[i].Product = preloadProduct.Data
	}

	return listCart, nil
}

func (cs cartService) AddToCart(cart model.Cart, id int) (model.Cart, error) {
	var preloadProduct utils.PreloadProduct

	cart.ProductID = id

	resp, err := http.Get(fmt.Sprintf("http://product-service:8081/product/%d", cart.ProductID))
	if err != nil {
		return cart, err
	}

	json.NewDecoder(resp.Body).Decode(&preloadProduct)

	cart.Product = preloadProduct.Data

	if err := validator.New().Struct(cart); err != nil {
		return cart, errors.New("request tidak sesuai")
	}

	addCart, err := cs.cartRepo.AddToCart(cart)
	if err != nil {
		return cart, errors.New("produk sudah ada")
	}

	return addCart, nil
}

func (cs cartService) UpdateCart(cart model.Cart, id uint) (model.Cart, error) {
	var preloadProduct utils.PreloadProduct
	
	getCart, _ := cs.cartRepo.GetCart(cart, id)

	resp, err := http.Get(fmt.Sprintf("http://product-service:8081/product/%d", getCart.ProductID))
	if err != nil {
		return cart, err
	}

	json.NewDecoder(resp.Body).Decode(&preloadProduct)

	getCart.Product = preloadProduct.Data
	getCart.Item += cart.Item

	if err := validator.New().Struct(getCart); err != nil {
		return cart, errors.New("request tidak sesuai")
	}
	
	updCart, err := cs.cartRepo.UpdateCart(getCart, id)
	if err != nil {
		return cart, err
	}

	return updCart, nil
}

func (cs cartService) DeleteProductInCart(cart model.Cart, id uint) error {
	if err := cs.cartRepo.DeleteProductInCart(cart, id); err != nil {
		return err
	}

	return nil
}

func (cs cartService) GetCart(cart model.Cart, id uint) (model.Cart, error) {
	getCart, err := cs.cartRepo.GetCart(cart, id)
	if err != nil {
		return cart, err
	}

	var preloadProduct utils.PreloadProduct

	resp, err := http.Get(fmt.Sprintf("http://product-service:8081/product/%d", getCart.ProductID))
	if err != nil {
		return cart, err
	}

	json.NewDecoder(resp.Body).Decode(&preloadProduct)

	getCart.Product = preloadProduct.Data

	return getCart, nil
}

func (cs cartService) GetCartByUser(cart []model.Cart, id uint) ([]model.Cart, error) {
	getCart, err := cs.cartRepo.GetCartByUser(cart, id)
	if err != nil {
		return nil, err
	}

	for i := range getCart {
		var preloadProduct utils.PreloadProduct
	
		resp, err := http.Get(fmt.Sprintf("http://product-service:8081/product/%d", getCart[i].ProductID))
		if err != nil {
			return cart, err
		}
	
		json.NewDecoder(resp.Body).Decode(&preloadProduct)

		getCart[i].Product = preloadProduct.Data
	}

	return getCart, nil
}