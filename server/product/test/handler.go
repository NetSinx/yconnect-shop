package test

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	"github.com/NetSinx/yconnect-shop/server/product/handler/dto"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type entityProduct struct {
	Id          uint               `json:"id" gorm:"primaryKey"`
	Nama        string             `json:"nama" form:"nama" gorm:"unique" validate:"required,max=255"`
	Slug        string             `json:"slug" form:"slug" gorm:"unique" validate:"required"`
	Gambar      []entityGambar     `json:"gambar" form:"gambar" validate:"required"`
	Deskripsi   string             `json:"deskripsi" form:"deskripsi" validate:"required"`
	KategoriId  uint               `json:"kategori_id" form:"kategori_id" validate:"required"`
	Harga       int                `json:"harga" form:"harga" validate:"required"`
	Stok        int                `json:"stok" form:"stok" validate:"required"`
	Rating      float32            `json:"rating" form:"rating" validate:"required"`
	Kategori    entityKategori     `json:"kategori" gorm:"-"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
}

type entityKategori struct {
	Id        uint        `json:"id"`
	Name      string      `json:"name"`
	Slug      string      `json:"slug"`
}

type entityGambar struct {
	Id         uint   `json:"id"`
	Nama       string `json:"nama"`
	ProductID  uint   `json:"product_id"`
}

var productModel = append([]entityProduct{}, 
	entityProduct{
		Id: 1,
		Nama: "Baju Muslim",
		Slug: "baju-muslim",
		Gambar: append([]entityGambar{}, 
			entityGambar{
				Nama: "../assets/images/6c6170746f7031d41d8cd98f00b204e9800998ecf8427e.jpg",
				ProductID: 1,
			},
			entityGambar{
				Nama: "../assets/images/6c6170746f7032d41d8cd98f00b204e9800998ecf8427e.jpg",
				ProductID: 1,
			},
		),
		Deskripsi: "Baju muslim yang nyaman digunakan untuk beribadah",
		KategoriId: 1,
		Harga: 95000,
		Stok: 15,
		Rating: 4.9,
	},
	entityProduct{
		Id: 2,
		Nama: "Baju Muslim Koko",
		Slug: "baju-muslim-koko",
		Gambar: append([]entityGambar{}, 
			entityGambar{
				Nama: "../assets/images/6c6170746f7031d41d8cd98f00b204e9800998ecf8427e.jpg",
				ProductID: 1,
			},
			entityGambar{
				Nama: "../assets/images/6c6170746f7032d41d8cd98f00b204e9800998ecf8427e.jpg",
				ProductID: 1,
			},
		),
		Deskripsi: "Baju muslim yang nyaman digunakan untuk beribadah",
		KategoriId: 1,
		Harga: 105000,
		Stok: 25,
		Rating: 5.0,
	},
)

func ListProduct(c echo.Context) error {
	return c.JSON(http.StatusOK, productModel)
}

func CreateProduct(c echo.Context) error {
	var reqProduct entityProduct
	var reqImg []entityGambar

	images, err := c.MultipartForm()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, dto.MessageResp{
			Message: err.Error(),
		})
	}

	if err := c.Bind(&reqProduct); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, dto.MessageResp{
			Message: err.Error(),
		})
	}

	if err := os.MkdirAll("../assets/images", 0750); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, dto.MessageResp{
			Message: err.Error(),
		})
	}

	for _, image := range images.File["gambar"] {
		src, err := image.Open()
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, dto.MessageResp{
				Message: err.Error(),
			})
		}
		defer src.Close()

		fileName := strings.Split(image.Filename, ".")[0]
		fileExt := strings.Split(image.Filename, ".")[1]
		hashedFileName := md5.New().Sum([]byte(fileName))

		dst, err := os.Create(fmt.Sprintf("../assets/images/%x.%s", hashedFileName, fileExt))
		if err != nil {
			return err
		}

		defer dst.Close()

		if _, err := io.Copy(dst, src); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, dto.MessageResp{
				Message: err.Error(),
			})
		}

		reqImg = append(reqImg, entityGambar{Nama: fmt.Sprintf("../assets/images/%x.%s", hashedFileName, fileExt)})
	}

	reqProduct.Gambar = reqImg

	if err := validator.New().Struct(reqProduct); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, dto.MessageResp{
			Message: err.Error(),
		})
	}

	productModel = append(productModel, reqProduct)

	return c.JSON(http.StatusOK, dto.MessageResp{
		Message: "Produk berhasil ditambahkan",
	})
}

func UpdateProduct(c echo.Context) error {
	var reqProduct entityProduct

	slug := c.Param("slug")

	if err := c.Bind(&reqProduct); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, dto.MessageResp{
			Message: err.Error(),
		})
	}

	imageProduct, err := c.MultipartForm()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, dto.MessageResp{
			Message: err.Error(),
		})
	}

	images := imageProduct.File["gambar"]

	for i, prod := range productModel {
		if prod.Slug == slug {
			for _, p := range productModel {
				if (reqProduct.Nama == p.Nama && reqProduct.Slug == p.Slug) && p.Slug != slug {
					return echo.NewHTTPError(http.StatusConflict, dto.MessageResp{
						Message: "Produk sudah terdaftar",
					})
				}
			}

			for _, gambar := range prod.Gambar {
				os.Remove(gambar.Nama)
			}

			for _, image := range images {
				src, err := image.Open()
				if err != nil {
					return echo.NewHTTPError(http.StatusInternalServerError, dto.MessageResp{
						Message: err.Error(),
					})
				}
				defer src.Close()
		
				hashedFileName := md5.New().Sum([]byte(strings.Split(image.Filename, ".")[0]))
				fileExt := strings.Split(image.Filename, ".")[1]
		
				dst, err := os.Create(fmt.Sprintf("../assets/images/%x.%v", hashedFileName, fileExt))
				if err != nil {
					return echo.NewHTTPError(http.StatusInternalServerError, dto.MessageResp{
						Message: err.Error(),
					})
				}
				defer dst.Close()

				if _, err := io.Copy(dst, src); err != nil {
					return echo.NewHTTPError(http.StatusInternalServerError, dto.MessageResp{
						Message: err.Error(),
					})
				}
				
				reqProduct.Gambar = append(reqProduct.Gambar, entityGambar{Nama: fmt.Sprintf("../assets/images/%x.%v", hashedFileName, fileExt), ProductID: prod.Id})
			}
			
			productModel[i] = reqProduct

			return c.JSON(http.StatusOK, dto.MessageResp{
				Message: "Produk berhasil diubah",
			})
		}
	}
	
	return echo.NewHTTPError(http.StatusNotFound, dto.MessageResp{
		Message: "Produk tidak ditemukan",
	})
}

func DeleteProduct(c echo.Context) error {
	slug := c.Param("slug")

	for _, product := range productModel {
		if product.Slug == slug {
			for _, img := range product.Gambar {
				os.Remove(img.Nama)
			}

			product = entityProduct{}

			return c.JSON(http.StatusOK, dto.MessageResp{
				Message: "Produk berhasil dihapus",
			})
		}
	}

	return echo.NewHTTPError(http.StatusNotFound, dto.MessageResp{
		Message: "Produk tidak ditemukan",
	})
}

func GetProduct(c echo.Context) error {
	slug := c.Param("slug")

	for _, product := range productModel {
		if product.Slug == slug {
			return c.JSON(http.StatusOK, product)
		}
	}

	return echo.NewHTTPError(http.StatusNotFound, dto.MessageResp{
		Message: "Produk tidak ditemukan",
	})
}

func GetProductByCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var products []entityProduct

	for i, product := range productModel {
		if product.KategoriId == uint(id) {
			products = append(products, product)

			if i == len(productModel) - 1 {
				return c.JSON(http.StatusOK, products)
			}
		}
	}

	return echo.NewHTTPError(http.StatusNotFound, dto.MessageResp{
		Message: "Produk tidak ditemukan",
	})
}