package middlewares

import (
	"github.com/louistwiice/go/fripclose/entity"
	"github.com/louistwiice/go/fripclose/repository/user"
	"github.com/louistwiice/go/fripclose/usecase/authentication"
)

func NewMiddlewareRouters(server *entity.Routers) *controller {
	userRepo := repository_user.NewUserClient(server.Database, server.Redis)
	authService := service_authentication.NewAuthService(userRepo)

	return NewMiddlewareControllers(authService)
}