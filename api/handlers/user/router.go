package handler_user

import (

	"github.com/louistwiice/go/fripclose/entity"
	"github.com/louistwiice/go/fripclose/repository/user"
	"github.com/louistwiice/go/fripclose/usecase/authentication"
	"github.com/louistwiice/go/fripclose/usecase/user"
)

func NewUserRouters(server *entity.Routers) {
	userRepo := repository_user.NewUserClient(server.Database, server.Redis)

	userService := service_user.NewUserService(userRepo)
	authService := service_authentication.NewAuthService(userRepo)

	userController := NewUserController(userService)
	authController := NewAuthController(authService)

	// Basic connection system
	api_connection := server.OpenApp.Group("auth/")
	api_connection.POST("login", authController.login)
	api_connection.POST("register", authController.register)
	api_connection.POST("activation", authController.activation)
	api_connection.POST("refresh", authController.refreshToken)
	api_connection.GET("logout", authController.logout)

	// Access allowed for everybody
	api_open := server.OpenApp.Group("user/")
	api_open.GET("", userController.listUsers)
	api_open.GET(":id", userController.getUser)

	// Access allowed for connected users
	api_auth := server.RestrictedApp.Group("user/")
	api_auth.GET("", userController.connectedUser)
	api_auth.PUT(":id", userController.updateUser)
	api_auth.POST(":id/reset_password", userController.updatePassword)
	api_auth.DELETE("delete_account/:id", userController.deleteUser)

	api_auth.POST("image", userController.uploadProfileImage)

}