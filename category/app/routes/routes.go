package routes

import (
	"net/http"
	"github.com/NetSinx/yconnect-shop/category/app/config"
	"github.com/NetSinx/yconnect-shop/category/controller"
	"github.com/NetSinx/yconnect-shop/category/repository"
	"github.com/NetSinx/yconnect-shop/category/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ApiRoutes() *echo.Echo {
	categoryRepository := repository.CategoryRepository(config.DB)
	categoryService := service.CategoryService(categoryRepository)
	categoryController := controller.CategoryController(categoryService)

	router := echo.New()
	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"http://localhost:4200"},
			AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		}),
	)
	router.GET("/categories", categoryController.ListCategory)
	router.GET("/categories/:slug", categoryController.GetCategory)
	router.POST("/categories", categoryController.CreateCategory)
	router.PUT("/categories/:slug", categoryController.UpdateCategory)
	router.DELETE("/categories/:slug", categoryController.DeleteCategory)

	return router
}