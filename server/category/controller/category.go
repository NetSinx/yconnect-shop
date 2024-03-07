package controller

import (
	"net/http"
	"strconv"
	"github.com/NetSinx/yconnect-shop/server/category/app/model"
	"github.com/NetSinx/yconnect-shop/server/category/service"
	"github.com/NetSinx/yconnect-shop/server/category/utils"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type categoryController struct {
	categoryService service.CategoryServ
}

func CategoryController(categoryservice service.CategoryServ) categoryController {
	return categoryController{
		categoryService: categoryservice,
	}
}

func (cc categoryController) ListCategory(c echo.Context) error {
	var categories []model.Category

	listCategories, err := cc.categoryService.ListCategory(categories)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, utils.ErrServer{
			Code: http.StatusInternalServerError,
			Status: http.StatusText(http.StatusInternalServerError),
			Message: "Maaf, ada kesalahan pada server",
		})
	}

	return c.JSON(http.StatusOK, utils.SuccessData{
		Code: http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data: listCategories,
	})
}

func (cc categoryController) CreateCategory(c echo.Context) error {
	var categories model.Category

	if err := c.Bind(&categories); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrServer{
			Code: http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Message: "Request yang dikirimkan tidak sesuai!",
		})
	}

	category, err := cc.categoryService.CreateCategory(categories)
	if err != nil && err.Error() == "request tidak sesuai" {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrServer{
			Code: http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Message: "Request yang dikirimkan tidak sesuai!",
		})
	} else if err != nil && err.Error() == "kategori sudah tersedia" {
		return echo.NewHTTPError(http.StatusConflict, utils.ErrServer{
			Code: http.StatusConflict,
			Status: http.StatusText(http.StatusConflict),
			Message: "Kategori sudah tersedia!",
		})
	}

	return c.JSON(http.StatusOK, utils.SuccessData{
		Code: http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data: category,
	})
}

func (cc categoryController) UpdateCategory(c echo.Context) error {
	var categories model.Category

	id := c.Param("id")

	if err := c.Bind(&categories); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrServer{
			Code: http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Message: "Request yang dikirimkan tidak sesuai!",
		})
	}
	
	category, err := cc.categoryService.UpdateCategory(categories, id)
	if err != nil && err == gorm.ErrRecordNotFound {
		return echo.NewHTTPError(http.StatusNotFound, utils.ErrServer{
			Code: http.StatusNotFound,
			Status: http.StatusText(http.StatusNotFound),
			Message: "Kategori tidak bisa ditemukan!",
		})
	} else if err != nil && err.Error() == "request tidak sesuai" {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrServer{
			Code: http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Message: "Request yang dikirimkan tidak sesuai!",
		})
	} else if err != nil && err.Error() == "kategori sudah tersedia" {
		return echo.NewHTTPError(http.StatusConflict, utils.ErrServer{
			Code: http.StatusConflict,
			Status: http.StatusText(http.StatusConflict),
			Message: "Kategori sudah tersedia!",
		})
	}

	getId, _ := strconv.ParseUint(id, 32, 10)

	category.Id = uint(getId)
	
	return c.JSON(http.StatusOK, utils.SuccessData{
		Code: http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data: category,
	})
}

func (cc categoryController) DeleteCategory(c echo.Context) error {
	var category model.Category

	id := c.Param("id")

	err := cc.categoryService.DeleteCategory(category, id)
	if err != nil && err == gorm.ErrRecordNotFound {
		return echo.NewHTTPError(http.StatusNotFound, utils.ErrServer{
			Code: http.StatusNotFound,
			Status: http.StatusText(http.StatusNotFound),
			Message: "Kategori tidak bisa ditemukan!",
		})
	}

	return c.JSON(http.StatusOK, utils.SuccessDelete{
		Code: http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Message: "Kategori berhasil dihapus!",
	})
}

func (cc categoryController) GetCategory(c echo.Context) error {
	var categories model.Category

	id := c.Param("id")

	getCategory, err := cc.categoryService.GetCategory(categories, id); if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, utils.ErrServer{
			Code: http.StatusNotFound,
			Status: http.StatusText(http.StatusNotFound),
			Message: "Kategori tidak bisa ditemukan!",
		})
	}

	return c.JSON(http.StatusOK, utils.SuccessData{
		Code: http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data: getCategory,
	})
}