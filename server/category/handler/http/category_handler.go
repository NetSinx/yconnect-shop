package http

import (
	"net/http"
	"strconv"
	"github.com/NetSinx/yconnect-shop/server/category/model"
	"github.com/NetSinx/yconnect-shop/server/category/handler/dto"
	"github.com/NetSinx/yconnect-shop/server/category/service"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CategoryHandl interface {
	ListCategory(c echo.Context) error
	CreateCategory(c echo.Context) error
	UpdateCategory(c echo.Context) error
	DeleteCategory(c echo.Context) error
	GetCategory(c echo.Context) error
}

type categoryHandler struct {
	categoryService service.CategoryServ
}

func CategoryHandler(categoryservice service.CategoryServ) categoryHandler {
	return categoryHandler{
		categoryService: categoryservice,
	}
}

func (cc categoryHandler) ListCategory(c echo.Context) error {
	var categories []model.Category

	listCategories, err := cc.categoryService.ListCategory(categories)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, dto.MessageResp{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.RespData{
		Data: listCategories,
	})
}

func (cc categoryHandler) CreateCategory(c echo.Context) error {
	var categories model.Category

	if err := c.Bind(&categories); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, dto.MessageResp{
			Message: err.Error(),
		})
	}

	category, err := cc.categoryService.CreateCategory(categories)
	if err != nil && err == gorm.ErrDuplicatedKey {
		return echo.NewHTTPError(http.StatusConflict, dto.MessageResp{
			Message: "Kategori sudah tersedia.",
		})
	} else if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, dto.MessageResp{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.RespData{
		Data: category,
	})
}

func (cc categoryHandler) UpdateCategory(c echo.Context) error {
	var categories model.Category

	id := c.Param("id")

	if err := c.Bind(&categories); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, dto.MessageResp{
			Message: err.Error(),
		})
	}
	
	category, err := cc.categoryService.UpdateCategory(categories, id)
	if err != nil && err == gorm.ErrRecordNotFound {
		return echo.NewHTTPError(http.StatusNotFound, dto.MessageResp{
			Message: "Kategori tidak bisa ditemukan.",
		})
	} else if err != nil && err == gorm.ErrDuplicatedKey {
		return echo.NewHTTPError(http.StatusConflict, dto.MessageResp{
			Message: "Kategori sudah tersedia.",
		})
	} else if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, dto.MessageResp{
			Message: err.Error(),
		})
	}

	getId, _ := strconv.ParseUint(id, 32, 10)

	category.Id = uint(getId)
	
	return c.JSON(http.StatusOK, dto.RespData{
		Data: category,
	})
}

func (cc categoryHandler) DeleteCategory(c echo.Context) error {
	var category model.Category

	id := c.Param("id")

	err := cc.categoryService.DeleteCategory(category, id)
	if err != nil && err == gorm.ErrRecordNotFound {
		return echo.NewHTTPError(http.StatusNotFound, dto.MessageResp{
			Message: "Kategori tidak bisa ditemukan.",
		})
	}

	return c.JSON(http.StatusOK, dto.MessageResp{
		Message: "Kategori berhasil dihapus.",
	})
}

func (cc categoryHandler) GetCategory(c echo.Context) error {
	var categories model.Category

	id := c.Param("id")

	getCategory, err := cc.categoryService.GetCategory(categories, id)
	if err != nil && err == gorm.ErrRecordNotFound {
		return echo.NewHTTPError(http.StatusNotFound, dto.MessageResp{
			Message: "Kategori tidak bisa ditemukan.",
		})
	}

	return c.JSON(http.StatusOK, dto.RespData{
		Data: getCategory,
	})
}