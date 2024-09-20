package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"github.com/NetSinx/yconnect-shop/server/category/model/domain"
	"github.com/NetSinx/yconnect-shop/server/category/model/entity"
	"github.com/NetSinx/yconnect-shop/server/category/repository"
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

func (c categoryService) ListCategory(categories []entity.Kategori) ([]entity.Kategori, error) {
	categories, err := c.categoryRepo.ListCategory(categories)
	if err != nil {
		return nil, err
	}

	for i, category := range categories {
		var preloadProduct domain.PreloadProducts

		responseData, err := http.Get(fmt.Sprintf("http://product-service:8081/product/category/%d", category.Id))
		if err != nil || responseData.StatusCode != 200 {
			return categories, nil
		}
		
		json.NewDecoder(responseData.Body).Decode(&preloadProduct)
		
		categories[i].Product = preloadProduct.Data
	}

	return categories, nil
}

func (c categoryService) CreateCategory(category entity.Kategori) (entity.Kategori, error) {
	category.Name = strings.Title(category.Name)
	category.Slug = strings.ToLower(category.Slug)

	if err := validator.New().Struct(category); err != nil {
		return category, err
	}

	createCategory, err := c.categoryRepo.CreateCategory(category)
	if err != nil {
		return category, err
	}

	return createCategory, nil
}

func (c categoryService) UpdateCategory(category entity.Kategori, id string) (entity.Kategori, error) {
	category.Name = strings.Title(category.Name)
	category.Slug = strings.ToLower(category.Slug)
	
	if err := validator.New().Struct(category); err != nil {
		return category, err
	}

	updCategory, err := c.categoryRepo.UpdateCategory(category, id)
	if err != nil {
		return category, err
	}

	return updCategory, nil
}

func (c categoryService) DeleteCategory(category entity.Kategori, id string) error {
	if err := c.categoryRepo.DeleteCategory(category, id); err != nil {
		return err
	}

	return nil
}

func (c categoryService) GetCategory(category entity.Kategori, id string) (entity.Kategori, error) {
	var preloadProduct domain.PreloadProducts
	
	getCategory, err := c.categoryRepo.GetCategory(category, id)
	if err != nil {
		return getCategory, err
	}
	
	responseData, err := http.Get(fmt.Sprintf("http://product-service:8081/product/category/%d", getCategory.Id))
	if err != nil || responseData.StatusCode != 200 {
		return getCategory, nil
	}
	json.NewDecoder(responseData.Body).Decode(&preloadProduct)

	getCategory.Product = preloadProduct.Data
	
	return getCategory, nil
}