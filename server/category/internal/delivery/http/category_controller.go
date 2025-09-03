package http

import (
	"math"
	"net/http"
	"strconv"
	"github.com/NetSinx/yconnect-shop/server/category/internal/model"
	"github.com/NetSinx/yconnect-shop/server/category/internal/usecase"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
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
		PageMetadata: &model.PageMetadataResponse{
			Page: categoryRequest.Page,
			Size: categoryRequest.Size,
			TotalItem: total,
			TotalPage: int64(math.Ceil(float64(total) / float64(categoryRequest.Size))),
		},
	}

	return ctx.JSON(http.StatusOK, listCategoriesResponse)
}

func (c *CategoryController) CreateCategory(ctx echo.Context) error {
	categoryRequest := new(model.CreateCategoryRequest)

	if err := ctx.Bind(categoryRequest); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return err
	}

	response, err := c.CategoryUseCase.CreateCategory(ctx.Request().Context(), categoryRequest)
	if err != nil {
		c.Log.WithError(err).Error("error creating category")
		return err
	}

	return ctx.JSON(http.StatusOK, model.CategoryResponseWithMessage{
		Message: "Data created successfully",
		Data: response,
	})
}

func (c *CategoryController) UpdateCategory(ctx echo.Context) error {
	categoryRequest := new(model.UpdateCategoryRequest)
	slug := ctx.Param("slug")

	categoryRequest.Slug = slug
	if err := ctx.Bind(categoryRequest); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return err
	}

	response, err := c.CategoryUseCase.UpdateCategory(ctx.Request().Context(), categoryRequest)
	if err != nil {
		c.Log.WithError(err).Error("error updating category")
		return err
	}

	return ctx.JSON(http.StatusOK, model.CategoryResponseWithMessage{
		Message: "Data updated successfully",
		Data: response,
	})
}

func (c *CategoryController) DeleteCategory(ctx echo.Context) error {
	categoryRequest := new(model.DeleteCategoryRequest)
	slug := ctx.Param("slug")

	categoryRequest.Slug = slug
	response, err := c.CategoryUseCase.DeleteCategory(ctx.Request().Context(), categoryRequest)
	if err != nil {
		c.Log.WithError(err).Error("error deleting category")
		return err
	}

	return ctx.JSON(http.StatusOK, model.CategoryResponseWithMessage{
		Message: "Data deleted successfully",
		Data: response,
	})
}

func (c *CategoryController) GetCategoryBySlug(ctx echo.Context) error {
	categoryRequest := new(model.GetCategoryBySlugRequest)
	slug := ctx.Param("slug")

	categoryRequest.Slug = slug
	response, err := c.CategoryUseCase.GetCategoryBySlug(ctx.Request().Context(), categoryRequest)
	if err != nil {
		c.Log.WithError(err).Error("error getting category")
		return err
	}

	return ctx.JSON(http.StatusOK, model.GetCategoryResponse{
		Data: response,
	})
}
