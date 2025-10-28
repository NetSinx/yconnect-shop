package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/NetSinx/yconnect-shop/server/cart/model/entity"
	"github.com/NetSinx/yconnect-shop/server/cart/repository"
	"github.com/NetSinx/yconnect-shop/server/cart/model/domain"
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

func (cr cartService) ListCart(carts []entity.Cart) ([]entity.Cart, error) {
	listCart, err := cr.cartRepo.ListCart(carts)
	if err != nil {
		return nil, err
	}

	for i := range listCart {
		var preloadProduct domain.PreloadProduct

		resp, err := http.Get(fmt.Sprintf("http://product-service:8081/product/%d", listCart[i].ProductID))
		if err != nil {
			return carts, err
		}

		json.NewDecoder(resp.Body).Decode(&preloadProduct)

		listCart[i].Product = preloadProduct.Data
	}

	return listCart, nil
}

func (cs cartService) AddToCart(cart entity.Cart, id int) (entity.Cart, error) {
	var preloadProduct domain.PreloadProduct

	cart.ProductID = id

	resp, err := http.Get(fmt.Sprintf("http://product-service:8081/product/%d", cart.ProductID))
	if err != nil {
		return cart, fmt.Errorf("Error get product: %v", err)
	}

	json.NewDecoder(resp.Body).Decode(&preloadProduct)

	cart.Product = preloadProduct.Data

	if err := validator.New().Struct(cart); err != nil {
		return cart, err
	}

	addCart, err := cs.cartRepo.AddToCart(cart)
	if err != nil {
		return cart, err
	}

	return addCart, nil
}

func (cs cartService) UpdateCart(cart entity.Cart, id uint) (entity.Cart, error) {
	var preloadProduct domain.PreloadProduct
	
	getCart, _ := cs.cartRepo.GetCart(cart, id)

	resp, err := http.Get(fmt.Sprintf("http://product-service:8081/product/%d", getCart.ProductID))
	if err != nil {
		return cart, fmt.Errorf("Error get product: %v", err)
	}

	json.NewDecoder(resp.Body).Decode(&preloadProduct)

	getCart.Product = preloadProduct.Data
	getCart.Item += cart.Item

	if err := validator.New().Struct(getCart); err != nil {
		return cart, err
	}
	
	updCart, err := cs.cartRepo.UpdateCart(getCart, id)
	if err != nil {
		return cart, err
	}

	return updCart, nil
}

func (cs cartService) DeleteProductInCart(cart entity.Cart, id uint) error {
	if err := cs.cartRepo.DeleteProductInCart(cart, id); err != nil {
		return err
	}

	return nil
}

func (cs cartService) GetCart(cart entity.Cart, id uint) (entity.Cart, error) {
	getCart, err := cs.cartRepo.GetCart(cart, id)
	if err != nil {
		return cart, err
	}

	var preloadProduct domain.PreloadProduct

	resp, err := http.Get(fmt.Sprintf("http://product-service:8081/product/%d", getCart.ProductID))
	if err != nil {
		return cart, err
	}

	json.NewDecoder(resp.Body).Decode(&preloadProduct)

	getCart.Product = preloadProduct.Data

	return getCart, nil
}

func (cs cartService) GetCartByUser(cart []entity.Cart, id uint) ([]entity.Cart, error) {
	getCart, err := cs.cartRepo.GetCartByUser(cart, id)
	if err != nil {
		return nil, err
	}

	for i := range getCart {
		var preloadProduct domain.PreloadProduct
	
		resp, err := http.Get(fmt.Sprintf("http://product-service:8081/product/%d", getCart[i].ProductID))
		if err != nil {
			return cart, fmt.Errorf("Error get product: %v", err)
		}
	
		json.NewDecoder(resp.Body).Decode(&preloadProduct)

		getCart[i].Product = preloadProduct.Data
	}

	return getCart, nil
}