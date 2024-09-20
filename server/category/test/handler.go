package test

import (
	"net/http"
	"strconv"
	"strings"
	"github.com/NetSinx/yconnect-shop/server/category/model/domain"
	"github.com/NetSinx/yconnect-shop/server/category/model/entity"
	prodEntity "github.com/NetSinx/yconnect-shop/server/product/model/entity"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var categoryModel = map[string][]entity.Kategori{
	"data": append([]entity.Kategori{},
		entity.Kategori{
			Id: 1,
			Name: "Baju",
			Slug: "baju",
			Product: append([]prodEntity.Product{}, prodEntity.Product{
				Id: 1,
				Nama: "Baju Muslim",
				Slug: "baju-muslim",
				Gambar: append([]prodEntity.Gambar{}, 
					prodEntity.Gambar{
						Nama: "baju_muslim1.jpg",
						ProductID: 1,
					},
					prodEntity.Gambar{
						Nama: "baju_muslim2.jpg",
						ProductID: 1,
					},
				),
				Deskripsi: "Baju muslim yang nyaman digunakan untuk beribadah",
				KategoriId: 1,
				Harga: 95000,
				Stok: 15,
				Rating: 4.9,
			}),
		},
		entity.Kategori{
			Id: 2,
			Name: "Celana",
			Slug: "celana",
		},
	),
}

var productModel = append([]prodEntity.Product{}, 
	prodEntity.Product{
		Id: 1,
		Nama: "Baju Muslim",
		Slug: "baju-muslim",
		Gambar: append([]prodEntity.Gambar{}, 
			prodEntity.Gambar{
				Nama: "baju_muslim1.jpg",
				ProductID: 1,
			},
			prodEntity.Gambar{
				Nama: "baju_muslim2.jpg",
				ProductID: 1,
			},
		),
		Deskripsi: "Baju muslim yang nyaman digunakan untuk beribadah",
		KategoriId: 1,
		Harga: 95000,
		Stok: 15,
		Rating: 4.9,
	},
)

func ListCategory(c echo.Context) error {
	for _, category := range categoryModel["data"] {
		for _, product := range productModel {
			if category.Id == product.KategoriId {
				category.Product = append(category.Product, product)
			}
		}
	}

	return c.JSON(http.StatusOK, domain.RespData{
		Data: categoryModel["data"],
	})
}

func CreateCategory(c echo.Context) error {
	var reqCategory entity.Kategori
	
	if err := c.Bind(&reqCategory); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
			Message: err.Error(),
		})
	}

	reqCategory.Name = strings.Title(reqCategory.Name)
	reqCategory.Slug = strings.ToLower(reqCategory.Slug)

	if err := validator.New().Struct(reqCategory); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
			Message: err.Error(),
		})
	}

	for _, category := range categoryModel["data"] {
		if reqCategory.Name == category.Name && reqCategory.Slug == category.Slug {
			return echo.NewHTTPError(http.StatusConflict, domain.MessageResp{
				Message: "Kategori tersebut sudah terdaftar",
			})
		}
	}

	return c.JSON(http.StatusOK, domain.MessageResp{
		Message: "Kategori berhasil ditambahkan",
	})
}

func UpdateCategory(c echo.Context) error {
	var reqCategory entity.Kategori

	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.Bind(&reqCategory); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
			Message: err.Error(),
		})
	}

	reqCategory.Name = strings.Title(reqCategory.Name)
	reqCategory.Slug = strings.ToLower(reqCategory.Slug)

	if err := validator.New().Struct(reqCategory); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
			Message: err.Error(),
		})
	}
	
	for _, category := range categoryModel["data"] {
		if uint(id) == category.Id {
			for _, c := range categoryModel["data"] {
				if (reqCategory.Name == c.Name && reqCategory.Slug == c.Slug) && category.Id != c.Id {
					return echo.NewHTTPError(http.StatusConflict, domain.MessageResp{
						Message: "Kategori tersebut sudah terdaftar",
					})
				}
			}

			category.Name = reqCategory.Name
			category.Slug = reqCategory.Slug
			
			return c.JSON(http.StatusOK, domain.MessageResp{
				Message: "Kategori berhasil diubah",
			})
		}
	}

	return echo.NewHTTPError(http.StatusNotFound, domain.MessageResp{
		Message: "Kategori tidak ditemukan",
	})
}

func DeleteCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	for _, category := range categoryModel["data"] {
		if uint(id) == category.Id {
			category = entity.Kategori{}
			
			return c.JSON(http.StatusOK, domain.MessageResp{
				Message: "Kategori berhasil dihapus",
			})
		}
	}

	return echo.NewHTTPError(http.StatusNotFound, domain.MessageResp{
		Message: "Kategori tidak ditemukan",
	})
}

func GetCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	for _, category := range categoryModel["data"] {
		if uint(id) == category.Id {
			return c.JSON(http.StatusOK, category)
		}
	}

	return echo.NewHTTPError(http.StatusNotFound, domain.MessageResp{
		Message: "Kategori tidak ditemukan",
	})
}