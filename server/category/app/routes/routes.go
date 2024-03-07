package routes

import (
	"github.com/NetSinx/yconnect-shop/server/category/app/config"
	"github.com/NetSinx/yconnect-shop/server/category/controller"
	"github.com/NetSinx/yconnect-shop/server/category/repository"
	"github.com/NetSinx/yconnect-shop/server/category/service"
	"github.com/labstack/echo/v4"
)

func ApiRoutes() *echo.Echo {
	categoryRepository := repository.CategoryRepository(config.DB)
	categoryService := service.CategoryService(categoryRepository)
	categoryController := controller.CategoryController(categoryService)

	router := echo.New()
	router.GET("/category", categoryController.ListCategory)
	router.GET("/category/:id", categoryController.GetCategory)
	router.POST("/category", categoryController.CreateCategory)
	router.PUT("/category/:id", categoryController.UpdateCategory)
	router.DELETE("/category/:id", categoryController.DeleteCategory)

	return router
}