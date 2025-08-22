package test

import (
	"net/http"
	"github.com/NetSinx/yconnect-shop/server/product/handler/dto"
	"github.com/NetSinx/yconnect-shop/server/product/helpers"
	"github.com/NetSinx/yconnect-shop/server/product/model"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var productModel = append([]model.Product{}, 
	model.Product{
		Id: 1,
		Nama: "Baju Muslim",
		Slug: "baju-muslim",
		Gambar: append([]model.Gambar{}, 
			model.Gambar{
				Path: "../assets/images/6c6170746f7031d41d8cd98f00b204e9800998ecf8427e.jpg",
				ProductID: 1,
			},
			model.Gambar{
				Path: "../assets/images/6c6170746f7032d41d8cd98f00b204e9800998ecf8427e.jpg",
				ProductID: 1,
			},
		),
		Deskripsi: "Baju muslim yang nyaman digunakan untuk beribadah",
		KategoriSlug: "pakaian",
		Harga: 95000,
		Stok: 15,
	},
	model.Product{
		Id: 2,
		Nama: "Baju Muslim Koko",
		Slug: "baju-muslim-koko",
		Gambar: append([]model.Gambar{}, 
			model.Gambar{
				Path: "../assets/images/6c6170746f7031d41d8cd98f00b204e9800998ecf8427e.jpg",
				ProductID: 1,
			},
			model.Gambar{
				Path: "../assets/images/6c6170746f7032d41d8cd98f00b204e9800998ecf8427e.jpg",
				ProductID: 1,
			},
		),
		Deskripsi: "Baju muslim yang nyaman digunakan untuk beribadah",
		KategoriSlug: "pakaian",
		Harga: 105000,
		Stok: 25,
	},
)

func ListProduct(c echo.Context) error {
	return c.JSON(http.StatusOK, productModel)
}

func CreateProduct(c echo.Context) error {
	var productReq dto.ProductRequest

	if err := c.Bind(&productReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, dto.MessageResp{
			Message: err.Error(),
		})
	}

	if err := validator.New().Struct(productReq); err != nil {
		return err
	}

	slug, err := helpers.GenerateSlugByName(productReq.Nama)
	if err != nil {
		return err
	}

	product := model.Product{
		Nama: productReq.Nama,
		Deskripsi: productReq.Deskripsi,
		Slug: slug,
		Gambar: productReq.Gambar,
		KategoriSlug: productReq.KategoriSlug,
		Harga: productReq.Harga,
		Stok: productReq.Stok,
	}

	productModel = append(productModel, product)

	return c.JSON(http.StatusOK, dto.MessageResp{
		Message: "Produk berhasil ditambahkan",
	})
}

func UpdateProduct(c echo.Context) error {
	var productReq dto.ProductRequest

	slug := c.Param("slug")

	if err := c.Bind(&productReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, dto.MessageResp{
			Message: err.Error(),
		})
	}

	if err := validator.New().Struct(productReq); err != nil {
		return err
	}

	slugGenerator, err := helpers.GenerateSlugByName(productReq.Nama)
	if err != nil {
		return err
	}
	
	product := model.Product{
		Nama: productReq.Nama,
		Deskripsi: productReq.Deskripsi,
		Slug: slugGenerator,
		Gambar: productReq.Gambar,
		KategoriSlug: productReq.KategoriSlug,
		Harga: productReq.Harga,
		Stok: productReq.Stok,
	}

	for _, p := range productModel {
		if p.Slug == slug {
			p = product
			
			return c.JSON(http.StatusOK, dto.MessageResp{
				Message: "Produk berhasil diubah",
			})
		}
	}
	
	return echo.NewHTTPError(http.StatusNotFound, dto.MessageResp{
		Message: "Produk tidak ditemukan.",
	})
}

func DeleteProduct(c echo.Context) error {
	slug := c.Param("slug")

	for _, product := range productModel {
		if product.Slug == slug {
			product = model.Product{}

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