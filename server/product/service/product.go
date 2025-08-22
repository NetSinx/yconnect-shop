package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	categoryEntity "github.com/NetSinx/yconnect-shop/server/category/model"
	"github.com/NetSinx/yconnect-shop/server/product/handler/dto"
	"github.com/NetSinx/yconnect-shop/server/product/helpers"
	"github.com/NetSinx/yconnect-shop/server/product/model"
	"github.com/NetSinx/yconnect-shop/server/product/repository"
	"github.com/go-playground/validator/v10"
)

type ProductServ interface {
	ListProduct(products []model.Product) ([]model.Product, error)
	CreateProduct(productReq dto.ProductRequest) error
	UpdateProduct(productReq dto.ProductRequest, slug string) error
	DeleteProduct(product model.Product, slug string) error
	GetProductByID(product model.Product, id string) (model.Product, error)
	GetProductBySlug(product model.Product, slug string) (model.Product, error)
	GetCategoryProduct(product model.Product, slug string) (categoryEntity.Category, error)
	GetProductByCategory(products []model.Product, slug string) ([]model.Product, error)
}

type productService struct {
	productRepository repository.ProductRepo
}

func ProductService(prodRepo repository.ProductRepo) productService {
	return productService{
		productRepository: prodRepo,
	}
}

func (p productService) ListProduct(products []model.Product) ([]model.Product, error) {
	listProduct, err := p.productRepository.ListProduct(products)
	if err != nil {
		return products, err
	}

	return listProduct, nil
}

func (p productService) CreateProduct(productReq dto.ProductRequest) error {
	if err := validator.New().Struct(productReq); err != nil {
		return err
	}

	slug, err := helpers.GenerateSlugByName(productReq.Nama);
	if err != nil {
		return err
	}

	product := model.Product{
		Nama: productReq.Nama,
		Deskripsi: productReq.Deskripsi,
		Slug: slug,
		Gambar: productReq.Gambar,
		KategoriSlug: productReq.KategoriSlug,
		Harga: productReq.Harga,
		Stok: productReq.Stok,
	}

	err = p.productRepository.CreateProduct(product)
	if err != nil {
		return err
	}

	return nil
}

func (p productService) UpdateProduct(productReq dto.ProductRequest, slug string) error {
	if err := validator.New().Struct(productReq); err != nil {
		return err
	}

	product := model.Product{
		Nama: productReq.Nama,
		Deskripsi: productReq.Deskripsi,
		Slug: slug,
		KategoriSlug: productReq.KategoriSlug,
		Harga: productReq.Harga,
		Stok: productReq.Stok,
	}

	gambar := productReq.Gambar

	getProductName, err := p.productRepository.GetProductName(product, slug)
	if err != nil {
		return err
	}

	if getProductName.Nama != productReq.Nama {
		newSlug := helpers.ReplaceProductSlug(getProductName.Slug, productReq.Nama)
		product.Slug = newSlug
	}

	err = p.productRepository.UpdateProduct(product, gambar, slug)
	if err != nil {
		return err
	}

	return nil
}

func (p productService) DeleteProduct(product model.Product, slug string) error {
	err := p.productRepository.DeleteProduct(product, slug)
	if err != nil {
		return err
	}

	return nil
}

func (p productService) GetProductByID(product model.Product, id string) (model.Product, error) {
	getProduct, err := p.productRepository.GetProductByID(product, id)
	if err != nil {
		return product, err
	}

	return getProduct, nil
}

func (p productService) GetProductBySlug(product model.Product, slug string) (model.Product, error) {
	getProduct, err := p.productRepository.GetProductBySlug(product, slug)
	if err != nil {
		return product, err
	}

	return getProduct, nil
}

func (p productService) GetCategoryProduct(product model.Product, slug string) (categoryEntity.Category, error) {
	if err := p.productRepository.GetCategoryProduct(product, slug); err != nil {
		return categoryEntity.Category{}, err
	}
	
	baseUrl := os.Getenv("BASE_URL")

	respCategory, err := http.Get(fmt.Sprintf(baseUrl + ":8080/category/%s", product.KategoriSlug))
	if err != nil {
		return categoryEntity.Category{}, err
	}

	var respDataCategory categoryEntity.Category
	json.NewDecoder(respCategory.Body).Decode(&respDataCategory)

	return respDataCategory, nil
}

func (p productService) GetProductByCategory(products []model.Product, slug string) ([]model.Product, error) {
	getProductByCategory, err := p.productRepository.GetProductByCategory(products, slug)
	if err != nil {
		return products, err
	}

	return getProductByCategory, nil
}