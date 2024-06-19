package controller

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"github.com/NetSinx/yconnect-shop/server/product/model/domain"
	"github.com/NetSinx/yconnect-shop/server/product/model/entity"
	"github.com/NetSinx/yconnect-shop/server/product/service"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
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
	var img []entity.Image

	imageProduct, err := c.MultipartForm()
	if err != nil {
		return err
	}

	images := imageProduct.File["images"]

	for _, image := range images {
		src, err := image.Open()
		if err != nil {
			return err
		}
		defer src.Close()
	
		fileName := strings.Split(image.Filename, ".")[0]
		fileExt := strings.Split(image.Filename, ".")[1]
		hashedFileName := md5.New().Sum([]byte(fileName))
	
		if err := os.MkdirAll("assets/images", os.ModePerm); err != nil {
			return err
		}
	
		dst, err := os.Create(fmt.Sprintf("assets/images/%x.%s", hashedFileName, fileExt))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
				Message: err.Error(),
			})
		}
		defer dst.Close()
	
		if _, err := io.Copy(dst, src); err != nil {
			return err
		}
	
		img = append(img, entity.Image{Name: fmt.Sprintf("/assets/images/%x.%s", hashedFileName, fileExt)})
	}

	products.Image = img

	if err := c.Bind(&products); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
			Message: err.Error(),
		})
	}
	
	product, err := p.productService.CreateProduct(products, img)
	if err != nil && err == gorm.ErrDuplicatedKey {
		return echo.NewHTTPError(http.StatusConflict, domain.MessageResp{
			Message: "Produk sudah tersedia.",
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

	imageProduct, err := c.MultipartForm()
	if err != nil {
		return err
	}

	images := imageProduct.File["images"]

	getProduct, err := p.productService.GetProduct(products, slug)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, domain.MessageResp{
			Message: err.Error(),
		})
	}

	for i, image := range images {
		src, err := image.Open()
		if err != nil {
			return err
		}
		defer src.Close()
	
		fileName := strings.Split(image.Filename, ".")[0]
		fileExt := strings.Split(image.Filename, ".")[1]
		hashedFileName := md5.New().Sum([]byte(fileName))
	
		if (len(getProduct.Image) <= i) {
			dst, err := os.Create(fmt.Sprintf("assets/images/%x.%s", hashedFileName, fileExt))
			if err != nil {
				return err
			}
			defer dst.Close()
			
			if _, err := io.Copy(dst, src); err != nil {
				return err
			}

			getProduct.Image = append(getProduct.Image, entity.Image{Name: fmt.Sprintf("/assets/images/%x.%s", hashedFileName, fileExt), ProductID: uint(getProduct.Id)})

		} else {
			dst, err := os.Create(fmt.Sprintf("assets/images/%x.%s", hashedFileName, fileExt))
			if err != nil {
				return err
			}
			defer dst.Close()
			
			if _, err := io.Copy(dst, src); err != nil {
				return err
			}
			
			os.Remove("." + getProduct.Image[i].Name)

			getProduct.Image[i].Name = fmt.Sprintf("/assets/images/%x.%s", hashedFileName, fileExt)
			getProduct.Image[i].ProductID = uint(getProduct.Id)
		}
	}

	products.Image = getProduct.Image
	
	if err := c.Bind(&products); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
			Message: err.Error(),
		})
	}

	product, err := p.productService.UpdateProduct(products, products.Image, slug, fmt.Sprintf("%d", getProduct.Id))
	if err != nil && err == gorm.ErrRecordNotFound {
		return echo.NewHTTPError(http.StatusNotFound, domain.MessageResp{
			Message: "Produk tidak ditemukan.",
		})
	} else if err != nil && err == gorm.ErrDuplicatedKey {
		return echo.NewHTTPError(http.StatusConflict, domain.MessageResp{
			Message: "Produk sudah tersedia.",
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

	getProduct, err := p.productService.GetProduct(products, slug)
	if err != nil && err == gorm.ErrRecordNotFound {
		return echo.NewHTTPError(http.StatusNotFound, domain.MessageResp{
			Message: "Produk tidak ditemukan.",
		})
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, domain.MessageResp{
			Message: err.Error(),
		})
	}

	for _, image := range getProduct.Image {
		os.Remove("." + image.Name)
	}

	if err := p.productService.DeleteProduct(products, getProduct.Image, slug, fmt.Sprintf("%d", getProduct.Id)); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, domain.MessageResp{
			Message: "Produk tidak ditemukan.",
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
	if err != nil && err == gorm.ErrRecordNotFound{
		return echo.NewHTTPError(http.StatusNotFound, domain.MessageResp{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, domain.RespData{
		Data: getProdByCate,
	})
}

func (p productController) GetProductBySeller(c echo.Context) error {
	var products []entity.Product

	id := c.Param("id")

	getProdBySeller, err := p.productService.GetProductBySeller(products, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, domain.MessageResp{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, domain.RespData{
		Data: getProdBySeller,
	})
}