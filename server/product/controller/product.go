package controller

import (
	"net/http"
	"github.com/NetSinx/yconnect-shop/server/product/model/domain"
	"github.com/NetSinx/yconnect-shop/server/product/model/entity"
	"github.com/NetSinx/yconnect-shop/server/product/service"
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
	var products []entity.Product

	listProducts, err := p.productService.ListProduct(products)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, domain.MessageResp{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, domain.RespData{
		Data: listProducts,
	})
}

func (p productController) CreateProduct(c echo.Context) error {
	var products entity.Product

	if err := c.Bind(&products); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
			Message: err.Error(),
		})
	}
	
	imageProduct, err := c.MultipartForm()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, domain.MessageResp{
			Message: err.Error(),
		})
	}
	
	product, err := p.productService.CreateProduct(products, imageProduct.File["images"])
	if err != nil && err.Error() == "produk sudah tersedia" {
		return echo.NewHTTPError(http.StatusConflict, domain.MessageResp{
			Message: err.Error(),
		})
	} else if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
			Message: err.Error(),
		})
	} 

	return c.JSON(http.StatusOK, domain.RespData{
		Data: product,
	})
}

func (p productController) UpdateProduct(c echo.Context) error {
	var products entity.Product

	slug := c.Param("slug")

	if err := c.Bind(&products); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
			Message: err.Error(),
		})
	}

	imageProduct, err := c.MultipartForm()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, domain.MessageResp{
			Message: err.Error(),
		})
	}

	product, err := p.productService.UpdateProduct(products, imageProduct.File["images"], slug)
	if err != nil && err.Error() == "produk tidak ditemukan" {
		return echo.NewHTTPError(http.StatusNotFound, domain.MessageResp{
			Message: err.Error(),
		})
	} else if err != nil && err.Error() == "produk sudah tersedia" {
		return echo.NewHTTPError(http.StatusConflict, domain.MessageResp{
			Message: err.Error(),
		})
	} else if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
			Message: err.Error(),
		})
	} 

	return c.JSON(http.StatusOK, domain.RespData{
		Data: product,
	})
}

func (p productController) DeleteProduct(c echo.Context) error {
	var products entity.Product

	slug := c.Param("slug")

	err := p.productService.DeleteProduct(products, slug)
	if err != nil && err.Error() == "produk tidak ditemukan" {
		return echo.NewHTTPError(http.StatusNotFound, domain.MessageResp{
			Message: err.Error(),
		})
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, domain.MessageResp{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, domain.MessageResp{
		Message: "Produk berhasil dihapus.",
	})
}

func (p productController) GetProduct(c echo.Context) error {
	var product entity.Product

	slug := c.Param("slug")

	getProduct, err := p.productService.GetProduct(product, slug)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, domain.MessageResp{
			Message: "Produk tidak ditemukan.",
		})
	}
	
	return c.JSON(http.StatusOK, domain.RespData{
		Data: getProduct,
	})
}

func (p productController) GetProductByCategory(c echo.Context) error {
	var products []entity.Product

	id := c.Param("id")

	getProdByCate, err := p.productService.GetProductByCategory(products, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, domain.MessageResp{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, domain.RespData{
		Data: getProdByCate,
	})
}