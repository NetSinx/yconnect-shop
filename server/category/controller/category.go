package controller

import (
	"net/http"
	"strconv"
	"github.com/NetSinx/yconnect-shop/server/category/model/entity"
	"github.com/NetSinx/yconnect-shop/server/category/model/domain"
	"github.com/NetSinx/yconnect-shop/server/category/service"
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
	var categories []entity.Kategori

	listCategories, err := cc.categoryService.ListCategory(categories)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, domain.MessageResp{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, domain.RespData{
		Data: listCategories,
	})
}

func (cc categoryController) CreateCategory(c echo.Context) error {
	var categories entity.Kategori

	if err := c.Bind(&categories); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
			Message: err.Error(),
		})
	}

	category, err := cc.categoryService.CreateCategory(categories)
	if err != nil && err == gorm.ErrDuplicatedKey {
		return echo.NewHTTPError(http.StatusConflict, domain.MessageResp{
			Message: "Kategori sudah tersedia.",
		})
	} else if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, domain.RespData{
		Data: category,
	})
}

func (cc categoryController) UpdateCategory(c echo.Context) error {
	var categories entity.Kategori

	id := c.Param("id")

	if err := c.Bind(&categories); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
			Message: err.Error(),
		})
	}
	
	category, err := cc.categoryService.UpdateCategory(categories, id)
	if err != nil && err == gorm.ErrRecordNotFound {
		return echo.NewHTTPError(http.StatusNotFound, domain.MessageResp{
			Message: "Kategori tidak bisa ditemukan.",
		})
	} else if err != nil && err == gorm.ErrDuplicatedKey {
		return echo.NewHTTPError(http.StatusConflict, domain.MessageResp{
			Message: "Kategori sudah tersedia.",
		})
	} else if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
			Message: err.Error(),
		})
	}

	getId, _ := strconv.ParseUint(id, 32, 10)

	category.Id = uint(getId)
	
	return c.JSON(http.StatusOK, domain.RespData{
		Data: category,
	})
}

func (cc categoryController) DeleteCategory(c echo.Context) error {
	var category entity.Kategori

	id := c.Param("id")

	err := cc.categoryService.DeleteCategory(category, id)
	if err != nil && err == gorm.ErrRecordNotFound {
		return echo.NewHTTPError(http.StatusNotFound, domain.MessageResp{
			Message: "Kategori tidak bisa ditemukan.",
		})
	}

	return c.JSON(http.StatusOK, domain.MessageResp{
		Message: "Kategori berhasil dihapus.",
	})
}

func (cc categoryController) GetCategory(c echo.Context) error {
	var categories entity.Kategori

	id := c.Param("id")

	getCategory, err := cc.categoryService.GetCategory(categories, id)
	if err != nil && err == gorm.ErrRecordNotFound {
		return echo.NewHTTPError(http.StatusNotFound, domain.MessageResp{
			Message: "Kategori tidak bisa ditemukan.",
		})
	}

	return c.JSON(http.StatusOK, domain.RespData{
		Data: getCategory,
	})
}