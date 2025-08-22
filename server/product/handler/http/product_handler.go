package http

import (
	"errors"
	"net/http"
	"net/url"
	"strings"
	"github.com/NetSinx/yconnect-shop/server/product/errs"
	"github.com/NetSinx/yconnect-shop/server/product/handler/dto"
	"github.com/NetSinx/yconnect-shop/server/product/model"
	"github.com/NetSinx/yconnect-shop/server/product/service"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ProductHandl interface {
	ListProduct(c echo.Context) error
	CreateProduct(c echo.Context) error
	UpdateProduct(c echo.Context) error
	DeleteProduct(c echo.Context) error
	GetProductByID(c echo.Context) error
	GetProductBySlug(c echo.Context) error
	GetCategoryProduct(c echo.Context) error
	GetProductByCategory(c echo.Context) error
}

type productHandler struct {
	productService service.ProductServ
}

func ProductHandler(prodService service.ProductServ) productHandler {
	return productHandler{
		productService: prodService,
	}
}

func (p productHandler) ListProduct(c echo.Context) error {
	var products []model.Product

	listProducts, err := p.productService.ListProduct(products)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, dto.MessageResp{
			Message: errs.ErrInternalServer,
		})
	}

	return c.JSON(http.StatusOK, dto.RespData{
		Data: listProducts,
	})
}

func (p productHandler) CreateProduct(c echo.Context) error {
	var productReq dto.ProductRequest

	if err := c.Bind(&productReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, dto.MessageResp{
			Message: err.Error(),
		})
	}

	err := p.productService.CreateProduct(productReq)
	if err != nil && errors.Is(err, gorm.ErrDuplicatedKey) {
		return echo.NewHTTPError(http.StatusConflict, dto.MessageResp{
			Message: errs.ErrDuplicatedKey,
		})
	} else if err != nil && strings.Contains(err.Error(), validator.ValidationErrors{}.Error()) {
		return echo.NewHTTPError(http.StatusBadRequest, dto.MessageResp{
			Message: err.Error(),
		})
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, dto.MessageResp{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.MessageResp{
		Message: dto.CreateResponse,
	})
}

func (p productHandler) UpdateProduct(c echo.Context) error {
	var productReq dto.ProductRequest

	slug := c.Param("slug")

	if err := c.Bind(&productReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, dto.MessageResp{
			Message: err.Error(),
		})
	}

	err := p.productService.UpdateProduct(productReq, slug)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, dto.MessageResp{
			Message: errs.ErrNotFound,
		})
	} else if err != nil && errors.Is(err, gorm.ErrDuplicatedKey) {
		return echo.NewHTTPError(http.StatusConflict, dto.MessageResp{
			Message: errs.ErrDuplicatedKey,
		})
	} else if err != nil && errors.Is(err, gorm.ErrForeignKeyViolated) {
		return echo.NewHTTPError(http.StatusBadRequest, dto.MessageResp{
			Message: err.Error(),
		})
	} else if err != nil && strings.Contains(err.Error(), validator.ValidationErrors{}.Error()) {
		return echo.NewHTTPError(http.StatusBadRequest, dto.MessageResp{
			Message: err.Error(),
		})
	} 

	return c.JSON(http.StatusOK, dto.MessageResp{
		Message: dto.UpdateResponse,
	})
}

func (p productHandler) DeleteProduct(c echo.Context) error {
	var product model.Product

	slug := c.Param("slug")

	err := p.productService.DeleteProduct(product, slug)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, dto.MessageResp{
			Message: errs.ErrNotFound,
		})
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, dto.MessageResp{
			Message: errs.ErrInternalServer,
		})
	}

	return c.JSON(http.StatusOK, dto.MessageResp{
		Message: dto.DeleteResponse,
	})
}

func (p productHandler) GetProductByID(c echo.Context) error {
	var product model.Product

	id := c.Param("id")

	getProduct, err := p.productService.GetProductByID(product, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, dto.MessageResp{
			Message: errs.ErrNotFound,
		})
	}
	
	return c.JSON(http.StatusOK, dto.RespData{
		Data: getProduct,
	})
}

func (p productHandler) GetProductBySlug(c echo.Context) error {
	var product model.Product

	slug := c.Param("slug")

	getProduct, err := p.productService.GetProductBySlug(product, slug)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, dto.MessageResp{
			Message: errs.ErrNotFound,
		})
	}
	
	return c.JSON(http.StatusOK, dto.RespData{
		Data: getProduct,
	})
}

func (p productHandler) GetCategoryProduct(c echo.Context) error {
	var product model.Product

	slug := c.Param("slug")

	getCategoryProduct, err := p.productService.GetCategoryProduct(product, slug)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, dto.MessageResp{
			Message: errs.ErrNotFound,
		})
	} else if err != nil && errors.Is(err, err.(*url.Error)) {
		return echo.NewHTTPError(http.StatusInternalServerError, dto.MessageResp{
			Message: "Service kategori sedang bermasalah.",
		})
	}

	return c.JSON(http.StatusOK, dto.RespData{
		Data: getCategoryProduct,
	})
}

func (p productHandler) GetProductByCategory(c echo.Context) error {
	var product []model.Product

	slug := c.Param("slug")

	getProductByCategory, err := p.productService.GetProductByCategory(product, slug)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusInternalServerError, dto.MessageResp{
			Message: errs.ErrNotFound,
		})
	}

	return c.JSON(http.StatusOK, dto.RespData{
		Data: getProductByCategory,
	})
}