package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/louistwiice/go/fripclose/entity"
)

type CategoryRepository interface {
	List() ([]*entity.Category, error)
	Create(data *entity.Category) error
	GetByID(id int) (*entity.Category, error)
	UpdateTitle(data *entity.Category) error
	UpdateParent(data *entity.Category) error
	ClearParent(data *entity.Category) error
	Delete(data *entity.Category) error
}

type CategoryService interface {
	List() ([]*entity.Category, error)
	Create(data *entity.Category) error
	GetByID(id int) (*entity.Category, error)
	UpdateTitle(data *entity.Category) error
	UpdateParent(data *entity.Category) error
	ClearParent(data *entity.Category) error
	Delete(data *entity.Category) error
}

type CategoryController interface {
	listCategory(ctx *gin.Context)
	createCategory(ctx *gin.Context)
	getGategory(ctx *gin.Context)
	updateTitleCategory(ctx *gin.Context)
	updateParentCategory(ctx *gin.Context)
}