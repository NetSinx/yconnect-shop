package http

import (
	"github.com/NetSinx/yconnect-shop/server/product/internal/entity"
	"github.com/NetSinx/yconnect-shop/server/product/internal/model"
	"github.com/NetSinx/yconnect-shop/server/product/internal/usecase"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"sync"
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
	var wg sync.WaitGroup
	errCh := make(chan error, 1)
	files := form.File["gambar"]
	for _, file := range files {
		wg.Add(1)
		go func(f *multipart.FileHeader) {
			defer wg.Done()
			src, err := file.Open()
			if err != nil {
				p.Log.WithError(err).Error("error opening file")
				errCh <- err
				return
			}
			defer src.Close()
	
			dst, err := os.Create(file.Filename)
			if err != nil {
				p.Log.WithError(err).Error("error creating file")
				errCh <- err
				return
			}
			defer dst.Close()
	
			if _, err := io.Copy(dst, src); err != nil {
				p.Log.WithError(err).Error("error copying file to destination")
				errCh <- err
				return
			}
	
			productRequest.Gambar = append(productRequest.Gambar, entity.Gambar{
				Path: file.Filename,
			})
		}(file)
	}

	go func() {
		wg.Wait()
		close(errCh)	
	}()

	if err := <-errCh; err != nil {
		return err
	}

	response, err := p.ProductUseCase.CreateProduct(c.Request().Context(), productRequest)
	if err != nil {
		p.Log.WithError(err).Error("error creating product")
		return err
	}

	return c.JSON(http.StatusOK, response)
}

func (p *ProductController) UpdateProduct(c echo.Context) error {
	slug := c.Param("slug")

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

	var wg sync.WaitGroup
	errCh := make(chan error, 1)
	files := form.File["gambar"]
	for _, file := range files {
		wg.Add(1)
		go func(f *multipart.FileHeader) {
			defer wg.Done()
			src, err := f.Open()
			if err != nil {
				p.Log.WithError(err).Error("error opening uploaded file")
				errCh <- err
				return
			}
			defer src.Close()

			dst, err := os.Create(file.Filename)
			if err != nil {
				p.Log.WithError(err).Error("error creating uploaded file")
				errCh <- err
				return
			}
			defer dst.Close()

			if _, err := io.Copy(dst, src); err != nil {
				p.Log.WithError(err).Error("error copying file to destination")
				errCh <- err
				return
			}

			productRequest.Gambar = append(productRequest.Gambar, entity.Gambar{
				Path: file.Filename,
			})
		}(file)
	}

	go func() {
		wg.Wait()
		close(errCh)
	}()

	if err := <-errCh; err != nil {
		return err
	}

	response, err := p.ProductUseCase.UpdateProduct(c.Request().Context(), productRequest, slug)
	if err != nil {
		p.Log.WithError(err).Error("error updating product")
		return err
	}

	return c.JSON(http.StatusOK, response)
}

func (p *ProductController) DeleteProduct(c echo.Context) error {
	productRequest := new(model.ParamRequest)
	if err := c.Bind(productRequest); err != nil {
		p.Log.WithError(err).Error("error binding request")
		return err
	}

	if err := p.ProductUseCase.DeleteProduct(c.Request().Context(), productRequest.Slug); err != nil {
		p.Log.WithError(err).Error("error deleting product")
		return echo.ErrInternalServerError
	}

	return c.NoContent(http.StatusNoContent)
}

func (p *ProductController) GetProductBySlug(c echo.Context) error {
	productRequest := new(model.ParamRequest)
	if err := c.Bind(productRequest); err != nil {
		p.Log.WithError(err).Error("error binding request")
		return err
	}

	response, err := p.ProductUseCase.GetProductBySlug(c.Request().Context(), productRequest.Slug)
	if err != nil {
		p.Log.WithError(err).Error("error getting product")
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, response)
}

func (p *ProductController) GetCategoryProduct(c echo.Context) error {
	productRequest := new(model.ParamRequest)
	if err := c.Bind(productRequest); err != nil {
		p.Log.WithError(err).Error("error binding request")
		return err
	}

	response, err := p.ProductUseCase.GetCategoryProduct(c.Request().Context(), productRequest.Slug)
	if err != nil {
		p.Log.WithError(err).Error("error getting product")
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, response)
}

func (p *ProductController) GetProductByCategory(c echo.Context) error {
	productRequestParam := new(model.ParamRequest)
	if err := c.Bind(productRequestParam); err != nil {
		p.Log.WithError(err).Error("error binding request")
		return err
	}

	productRequest := new(model.GetAllProductRequest)
	if err := c.Bind(productRequest); err != nil {
		p.Log.WithError(err).Error("error binding request")
		return err
	}

	response, err := p.ProductUseCase.GetProductByCategory(c.Request().Context(), productRequest, productRequestParam.Slug)
	if err != nil {
		p.Log.WithError(err).Error("error getting product")
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, response)
}
