package http

import (
	"errors"
	"math"
	"net/http"
	"strconv"
	"strings"
	"github.com/NetSinx/yconnect-shop/server/category/internal/model"
	"github.com/NetSinx/yconnect-shop/server/category/internal/usecase"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CategoryController struct {
	CategoryUseCase *usecase.CategoryUseCase
	Log             *logrus.Logger
}

func NewCategoryController(categoryUseCase *usecase.CategoryUseCase, log *logrus.Logger) *CategoryController {
	return &CategoryController{
		CategoryUseCase: categoryUseCase,
		Log: log,
	}
}

func (c *CategoryController) ListCategory(ctx echo.Context) error {
	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	pageSize, _ := strconv.Atoi(ctx.QueryParam("page_size"))

	if page <= 0 || pageSize <= 0 {
		page = 1
		pageSize = 20
	}

	categoryRequest := &model.ListCategoryRequest{
		Page: page,
		Size: pageSize,
	}

	listCategories, total, err := c.CategoryUseCase.ListCategory(ctx.Request().Context(), categoryRequest)
	if err != nil {
		c.Log.WithError(err).Error("error listing categories")
		return err
	}

	listCategoriesResponse := model.ListCategoryResponse{
		Data: listCategories,
		Page: categoryRequest.Page,
		Size: categoryRequest.Size,
		TotalItem: total,
		TotalPage: int64(math.Ceil(float64(total) / float64(categoryRequest.Size))),
	}

	return ctx.JSON(http.StatusOK, listCategoriesResponse)
}

func (c *CategoryController) CreateCategory(ctx echo.Context) error {
	var categoryRequest *model.CreateCategoryRequest

	if err := ctx.Bind(categoryRequest); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return err
	}

	if err := c.CategoryUseCase.CreateCategory(ctx.Request().Context(), categoryRequest); err != nil {
		c.Log.WithError(err).Error("error creating category")
		return err
	}

	return ctx.JSON(http.StatusOK, model.CategoryResponseMessage{
		Message: "Data was created successfully",
	})
}

func (c *CategoryController) UpdateCategory(ctx echo.Context) error {
	var categoryRequest *model.UpdateCategoryRequest

	slug := ctx.Param("slug")

	categoryRequest.Slug = slug

	if err := ctx.Bind(categoryRequest); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return err
	}

	if err := c.CategoryUseCase.UpdateCategory(ctx.Request().Context(), categoryRequest); err != nil {
		c.Log.WithError(err).Error("error updating category")
		return err
	}

	return ctx.JSON(http.StatusOK, model.CategoryResponseMessage{
		Message: "",
	})
}

func (c *CategoryController) DeleteCategory(ctx echo.Context) error {
	var category model.Category

	slug := c.Param("slug")

	err := c.categoryService.DeleteCategory(category, slug)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, dto.MessageResp{
			Message: errs.ErrNotFound,
		})
	}

	return c.JSON(http.StatusOK, dto.MessageResp{
		Message: dto.DeleteResponse,
	})
}

func (c *CategoryController) GetCategoryById(ctx echo.Context) error {
	var category model.Category

	id := c.Param("id")

	getCategory, err := c.categoryService.GetCategoryById(category, id)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, dto.MessageResp{
			Message: errs.ErrNotFound,
		})
	}

	return c.JSON(http.StatusOK, dto.RespData{
		Data: getCategory,
	})
}

func (c *CategoryController) GetCategoryBySlug(ctx echo.Context) error {
	var category model.Category

	slug := c.Param("slug")

	getCategory, err := c.categoryService.GetCategoryBySlug(category, slug)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, dto.MessageResp{
			Message: errs.ErrNotFound,
		})
	}

	return c.JSON(http.StatusOK, dto.RespData{
		Data: getCategory,
	})
}
