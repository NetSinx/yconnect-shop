package routes

import (
	"net/http"
	"github.com/NetSinx/yconnect-shop/category/app/config"
	authMiddleware "github.com/NetSinx/yconnect-shop/category/app/middleware"
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
			AllowMethods: []string{http.MethodGet},
		}),
	)
	router.GET("/categories", categoryController.ListCategory)
	router.GET("/categories/:slug", categoryController.GetCategory)

	router.Group("/api", middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"http://localhost:4200"},
			AllowMethods: []string{http.MethodPost, http.MethodPut, http.MethodDelete},
		}),
		authMiddleware.AuthMiddleware,
	)
	router.POST("/api/categories", categoryController.CreateCategory)
	router.PUT("/api/categories/:slug", categoryController.UpdateCategory)
	router.DELETE("/api/categories/:slug", categoryController.DeleteCategory)

	return router
}