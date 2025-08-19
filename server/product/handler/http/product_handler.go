package http

import (
	"net/http"
	"github.com/NetSinx/yconnect-shop/server/product/model"
	"github.com/NetSinx/yconnect-shop/server/product/service"
	"github.com/NetSinx/yconnect-shop/server/product/handler/dto"
	"github.com/labstack/echo/v4"
)

type ProductHandl interface {
	ListProduct(c echo.Context) error
	CreateProduct(c echo.Context) error
	UpdateProduct(c echo.Context) error
	DeleteProduct(c echo.Context) error
	GetProductByID(c echo.Context) error
	GetProductBySlug(c echo.Context) error
	GetCategoryProduct(c echo.Context) error
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
			Message: err.Error(),
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
	if err != nil && err.Error() == "produk sudah tersedia" {
		return echo.NewHTTPError(http.StatusConflict, dto.MessageResp{
			Message: "Produk sudah tersedia.",
		})
	} else if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, dto.MessageResp{
			Message: err.Error(),
		})
	} 

	return c.JSON(http.StatusOK, dto.MessageResp{
		Message: "Produk berhasil ditambahkan.",
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
	if err != nil && err.Error() == "produk tidak ditemukan" {
		return echo.NewHTTPError(http.StatusNotFound, dto.MessageResp{
			Message: "Produk tidak ditemukan.",
		})
	} else if err != nil && err.Error() == "produk sudah tersedia" {
		return echo.NewHTTPError(http.StatusConflict, dto.MessageResp{
			Message: "Produk sudah tersedia.",
		})
	} else if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, dto.MessageResp{
			Message: err.Error(),
		})
	} 

	return c.JSON(http.StatusOK, dto.MessageResp{
		Message: "Produk berhasil diubah.",
	})
}

func (p productHandler) DeleteProduct(c echo.Context) error {
	var product model.Product

	slug := c.Param("slug")

	err := p.productService.DeleteProduct(product, slug)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, dto.MessageResp{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.MessageResp{
		Message: "Produk berhasil dihapus.",
	})
}

func (p productHandler) GetProductByID(c echo.Context) error {
	var product model.Product

	id := c.Param("id")

	getProduct, err := p.productService.GetProductByID(product, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, dto.MessageResp{
			Message: "Produk tidak ditemukan",
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
			Message: "Produk tidak ditemukan.",
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
	if err != nil && err.Error() == "produk tidak ditemukan" {
		return echo.NewHTTPError(http.StatusNotFound, dto.MessageResp{
			Message: "Produk tidak ditemukan.",
		})
	} else if err != nil && err.Error() == "service kategori sedang bermasalah" {
		return echo.NewHTTPError(http.StatusInternalServerError, dto.MessageResp{
			Message: "Service kategori sedang bermasalah.",
		})
	}

	return c.JSON(http.StatusOK, dto.RespData{
		Data: getCategoryProduct,
	})
}