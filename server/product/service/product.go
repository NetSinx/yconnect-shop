package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"github.com/NetSinx/yconnect-shop/server/product/app/model"
	"github.com/NetSinx/yconnect-shop/server/product/repository"
	"github.com/NetSinx/yconnect-shop/server/product/utils"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type productService struct {
	productRepository repository.ProductRepo
}

func ProductService(prodRepo repository.ProductRepo) productService {
	return productService{
		productRepository: prodRepo,
	}
}

func (p productService) ListProduct(products []model.Product) ([]model.Product, error) {
	product, err := p.productRepository.ListProduct(products)
	if err != nil {
		return nil, err
	}

	for i, prod := range product {
		var preloadCategory utils.PreloadCategory
		var preloadUser utils.PreloadUser

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

func (p productService) CreateProduct(products model.Product, image []model.Image) (model.Product, error) {
	if err := validator.New().Struct(products); err != nil {
		return products, errors.New("request tidak sesuai")
	}

	product, err := p.productRepository.CreateProduct(products, image)
	if err != nil {
		return products, errors.New("produk sudah tersedia")
	}

	return product, nil
}

func (p productService) UpdateProduct(products model.Product, image []model.Image, id uint) (model.Product, error) {
	if err := validator.New().Struct(products); err != nil {
		return products, errors.New("request tidak sesuai")
	}

	product, err := p.productRepository.UpdateProduct(products, image, id)
	if err != nil && err == gorm.ErrRecordNotFound {
		return products, errors.New("produk tidak bisa ditemukan")
	} else if err != nil {
		return products, err
	}

	return product, nil
}

func (p productService) DeleteProduct(products model.Product, image []model.Image, id string) error {
	err := p.productRepository.DeleteProduct(products, image, id)
	if err != nil && err == gorm.ErrRecordNotFound {
		return errors.New("produk tidak ditemukan")
	} else if err != nil {
		return err
	}

	return nil
}

func (p productService) GetProduct(products model.Product, id string) (model.Product, error) {
	var preloadCategory utils.PreloadCategory
	var preloadUser utils.PreloadUser

	getProducts, err := p.productRepository.GetProduct(products, id)
	if err != nil && err == gorm.ErrRecordNotFound {
		return products, errors.New("produk tidak ditemukan")
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