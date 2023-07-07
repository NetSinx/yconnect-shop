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
			AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		}),
	)
	router.GET("/categories", categoryController.ListCategory)
	router.GET("/categories/:slug", categoryController.GetCategory)

	routerAuth := router.Group("/api", authMiddleware.AuthMiddleware)
	routerAuth.POST("/categories", categoryController.CreateCategory)
	routerAuth.PUT("/categories/:slug", categoryController.UpdateCategory)
	routerAuth.DELETE("/categories/:slug", categoryController.DeleteCategory)

	return router
}