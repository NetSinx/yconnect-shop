package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	categoryEntity "github.com/NetSinx/yconnect-shop/server/category/model/entity"
	"github.com/NetSinx/yconnect-shop/server/product/model/domain"
	"github.com/NetSinx/yconnect-shop/server/product/model/entity"
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

func (p productService) ListProduct(products []entity.Product) ([]entity.Product, error) {
	listProduct, err := p.productRepository.ListProduct(products)
	if err != nil {
		return products, err
	}

	return listProduct, nil
}

func (p productService) CreateProduct(productReq domain.ProductRequest) error {
	if err := validator.New().Struct(productReq); err != nil {
		return err
	}

	slug := utils.GenerateSlugByName(productReq.Nama)
	product := entity.Product{
		Nama: productReq.Nama,
		Deskripsi: productReq.Deskripsi,
		Slug: slug,
		Gambar: productReq.Gambar,
		KategoriID: productReq.KategoriID,
		Harga: productReq.Harga,
		Stok: productReq.Stok,
	}

	err := p.productRepository.CreateProduct(product)
	if err != nil {
		return fmt.Errorf("produk sudah tersedia")
	}

	return nil
}

func (p productService) UpdateProduct(productReq domain.ProductRequest, slug string) error {
	if err := validator.New().Struct(productReq); err != nil {
		return err
	}

	err := p.productRepository.UpdateProduct(productReq, slug)
	if err != nil && err == gorm.ErrRecordNotFound {
		return fmt.Errorf("produk tidak ditemukan")
	} else if err != nil && err == gorm.ErrDuplicatedKey {
		return fmt.Errorf("produk sudah tersedia")
	}

	return nil
}

func (p productService) DeleteProduct(product entity.Product, slug string) error {
	err := p.productRepository.DeleteProduct(product, slug)
	if err != nil {
		return err
	}

	return nil
}

func (p productService) GetProductByID(product entity.Product, id string) (entity.Product, error) {
	getProduct, err := p.productRepository.GetProductByID(product, id)
	if err != nil {
		return product, err
	}

	return getProduct, nil
}

func (p productService) GetProductBySlug(product entity.Product, slug string) (entity.Product, error) {
	getProduct, err := p.productRepository.GetProductBySlug(product, slug)
	if err != nil {
		return product, err
	}

	return getProduct, nil
}

func (p productService) GetCategoryProduct(product entity.Product, slug string) (categoryEntity.Category, error) {
	if err := p.productRepository.GetCategoryProduct(product, slug); err != nil {
		return categoryEntity.Category{}, fmt.Errorf("produk tidak ditemukan")
	}
	
	baseUrl := os.Getenv("BASE_URL")

	respCategory, err := http.Get(fmt.Sprintf(baseUrl + ":8080/category/%d", product.KategoriID))
	if err != nil {
		return categoryEntity.Category{}, fmt.Errorf("service kategori sedang bermasalah")
	}

	var respDataCategory categoryEntity.Category
	json.NewDecoder(respCategory.Body).Decode(&respDataCategory)

	return respDataCategory, nil
}