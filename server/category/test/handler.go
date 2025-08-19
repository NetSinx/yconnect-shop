package test

import (
	"net/http"
	"strconv"
	"strings"
	"github.com/NetSinx/yconnect-shop/server/category/handler/dto"
	"github.com/NetSinx/yconnect-shop/server/category/model"
	prodEntity "github.com/NetSinx/yconnect-shop/server/product/model"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var categoryModel = append([]model.Category{},
	model.Category{
		Id: 1,
		Name: "Baju",
		Slug: "baju",
	},
	model.Category{
		Id: 2,
		Name: "Celana",
		Slug: "celana",
	},
)

var productModel = append([]prodEntity.Product{}, 
	prodEntity.Product{
		Id: 1,
		Nama: "Baju Muslim",
		Slug: "baju-muslim",
		Gambar: append([]prodEntity.Gambar{}, 
			prodEntity.Gambar{
				Path: "baju_muslim1.jpg",
				ProductID: 1,
			},
			prodEntity.Gambar{
				Path: "baju_muslim2.jpg",
				ProductID: 1,
			},
		),
		Deskripsi: "Baju muslim yang nyaman digunakan untuk beribadah",
		KategoriID: 1,
		Harga: 95000,
		Stok: 15,
	},
)

func ListCategory(c echo.Context) error {
	return c.JSON(http.StatusOK, dto.RespData{
		Data: categoryModel,
	})
}

func CreateCategory(c echo.Context) error {
	var categoryReq model.Category
	
	if err := c.Bind(&categoryReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, dto.MessageResp{
			Message: err.Error(),
		})
	}

	categoryReq.Name = strings.ToTitle(categoryReq.Name)
	categoryReq.Slug = strings.ToLower(categoryReq.Slug)

	if err := validator.New().Struct(categoryReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, dto.MessageResp{
			Message: err.Error(),
		})
	}

	for _, category := range categoryModel {
		if categoryReq.Name == category.Name || categoryReq.Slug == category.Slug {
			return echo.NewHTTPError(http.StatusConflict, dto.MessageResp{
				Message: "Kategori tersebut sudah terdaftar",
			})
		}
	}

	return c.JSON(http.StatusOK, dto.MessageResp{
		Message: "Kategori berhasil ditambahkan",
	})
}

func UpdateCategory(c echo.Context) error {
	var reqCategory model.Category

	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.Bind(&reqCategory); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, dto.MessageResp{
			Message: err.Error(),
		})
	}

	reqCategory.Name = strings.Title(reqCategory.Name)
	reqCategory.Slug = strings.ToLower(reqCategory.Slug)

	if err := validator.New().Struct(reqCategory); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, dto.MessageResp{
			Message: err.Error(),
		})
	}
	
	for _, category := range categoryModel {
		if uint(id) == category.Id {
			for _, c := range categoryModel {
				if (reqCategory.Name == c.Name && reqCategory.Slug == c.Slug) && category.Id != c.Id {
					return echo.NewHTTPError(http.StatusConflict, dto.MessageResp{
						Message: "Kategori tersebut sudah terdaftar",
					})
				}
			}

			category.Name = reqCategory.Name
			category.Slug = reqCategory.Slug
			
			return c.JSON(http.StatusOK, dto.MessageResp{
				Message: "Kategori berhasil diubah",
			})
		}
	}

	return echo.NewHTTPError(http.StatusNotFound, dto.MessageResp{
		Message: "Kategori tidak ditemukan",
	})
}

func DeleteCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	for _, category := range categoryModel {
		if uint(id) == category.Id {
			category = model.Category{}
			
			return c.JSON(http.StatusOK, dto.MessageResp{
				Message: "Kategori berhasil dihapus",
			})
		}
	}

	return echo.NewHTTPError(http.StatusNotFound, dto.MessageResp{
		Message: "Kategori tidak ditemukan",
	})
}

func GetCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	for _, category := range categoryModel {
		if uint(id) == category.Id {
			return c.JSON(http.StatusOK, category)
		}
	}

	return echo.NewHTTPError(http.StatusNotFound, dto.MessageResp{
		Message: "Kategori tidak ditemukan",
	})
}