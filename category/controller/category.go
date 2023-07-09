package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/NetSinx/yconnect-shop/category/app/model"
	"github.com/NetSinx/yconnect-shop/category/service"
	"github.com/NetSinx/yconnect-shop/category/utils"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type categoryController struct {
	categoryService service.CategoryServ
}

func CategoryController(categoryservice service.CategoryServ) categoryController {
	return categoryController{
		categoryService: categoryservice,
	}
}

func (cc categoryController) ListCategory(c echo.Context) error {
	var categories []model.Category

	listCategories, err := cc.categoryService.ListCategory(categories)
	if err != nil {
		fmt.Printf("Error message: %v", err)

		return echo.NewHTTPError(http.StatusInternalServerError, utils.ErrServer{
			Code: http.StatusInternalServerError,
			Status: "Internal Server Error",
			Message: "Sorry, there was a server failure!",
		})
	}

	for i, category := range listCategories {
		var preloadProduct utils.PreloadProducts

		responseData, err := http.Get(fmt.Sprintf("http://product-service:8081/products/category/%d", category.Id))
		if err != nil {
			return c.JSON(http.StatusOK, utils.SuccessGetData{
				Code: http.StatusOK,
				Status: "OK",
				Data: listCategories,
			})
		}

		if err := json.NewDecoder(responseData.Body).Decode(&preloadProduct); err != nil {
			return err
		}

		listCategories[i].Product = append(listCategories[i].Product, preloadProduct.Data...)
	}

	return c.JSON(http.StatusOK, utils.SuccessGetData{
		Code: http.StatusOK,
		Status: "OK",
		Data: listCategories,
	})
}

func (cc categoryController) CreateCategory(c echo.Context) error {
	var categories *model.Category

	if err := c.Bind(&categories); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrServer{
			Code: http.StatusBadRequest,
			Status: "Bad Request",
			Message: "Request doesn't match!",
		})
	}

	if err := validator.New().Struct(categories); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := cc.categoryService.CreateCategory(*categories); err != nil {
		fmt.Printf("Error message: %v", err)

		return echo.NewHTTPError(http.StatusConflict, utils.ErrServer{
			Code: http.StatusConflict,
			Status: "Data Conflict",
			Message: "Data was existing!",
		})
	}

	return c.JSON(http.StatusOK, utils.SuccessCUD{
		Code: http.StatusOK,
		Status: "OK",
		Message: "Category created successfully!",
	})
}

func (cc categoryController) UpdateCategory(c echo.Context) error {
	var categories *model.Category

	if err := c.Bind(&categories); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrServer{
			Code: http.StatusBadRequest,
			Status: "Bad Request",
			Message: "Request doesn't match!",
		})
	}
	
	if err := validator.New().Struct(categories); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	slug := c.Param("slug")

	if err := cc.categoryService.UpdateCategory(*categories, slug); (err != nil && err.Error() == "record not found") {
		fmt.Printf("Error message: %v", err.Error())

		return echo.NewHTTPError(http.StatusNotFound, utils.ErrServer{
			Code: http.StatusNotFound,
			Status: "Not Found",
			Message: "Category cannot be found!",
		})
	} 
	
	if err := cc.categoryService.UpdateCategory(*categories, slug); (err != nil && err.Error() != "record not found") {
		fmt.Printf("Error message: %v", err.Error())

		return echo.NewHTTPError(http.StatusConflict, utils.ErrServer{
			Code: http.StatusConflict,
			Status: "Data Conflict",
			Message: "Data was existing!",
		})
	}
	
	return c.JSON(http.StatusOK, utils.SuccessCUD{
		Code: http.StatusOK,
		Status: "OK",
		Message: "Category updated successfully!",
	})
}

func (cc categoryController) DeleteCategory(c echo.Context) error {
	var category model.Category

	slug := c.Param("slug")

	if err := cc.categoryService.DeleteCategory(category, slug); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, utils.ErrServer{
			Code: http.StatusNotFound,
			Status: "Not Found",
			Message: "Category cannot be found!",
		})
	}

	return c.JSON(http.StatusOK, utils.SuccessCUD{
		Code: http.StatusOK,
		Status: "OK",
		Message: "Category deleted successfully!",
	})
}

func (cc categoryController) GetCategory(c echo.Context) error {
	var categories model.Category
	var preloadProduct utils.PreloadProducts

	slug := c.Param("slug")

	getCategory, err := cc.categoryService.GetCategory(categories, slug); if err != nil {
		fmt.Printf("Error message: %v", err)

		return echo.NewHTTPError(http.StatusNotFound, utils.ErrServer{
			Code: http.StatusNotFound,
			Status: "Not Found",
			Message: "Category cannot be found!",
		})
	}

	responseData, err := http.Get(fmt.Sprintf("http://kong-gateway:8000/products/category/%d", getCategory.Id))
	if err != nil {
		return c.JSON(http.StatusOK, utils.SuccessGetData{
			Code: http.StatusOK,
			Status: "OK",
			Data: getCategory,
		})
	}

	if err := json.NewDecoder(responseData.Body).Decode(&preloadProduct); err != nil {
		return err
	}

	getCategory.Product = append(getCategory.Product, preloadProduct.Data...)

	return c.JSON(http.StatusOK, utils.SuccessGetData{
		Code: http.StatusOK,
		Status: "OK",
		Data: getCategory,
	})
}