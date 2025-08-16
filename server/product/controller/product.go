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
	var product entity.Product

	if err := c.Bind(&product); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
			Message: err.Error(),
		})
	}

	err := p.productService.CreateProduct(product)
	if err != nil && err.Error() == "produk sudah tersedia" {
		return echo.NewHTTPError(http.StatusConflict, domain.MessageResp{
			Message: "Produk sudah tersedia.",
		})
	} else if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
			Message: err.Error(),
		})
	} 

	return c.JSON(http.StatusOK, domain.MessageResp{
		Message: "Produk berhasil ditambahkan.",
	})
}

func (p productController) UpdateProduct(c echo.Context) error {
	var product entity.Product

	slug := c.Param("slug")

	if err := c.Bind(&product); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
			Message: err.Error(),
		})
	}

	err := p.productService.UpdateProduct(product, slug)
	if err != nil && err.Error() == "produk tidak ditemukan" {
		return echo.NewHTTPError(http.StatusNotFound, domain.MessageResp{
			Message: "Produk tidak ditemukan.",
		})
	} else if err != nil && err.Error() == "produk sudah tersedia" {
		return echo.NewHTTPError(http.StatusConflict, domain.MessageResp{
			Message: "Produk sudah tersedia.",
		})
	} else if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
			Message: err.Error(),
		})
	} 

	return c.JSON(http.StatusOK, domain.MessageResp{
		Message: "Produk berhasil diubah.",
	})
}

func (p productController) DeleteProduct(c echo.Context) error {
	var product entity.Product

	slug := c.Param("slug")

	err := p.productService.DeleteProduct(product, slug)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, domain.MessageResp{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, domain.MessageResp{
		Message: "Produk berhasil dihapus.",
	})
}

func (p productController) GetProductByID(c echo.Context) error {
	var product entity.Product

	id := c.Param("id")

	getProduct, err := p.productService.GetProductByID(product, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, domain.MessageResp{
			Message: "Produk tidak ditemukan",
		})
	}
	
	return c.JSON(http.StatusOK, domain.RespData{
		Data: getProduct,
	})
}

func (p productController) GetProductBySlug(c echo.Context) error {
	var product entity.Product

	slug := c.Param("slug")

	getProduct, err := p.productService.GetProductBySlug(product, slug)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, domain.MessageResp{
			Message: "Produk tidak ditemukan.",
		})
	}
	
	return c.JSON(http.StatusOK, domain.RespData{
		Data: getProduct,
	})
}

func (p productController) GetCategoryProduct(c echo.Context) error {
	var product entity.Product

	slug := c.Param("slug")

	getCategoryProduct, err := p.productService.GetCategoryProduct(product, slug)
	if err != nil && err.Error() == "produk tidak ditemukan" {
		return echo.NewHTTPError(http.StatusNotFound, domain.MessageResp{
			Message: "Produk tidak ditemukan.",
		})
	} else if err != nil && err.Error() == "service kategori sedang bermasalah" {
		return echo.NewHTTPError(http.StatusInternalServerError, domain.MessageResp{
			Message: "Service kategori sedang bermasalah.",
		})
	}

	return c.JSON(http.StatusOK, domain.RespData{
		Data: getCategoryProduct,
	})
}