package service

import (
	"errors"
	"github.com/NetSinx/yconnect-shop/product/app/model"
	"github.com/NetSinx/yconnect-shop/product/repository"
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
		return nil, errors.New("cannot get all products")
	}

	return product, nil
}

func (p productService) CreateProduct(products model.Product) error {
	if err := p.productRepository.CreateProduct(products); err != nil {
		return errors.New("cannot created data")
	}

	return nil
}

func (p productService) UpdateProduct(products model.Product, slug string) error {
	if err := p.productRepository.UpdateProduct(products, slug); err != nil {
		return errors.New("cannot updated data")
	}

	return nil
}

func (p productService) DeleteProduct(products model.Product, slug string) error {
	if err := p.productRepository.DeleteProduct(products, slug); err != nil {
		return errors.New("cannot deleted data")
	}

	return nil
}

func (p productService) GetProduct(products model.Product, slug string) (model.Product, error) {
	getProducts, err := p.productRepository.GetProduct(products, slug); if err != nil {
		return products, err
	}

	return getProducts, nil
}