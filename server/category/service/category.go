package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/NetSinx/yconnect-shop/server/category/model/entity"
	"github.com/NetSinx/yconnect-shop/server/category/model/domain"
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

func (c categoryService) ListCategory(categories []entity.Category) ([]entity.Category, error) {
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

func (c categoryService) CreateCategory(categories entity.Category) (entity.Category, error) {
	if err := validator.New().Struct(categories); err != nil {
		return categories, err
	}

	category, err := c.categoryRepo.CreateCategory(categories)
	if err != nil {
		return categories, err
	}

	return category, nil
}

func (c categoryService) UpdateCategory(categories entity.Category, id string) (entity.Category, error) {
	if err := validator.New().Struct(categories); err != nil {
		return categories, err
	}

	category, err := c.categoryRepo.UpdateCategory(categories, id)
	if err != nil {
		return categories, err
	}

	return category, nil
}

func (c categoryService) DeleteCategory(category entity.Category, id string) error {
	if err := c.categoryRepo.DeleteCategory(category, id); err != nil {
		return err
	}

	return nil
}

func (c categoryService) GetCategory(categories entity.Category, id string) (entity.Category, error) {
	var preloadProduct domain.PreloadProducts
	
	getCategory, err := c.categoryRepo.GetCategory(categories, id)
	if err != nil {
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