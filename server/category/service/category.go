package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"gorm.io/gorm"
	"github.com/NetSinx/yconnect-shop/server/category/app/model"
	"github.com/NetSinx/yconnect-shop/server/category/repository"
	"github.com/NetSinx/yconnect-shop/server/category/utils"
	"github.com/go-playground/validator/v10"
)

type categoryService struct {
	categoryRepo repository.CategoryRepo
}

func CategoryService(categoryrepo repository.CategoryRepo) categoryService {
	return categoryService{
		categoryRepo: categoryrepo,
	}
}

func (c categoryService) ListCategory(categories []model.Category) ([]model.Category, error) {
	categories, err := c.categoryRepo.ListCategory(categories)
	if err != nil {
		return nil, err
	}

	for i, category := range categories {
		var preloadProduct utils.PreloadProducts

		responseData, err := http.Get(fmt.Sprintf("http://product-service:8081/product/category/%d", category.Id))
		if err != nil {
			return categories, nil
		} else if responseData.StatusCode == 200 {
			json.NewDecoder(responseData.Body).Decode(&preloadProduct)
			
			categories[i].Product = preloadProduct.Data
		}
	}

	return categories, nil
}

func (c categoryService) CreateCategory(categories model.Category) (model.Category, error) {
	if err := validator.New().Struct(categories); err != nil {
		return categories, errors.New("request tidak sesuai")
	}

	category, err := c.categoryRepo.CreateCategory(categories)
	if err != nil {
		return categories, errors.New("kategori sudah tersedia")
	}

	return category, nil
}

func (c categoryService) UpdateCategory(categories model.Category, id string) (model.Category, error) {
	if err := validator.New().Struct(categories); err != nil {
		return categories, errors.New("request tidak sesuai")
	}

	category, err := c.categoryRepo.UpdateCategory(categories, id)
	if err != nil && err != gorm.ErrRecordNotFound {
		return categories, errors.New("kategori sudah tersedia")
	} else if err != nil && err == gorm.ErrRecordNotFound {
		return categories, err
	}

	return category, nil
}

func (c categoryService) DeleteCategory(category model.Category, id string) error {
	if err := c.categoryRepo.DeleteCategory(category, id); err != nil {
		return err
	}

	return nil
}

func (c categoryService) GetCategory(categories model.Category, id string) (model.Category, error) {
	var preloadProduct utils.PreloadProducts
	
	getCategory, err := c.categoryRepo.GetCategory(categories, id); if err != nil {
		return getCategory, err
	}
	
	responseData, err := http.Get(fmt.Sprintf("http://product-service:8081/product/category/%d", getCategory.Id))
	if err != nil {
		return getCategory, nil
	} else if responseData.StatusCode == 200 {
		json.NewDecoder(responseData.Body).Decode(&preloadProduct)

		getCategory.Product = preloadProduct.Data
	}
	
	return getCategory, nil
}