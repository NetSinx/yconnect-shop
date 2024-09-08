package test

import (
	"net/http"

	"github.com/NetSinx/yconnect-shop/server/product/model/domain"
	"github.com/NetSinx/yconnect-shop/server/product/model/entity"
	// "github.com/NetSinx/yconnect-shop/server/category/test"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var ProductModel = append([]entity.Product{}, 
	entity.Product{
		Id: 1,
		Nama: "Baju Muslim",
		Slug: "baju-muslim",
		Gambar: append([]entity.Gambar{}, 
			entity.Gambar{
				Nama: "baju_muslim1.jpg",
				ProductID: 1,
			},
			entity.Gambar{
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

func ListProduct(c echo.Context) error {
	// for i, category := range test. {
	// 	if category.Id == ProductModel
	// }

	return c.JSON(http.StatusOK, domain.RespData{
		Data: ProductModel,
	})
}

func CreateProduct(c echo.Context) error {
	var reqProduct entity.Product

	if err := c.Bind(&reqProduct); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
			Message: err.Error(),
		})
	}

	if err := validator.New().Struct(reqProduct); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
			Message: err.Error(),
		})
	}

	

	ProductModel = append(ProductModel, reqProduct)

	return c.JSON(http.StatusOK, domain.MessageResp{
		Message: "Produk berhasil ditambahkan",
	})
}