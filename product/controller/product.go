package controller

import (
	"fmt"
	"net/http"

	"github.com/NetSinx/yconnect-shop/product/app/config"
	"github.com/NetSinx/yconnect-shop/product/app/model"
	"github.com/NetSinx/yconnect-shop/product/service"
	"github.com/NetSinx/yconnect-shop/product/utils"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type productController struct {
	productService service.ProductServ
}

func ProductController(prodService service.ProductServ) productController {
	return productController{
		productService: prodService,
	}
}

func (p productController) ListProduct(c echo.Context) error {
	var products []model.Product

	listProducts, err := p.productService.ListProduct(products)
	if err != nil {
		fmt.Printf("Error message: %v", err)

		return echo.NewHTTPError(http.StatusInternalServerError, utils.ErrServer{
			Code: http.StatusInternalServerError,
			Status: "Internal Server Error",
			Message: "Sorry, there was a server failure!",
		})
	}

	return c.JSON(http.StatusOK, utils.SuccessGetData{
		Code: http.StatusOK,
		Status: "OK",
		Data: listProducts,
	})
}

func (p productController) CreateProduct(c echo.Context) error {
	var products *model.Product
	
	if err := c.Bind(&products); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrServer{
			Code: http.StatusBadRequest,
			Status: "Bad Request",
			Message: "Request doesn't match!",
		})
	}

	if err := validator.New().Struct(products); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	
	if err := p.productService.CreateProduct(*products); err != nil {
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
		Message: "Product created successfully!",
	})
}

func (p productController) UpdateProduct(c echo.Context) error {
	var products *model.Product

	if err := c.Bind(&products); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrServer{
			Code: http.StatusBadRequest,
			Status: "Bad Request",
			Message: "Request doesn't match!",
		})
	}

	if err := validator.New().Struct(products); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	slug := c.Param("slug")

	if err := p.productService.UpdateProduct(*products, slug); err != nil {
		fmt.Printf("Error message: %v", err)

		return echo.NewHTTPError(http.StatusConflict, utils.SuccessCUD{
			Code: http.StatusConflict,
			Status: "Data Conflict",
			Message: "Data was existing!",
		})
	}

	return c.JSON(http.StatusOK, utils.SuccessCUD{
		Code: http.StatusOK,
		Status: "OK",
		Message: "Product updated successfully!",
	})
}

func (p productController) DeleteProduct(c echo.Context) error {
	var products model.Product

	slug := c.Param("slug")

	if err := p.productService.DeleteProduct(products, slug); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, utils.ErrServer{
			Code: http.StatusNotFound,
			Status: "Not Found",
			Message: "Product cannot be found!",
		})
	}

	return c.JSON(http.StatusOK, utils.SuccessCUD{
		Code: http.StatusOK,
		Status: "OK",
		Message: "Product deleted successfully!",
	})
}

func (p productController) GetProduct(c echo.Context) error {
	var products model.Product

	slug := c.Param("slug")

	getProducts, err := p.productService.GetProduct(products, slug)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, utils.ErrServer{
			Code: http.StatusNotFound,
			Status: "Not Found",
			Message: "Product cannot be found!",
		})
	}

	return c.JSON(http.StatusOK, utils.SuccessGetData{
		Code: http.StatusOK,
		Status: "OK",
		Data: getProducts,
	})
}

func (p productController) GetProductByCategory(c echo.Context) error {
	var products []model.Product

	id := c.Param("id")

	if err := config.DB.Find(&products, "category_id = ?", id).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, utils.ErrServer{
			Code: http.StatusInternalServerError,
			Status: http.StatusText(http.StatusInternalServerError),
			Message: "Maaf, ada kesalahan pada server!",
		})
	}

	return c.JSON(http.StatusOK, utils.SuccessGetData{
		Code: http.StatusOK,
		Status: "OK",
		Data: products,
	})
}

func (p productController) GetProductByUser(c echo.Context) error {
	var products []model.Product

	id := c.Param("id")

	if err := config.DB.Find(&products, "seller_id = ?", id).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, utils.ErrServer{
			Code: http.StatusInternalServerError,
			Status: http.StatusText(http.StatusInternalServerError),
			Message: "Maaf, ada kesalahan pada server!",
		})
	}

	return c.JSON(http.StatusOK, utils.SuccessGetData{
		Code: http.StatusOK,
		Status: "OK",
		Data: products,
	})
}