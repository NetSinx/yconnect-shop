package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/NetSinx/yconnect-shop/server/product/model/entity"
	"github.com/NetSinx/yconnect-shop/server/product/repository"
	"github.com/NetSinx/yconnect-shop/server/product/model/domain"
	"github.com/go-playground/validator/v10"
)

type productService struct {
	productRepository repository.ProductRepo
}

func ProductService(prodRepo repository.ProductRepo) productService {
	return productService{
		productRepository: prodRepo,
	}
}

func (p productService) ListProduct(products []entity.Product) ([]entity.Product, error) {
	product, err := p.productRepository.ListProduct(products)
	if err != nil {
		return nil, err
	}

	for i, prod := range product {
		var preloadCategory domain.PreloadCategory
		var preloadUser domain.PreloadUser

		resCategory, err := http.Get(fmt.Sprintf("http://category-service:8080/category/%d", prod.CategoryId))
		if err != nil {
			product[i].Category = preloadCategory.Data
		} else if resCategory.StatusCode == 200 {
			json.NewDecoder(resCategory.Body).Decode(&preloadCategory)

			product[i].Category = preloadCategory.Data
		}
		
		resUser, err := http.Get(fmt.Sprintf("http://user-service:8082/user/%d", prod.SellerId))
		if err != nil {
			product[i].Seller = preloadUser.Data
		} else if resUser.StatusCode == 200 {
			json.NewDecoder(resUser.Body).Decode(&preloadUser)
			
			product[i].Seller = preloadUser.Data
		}
	}

	return product, nil
}

func (p productService) CreateProduct(products entity.Product, image []entity.Image) (entity.Product, error) {
	if err := validator.New().Struct(products); err != nil {
		return products, err
	}

	product, err := p.productRepository.CreateProduct(products, image)
	if err != nil {
		return products, err
	}

	return product, nil
}

func (p productService) UpdateProduct(products entity.Product, image []entity.Image, slug string, id string) (entity.Product, error) {
	if err := validator.New().Struct(products); err != nil {
		return products, err
	}

	product, err := p.productRepository.UpdateProduct(products, image, slug, id)
	if err != nil {
		return products, err
	}

	return product, nil
}

func (p productService) DeleteProduct(products entity.Product, image []entity.Image, slug string, id string) error {
	err := p.productRepository.DeleteProduct(products, image, slug, id)
	if err != nil {
		return err
	}

	return nil
}

func (p productService) GetProduct(products entity.Product, slug string) (entity.Product, error) {
	var preloadCategory domain.PreloadCategory
	var preloadUser domain.PreloadUser

	getProducts, err := p.productRepository.GetProduct(products, slug)
	if err != nil {
		return products, err
	}
	
	resCategory, err := http.Get(fmt.Sprintf("http://category-service:8080/category/%d", getProducts.CategoryId))
	if err != nil {
		getProducts.Category = preloadCategory.Data
	} else if resCategory.StatusCode == 200 {
		json.NewDecoder(resCategory.Body).Decode(&preloadCategory)

		getProducts.Category = preloadCategory.Data
	}

	resUser, err := http.Get(fmt.Sprintf("http://user-service:8082/seller/%d", getProducts.SellerId))
	if err != nil {
		getProducts.Seller = preloadUser.Data
	} else if resUser.StatusCode == 200 {
		json.NewDecoder(resUser.Body).Decode(&preloadUser)
	
		getProducts.Seller = preloadUser.Data
	}

	return getProducts, nil
}

func (p productService) GetProductByCategory(products []entity.Product, id string) ([]entity.Product, error) {
	getProdByCate, err := p.productRepository.GetProductByCategory(products, id)
	if err != nil {
		return nil, err
	}

	return getProdByCate, nil
}

func (p productService) GetProductBySeller(products []entity.Product, id string) ([]entity.Product, error) {
	getProdBySeller, err := p.productRepository.GetProductBySeller(products, id)
	if err != nil {
		return nil, err
	}

	return getProdBySeller, nil
}