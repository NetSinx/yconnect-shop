package routes

import (
	"github.com/NetSinx/yconnect-shop/server/category/config"
	"github.com/NetSinx/yconnect-shop/server/category/controller"
	"github.com/NetSinx/yconnect-shop/server/category/repository"
	"github.com/NetSinx/yconnect-shop/server/category/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	authMiddleware "github.com/NetSinx/yconnect-shop/server/category/middleware"
)

func ApiRoutes() *echo.Echo {
	categoryRepository := repository.CategoryRepository(config.DB)
	categoryService := service.CategoryService(categoryRepository)
	categoryController := controller.CategoryController(categoryService)

	router := echo.New()
	router.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		KeyLookup: "header:api-token",
		Validator: func(auth string, c echo.Context) (bool, error) {
			return auth == "dfkgjdgj#753846873248358645*&#%^*$54%hgdf", nil
		},
	}))
	router.GET("/category", categoryController.ListCategory)
	router.GET("/category/:id", categoryController.GetCategory)
	authRoute := router.Group("/auth")
	authRoute.Use(authMiddleware.JWTAuthMiddleware)
	authRoute.POST("/category", categoryController.CreateCategory)
	authRoute.PUT("/category/:id", categoryController.UpdateCategory)
	authRoute.DELETE("/category/:id", categoryController.DeleteCategory)

	return router
}