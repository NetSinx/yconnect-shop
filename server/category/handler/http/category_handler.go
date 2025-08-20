package http

import (
	"errors"
	"net/http"
	"github.com/NetSinx/yconnect-shop/server/category/errs"
	"github.com/NetSinx/yconnect-shop/server/category/handler/dto"
	"github.com/NetSinx/yconnect-shop/server/category/model"
	"github.com/NetSinx/yconnect-shop/server/category/service"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CategoryHandl interface {
	ListCategory(c echo.Context) error
	CreateCategory(c echo.Context) error
	UpdateCategory(c echo.Context) error
	DeleteCategory(c echo.Context) error
	GetCategoryById(c echo.Context) error
	GetCategoryBySlug(c echo.Context) error
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
	var categoryReq dto.CategoryRequest

	if err := c.Bind(&categoryReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, dto.MessageResp{
			Message: err.Error(),
		})
	}

	err := cc.categoryService.CreateCategory(categoryReq)
	if err != nil && errors.Is(err, gorm.ErrDuplicatedKey) {
		return echo.NewHTTPError(http.StatusConflict, dto.MessageResp{
			Message: errs.ErrDuplicatedKey,
		})
	} else if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, dto.MessageResp{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.MessageResp{
		Message: dto.CreateResponse,
	})
}

func (cc categoryHandler) UpdateCategory(c echo.Context) error {
	var categoryReq dto.CategoryRequest

	slug := c.Param("slug")

	if err := c.Bind(&categoryReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, dto.MessageResp{
			Message: err.Error(),
		})
	}
	
	err := cc.categoryService.UpdateCategory(categoryReq, slug)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, dto.MessageResp{
			Message: errs.ErrNotFound,
		})
	} else if err != nil && errors.Is(err, gorm.ErrDuplicatedKey) {
		return echo.NewHTTPError(http.StatusConflict, dto.MessageResp{
			Message: errs.ErrDuplicatedKey,
		})
	} else if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, dto.MessageResp{
			Message: err.Error(),
		})
	}
	
	return c.JSON(http.StatusOK, dto.MessageResp{
		Message: dto.UpdateResponse,
	})
}

func (cc categoryHandler) DeleteCategory(c echo.Context) error {
	var category model.Category

	slug := c.Param("slug")

	err := cc.categoryService.DeleteCategory(category, slug)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, dto.MessageResp{
			Message: errs.ErrNotFound,
		})
	}

	return c.JSON(http.StatusOK, dto.MessageResp{
		Message: dto.DeleteResponse,
	})
}

func (cc categoryHandler) GetCategoryById(c echo.Context) error {
	var category model.Category

	id := c.Param("id")

	getCategory, err := cc.categoryService.GetCategoryById(category, id)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, dto.MessageResp{
			Message: errs.ErrNotFound,
		})
	}

	return c.JSON(http.StatusOK, dto.RespData{
		Data: getCategory,
	})
}

func (cc categoryHandler) GetCategoryBySlug(c echo.Context) error {
	var category model.Category

	slug := c.Param("slug")

	getCategory, err := cc.categoryService.GetCategoryBySlug(category, slug)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, dto.MessageResp{
			Message: errs.ErrNotFound,
		})
	}

	return c.JSON(http.StatusOK, dto.RespData{
		Data: getCategory,
	})
}