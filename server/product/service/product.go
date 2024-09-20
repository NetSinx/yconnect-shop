package service

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"github.com/NetSinx/yconnect-shop/server/product/model/domain"
	"github.com/NetSinx/yconnect-shop/server/product/model/entity"
	"github.com/NetSinx/yconnect-shop/server/product/repository"
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
	product, err := p.productRepository.ListProduct(products)
	if err != nil {
		return nil, err
	}

	for _, prod := range product {
		resCategory, err := http.Get(fmt.Sprintf("http://category-service:8080/category/%d", prod.KategoriId))
		if err != nil || resCategory.StatusCode != 200 {
			return product, nil
		}
		
		var preloadCategory domain.PreloadCategory
		
		json.NewDecoder(resCategory.Body).Decode(&preloadCategory)

		prod.Kategori = preloadCategory.Data
	}

	return product, nil
}

func (p productService) CreateProduct(products entity.Product, images []*multipart.FileHeader) (entity.Product, error) {
	var img []entity.Gambar
	
	for _, image := range images {
		src, err := image.Open()
		if err != nil {
			return entity.Product{}, err
		}
		defer src.Close()
	
		fileName := strings.Split(image.Filename, ".")[0]
		fileExt := strings.Split(image.Filename, ".")[1]
		hashedFileName := md5.New().Sum([]byte(fileName))
	
		if err := os.MkdirAll("../assets/images", os.ModePerm); err != nil {
			return entity.Product{}, err
		}
	
		dst, err := os.Create(fmt.Sprintf("../assets/images/%x.%s", hashedFileName, fileExt))
		if err != nil {
			return entity.Product{}, err
		}
		defer dst.Close()
	
		if _, err := io.Copy(dst, src); err != nil {
			return entity.Product{}, err
		}
	
		img = append(img, entity.Gambar{Nama: fmt.Sprintf("../assets/images/%x.%s", hashedFileName, fileExt)})
	}

	products.Gambar = img

	if err := validator.New().Struct(products); err != nil {
		return entity.Product{}, err
	}

	product, err := p.productRepository.CreateProduct(products)
	if err != nil {
		return entity.Product{}, fmt.Errorf("produk sudah tersedia")
	}

	return product, nil
}

func (p productService) UpdateProduct(products entity.Product, images []*multipart.FileHeader, slug string) (entity.Product, error) {
	getProduct, err := p.productRepository.GetProduct(products, slug)
	if err != nil {
		return entity.Product{}, err
	}

	for _, gambar := range getProduct.Gambar {
		os.Remove(gambar.Nama)
	}

	for _, image := range images {
		src, err := image.Open()
		if err != nil {
			return entity.Product{}, err
		}
		defer src.Close()
	
		fileName := strings.Split(image.Filename, ".")[0]
		fileExt := strings.Split(image.Filename, ".")[1]
		hashedFileName := md5.New().Sum([]byte(fileName))
	
		dst, err := os.Create(fmt.Sprintf("../assets/images/%x.%s", hashedFileName, fileExt))
		if err != nil {
			return entity.Product{}, err
		}
		defer dst.Close()
		
		if _, err := io.Copy(dst, src); err != nil {
			return entity.Product{}, err
		}

		getProduct.Gambar = append(getProduct.Gambar, entity.Gambar{Nama: fmt.Sprintf("../assets/images/%x.%s", hashedFileName, fileExt), ProductID: uint(getProduct.Id)})
	}

	products.Gambar = getProduct.Gambar

	if err := validator.New().Struct(products); err != nil {
		return entity.Product{}, err
	}

	product, err := p.productRepository.UpdateProduct(products, slug)
	if err != nil && err == gorm.ErrRecordNotFound {
		return entity.Product{}, fmt.Errorf("produk tidak ditemukan")
	} else if err != nil && err == gorm.ErrDuplicatedKey {
		return entity.Product{}, fmt.Errorf("produk sudah tersedia")
	}

	return product, nil
}

func (p productService) DeleteProduct(products entity.Product, slug string) error {
	getProduct, err := p.productRepository.GetProduct(products, slug)
	if err != nil {
		return fmt.Errorf("produk tidak ditemukan")
	}

	for _, image := range getProduct.Gambar {
		if err := os.Remove("." + image.Nama); err != nil {
			return err
		}
	}

	err = p.productRepository.DeleteProduct(products, slug)
	if err != nil && err == gorm.ErrRecordNotFound{
		return fmt.Errorf("produk tidak ditemukan")
	}

	return nil
}

func (p productService) GetProduct(products entity.Product, slug string) (entity.Product, error) {
	getProducts, err := p.productRepository.GetProduct(products, slug)
	if err != nil {
		return entity.Product{}, err
	}
	
	resCategory, err := http.Get(fmt.Sprintf("http://category-service:8080/category/%d", getProducts.KategoriId))
	if err != nil || resCategory.StatusCode != 200 {
		return getProducts, nil
	}

	var preloadCategory domain.PreloadCategory

	json.NewDecoder(resCategory.Body).Decode(&preloadCategory)

	getProducts.Kategori = preloadCategory.Data

	return getProducts, nil
}

func (p productService) GetProductByCategory(products []entity.Product, id string) ([]entity.Product, error) {
	getProdByCate, err := p.productRepository.GetProductByCategory(products, id)
	if err != nil {
		return nil, err
	}

	return getProdByCate, nil
}