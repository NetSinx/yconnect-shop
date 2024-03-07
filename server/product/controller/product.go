package controller

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"github.com/NetSinx/yconnect-shop/server/product/app/config"
	"github.com/NetSinx/yconnect-shop/server/product/app/model"
	"github.com/NetSinx/yconnect-shop/server/product/service"
	"github.com/NetSinx/yconnect-shop/server/product/utils"
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
		return echo.NewHTTPError(http.StatusInternalServerError, utils.ErrServer{
			Code: http.StatusInternalServerError,
			Status: http.StatusText(http.StatusInternalServerError),
			Message: "Maaf, ada kesalahan pada server!",
		})
	}

	return c.JSON(http.StatusOK, utils.SuccessData{
		Code: http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data: listProducts,
	})
}

func (p productController) CreateProduct(c echo.Context) error {
	var products model.Product
	var img []model.Image

	category_id, _ := strconv.ParseUint(c.FormValue("category_id"), 10, 32)
	seller_id, _ := strconv.ParseUint(c.FormValue("seller_id"), 10, 32)

	products.Name = c.FormValue("name")
	products.Slug = c.FormValue("slug")
	products.Description = c.FormValue("description")
	products.CategoryId = uint(category_id)
	products.SellerId = uint(seller_id)
	products.Price, _ = strconv.Atoi(c.FormValue("price"))
	products.Stock, _ = strconv.Atoi(c.FormValue("stock"))

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
			return err
		}
		defer dst.Close()
	
		if _, err := io.Copy(dst, src); err != nil {
			return err
		}
	
		img = append(img, model.Image{Name: fmt.Sprintf("/assets/images/%x.%s", hashedFileName, fileExt)})
	}

	products.Image = img

	if err := c.Bind(&products); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrServer{
			Code: http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Message: err.Error(),
		})
	}
	
	product, err := p.productService.CreateProduct(products, img)
	if err != nil && err.Error() == "produk sudah tersedia" {
		return echo.NewHTTPError(http.StatusConflict, utils.ErrServer{
			Code: http.StatusConflict,
			Status: http.StatusText(http.StatusConflict),
			Message: "Produk sudah tersedia!",
		})
	} else if err != nil && err.Error() == "request tidak sesuai" {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrServer{
			Code: http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Message: "Request yang dikirimkan tidak sesuai!",
		})
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, utils.ErrServer{
			Code: http.StatusInternalServerError,
			Status: http.StatusText(http.StatusInternalServerError),
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, utils.SuccessData{
		Code: http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data: product,
	})
}

func (p productController) UpdateProduct(c echo.Context) error {
	var products model.Product

	id := c.Param("id")
	getId, _ := strconv.ParseUint(id, 32, 10)

	imageProduct, err := c.MultipartForm()
	if err != nil {
		return err
	}

	images := imageProduct.File["images"]

	getProduct, err := p.productService.GetProduct(products, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, utils.ErrServer{
			Code: http.StatusInternalServerError,
			Status: http.StatusText(http.StatusInternalServerError),
			Message: "Maaf, ada kesalahan pada server",
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

			getProduct.Image = append(getProduct.Image, model.Image{Name: fmt.Sprintf("/assets/images/%x.%s", hashedFileName, fileExt), ProductID: uint(getId)})

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
			getProduct.Image[i].ProductID = uint(getId)
		}
	}

	products.Image = getProduct.Image
	
	if err := c.Bind(&products); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrServer{
			Code: http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Message: "Request yang dikirim tidak sesuai!",
		})
	}

	product, err := p.productService.UpdateProduct(products, products.Image, uint(getId))
	if err != nil && err.Error() == "request tidak sesuai" {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrServer{
			Code: http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Message: "Request yang dikirim tidak sesuai!",
		})
	} else if err != nil && err.Error() == "produk tidak bisa ditemukan" {
		return echo.NewHTTPError(http.StatusNotFound, utils.ErrServer{
			Code: http.StatusNotFound,
			Status: http.StatusText(http.StatusNotFound),
			Message: "Produk tidak ditemukan!",
		})
	} else if err != nil && err.Error() == "produk sudah tersedia" {
		return echo.NewHTTPError(http.StatusConflict, utils.ErrServer{
			Code: http.StatusConflict,
			Status: http.StatusText(http.StatusConflict),
			Message: "Produk sudah tersedia!",
		})
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, utils.ErrServer{
			Code: http.StatusInternalServerError,
			Status: http.StatusText(http.StatusInternalServerError),
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, utils.SuccessData{
		Code: http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data: product,
	})
}

func (p productController) DeleteProduct(c echo.Context) error {
	var products model.Product

	id := c.Param("id")

	getProduct, err := p.productService.GetProduct(products, id)
	if err != nil && err.Error() == "produk tidak ditemukan" {
		return echo.NewHTTPError(http.StatusNotFound, utils.ErrServer{
			Code: http.StatusNotFound,
			Status: http.StatusText(http.StatusNotFound),
			Message: "Produk tidak ditemukan!",
		})
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, utils.ErrServer{
			Code: http.StatusInternalServerError,
			Status: http.StatusText(http.StatusInternalServerError),
			Message: "Maaf, ada kesalahan pada server",
		})
	}

	for _, image := range getProduct.Image {
		os.Remove("." + image.Name)
	}

	if err := p.productService.DeleteProduct(products, getProduct.Image, id); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, utils.ErrServer{
			Code: http.StatusNotFound,
			Status: http.StatusText(http.StatusNotFound),
			Message: "Produk tidak ditemukan!",
		})
	}

	return c.JSON(http.StatusOK, utils.ErrServer{
		Code: http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Message: "Produk berhasil dihapus!",
	})
}

func (p productController) GetProduct(c echo.Context) error {
	var product model.Product

	id := c.Param("id")

	getProduct, err := p.productService.GetProduct(product, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, utils.ErrServer{
			Code: http.StatusNotFound,
			Status: http.StatusText(http.StatusNotFound),
			Message: "Produk tidak ditemukan!",
		})
	}
	
	return c.JSON(http.StatusOK, utils.SuccessData{
		Code: http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data: getProduct,
	})
}

func (p productController) GetProductByCategory(c echo.Context) error {
	var products []model.Product

	id := c.Param("id")

	if err := config.DB.Preload("Image").Find(&products, "category_id = ?", id).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, utils.ErrServer{
			Code: http.StatusInternalServerError,
			Status: http.StatusText(http.StatusInternalServerError),
			Message: "Maaf, ada kesalahan pada server!",
		})
	}

	return c.JSON(http.StatusOK, utils.SuccessData{
		Code: http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data: products,
	})
}

func (p productController) GetProductBySeller(c echo.Context) error {
	var products []model.Product

	id := c.Param("id")

	if err := config.DB.Preload("Image").Find(&products, "seller_id = ?", id).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, utils.ErrServer{
			Code: http.StatusInternalServerError,
			Status: http.StatusText(http.StatusInternalServerError),
			Message: "Maaf, ada kesalahan pada server!",
		})
	}

	return c.JSON(http.StatusOK, utils.SuccessData{
		Code: http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data: products,
	})
}