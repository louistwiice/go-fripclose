package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/louistwiice/go/fripclose/entity"
)

type UserRepository interface {
	List() ([]*entity.UserDisplay, error)
	Create(u *entity.UserCreateUpdate) error
	GetByID(id string) (*entity.UserDisplay, string, error)
	SearchUser(identifier string) (*entity.UserDisplay, string, error)
	UpdateInfo(u *entity.UserCreateUpdate) error
	UploadPicture(u *entity.UserDisplay) error
	UpdatePassword(u *entity.UserCreateUpdate) error
	UpdateAuthenticationDate(u *entity.UserDisplay) error
	Delete(id string) error
	ActivateUser(username string) error

	SaveTokenInRedis(key, value string) (string, error)
	GetTokenInRedis(key string) (string, error)
}

type UserService interface {
	List() ([]*entity.UserDisplay, error)
	Create(u *entity.UserCreateUpdate) error
	GetByID(id string) (*entity.UserDisplay, string, error)
	SearchUser(identifier string) (*entity.UserDisplay, string, error)
	UpdateUser(u *entity.UserCreateUpdate) error
	UploadPicture(u *entity.UserDisplay) error
	UpdatePassword(u *entity.UserCreateUpdate) error
	Delete(id string) error

	IsAdminOrHasRight(id string) (bool, error)
}

type UserController interface {
	listUser(ctx *gin.Context)
	getUser(ctx *gin.Context)
	connectedUser(ctx *gin.Context)
	updateUser(ctx *gin.Context)
	updatePassword(ctx *gin.Context)
	deleteUser(ctx *gin.Context)

	uploadProfileImage(ctx *gin.Context)
}
