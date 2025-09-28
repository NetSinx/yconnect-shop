package http

import (
	"errors"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/NetSinx/yconnect-shop/server/product/internal/entity"
	"github.com/NetSinx/yconnect-shop/server/product/internal/model"
	"github.com/NetSinx/yconnect-shop/server/product/internal/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProductController struct {
	Log            *logrus.Logger
	ProductUseCase *usecase.ProductUseCase
}

func NewProductController(log *logrus.Logger, productUseCase *usecase.ProductUseCase) *ProductController {
	return &ProductController{
		Log:            log,
		ProductUseCase: productUseCase,
	}
}

func (p *ProductController) GetAllProduct(c echo.Context) error {
	productRequest := new(model.GetAllProductRequest)
	if err := c.Bind(productRequest); err != nil {
		p.Log.WithError(err).Error("error binding request")
		return err
	}

	response, err := p.ProductUseCase.GetAllProduct(c.Request().Context(), productRequest)
	if err != nil {
		p.Log.WithError(err).Error("error getting all products")
		return err
	}

	return c.JSON(http.StatusOK, response)
}

func (p *ProductController) CreateProduct(c echo.Context) error {
	productRequest := new(model.ProductRequest)
	if err := c.Bind(productRequest); err != nil {
		p.Log.WithError(err).Error("error binding request")
		return err
	}

	form, err := c.MultipartForm()
	if err != nil {
		p.Log.WithError(err).Error("error uploading file")
		return err
	}
	files := form.File["gambar"]

	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			p.Log.WithError(err).Error("error opening file")
			return err
		}
		defer src.Close()
	
		dst, err := os.Create(file.Filename)
		if err != nil {
			p.Log.WithError(err).Error("error creating file")
			return err
		}
		defer dst.Close()
	
		if _, err := io.Copy(dst, src); err != nil {
			p.Log.WithError(err).Error("error copying file to destination")
			return err
		}

		productRequest.Gambar = append(productRequest.Gambar, entity.Gambar{
			Path: file.Filename,
		})
	}

	response, err := p.ProductUseCase.CreateProduct(c.Request().Context(), productRequest)
	if err != nil {
		p.Log.WithError(err).Error("error creating product")
		return err
	}

	return c.JSON(http.StatusOK, response)
}

func (p *ProductController) UpdateProduct(c echo.Context) error {
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

func (p *ProductController) DeleteProduct(c echo.Context) error {
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

func (p *ProductController) GetProductByID(c echo.Context) error {
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

func (p *ProductController) GetProductBySlug(c echo.Context) error {
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

func (p *ProductController) GetCategoryProduct(c echo.Context) error {
	var product model.Product

	slug := c.Param("slug")

	getCategoryProduct, err := p.productService.GetCategoryProduct(product, slug)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, dto.MessageResp{
			Message: errs.ErrNotFound,
		})
	}

	return c.JSON(http.StatusOK, dto.RespData{
		Data: getCategoryProduct,
	})
}

func (p *ProductController) GetProductByCategory(c echo.Context) error {
	var product []model.Product

	slug := c.Param("slug")

	getProductByCategory, err := p.productService.GetProductByCategory(product, slug)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, dto.MessageResp{
			Message: errs.ErrNotFound,
		})
	}

	return c.JSON(http.StatusOK, dto.RespData{
		Data: getProductByCategory,
	})
}
