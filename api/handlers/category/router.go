package handler_category

import (
	"github.com/louistwiice/go/fripclose/entity"
	"github.com/louistwiice/go/fripclose/repository/category"
	"github.com/louistwiice/go/fripclose/repository/user"
	"github.com/louistwiice/go/fripclose/usecase/category"
	"github.com/louistwiice/go/fripclose/usecase/user"
)

func NewCategoryRouters(server *entity.Routers) {
	categeryRepo := repository_category.NewCategoryClient(server.Database)
	userRepo := repository_user.NewUserClient(server.Database, server.Redis)
	userRightService := service_user.NewUserService(userRepo)

	categoryService := service_category.NewCategoryService(categeryRepo)
	categoryController := NewCategoryController(categoryService, userRightService)

	// API endpoints that don't need authentication
	api_open := server.OpenApp.Group("category")
	api_open.GET("", categoryController.listCategory)
	api_open.GET("/:id", categoryController.getGategory)

	// APIs that need users to be authenticated
	api_restricted := server.RestrictedApp.Group("category")
	api_restricted.POST("", categoryController.createCategory)
	api_restricted.PUT("/:id/title", categoryController.updateTitleCategory)
	api_restricted.PUT("/:id/parent", categoryController.updateParentCategory)
	api_restricted.DELETE("/:id", categoryController.deleteCategory)
}