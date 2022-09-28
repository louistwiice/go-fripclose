package category

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/louistwiice/go/fripclose/domain"
	"github.com/louistwiice/go/fripclose/entity"
	"github.com/louistwiice/go/fripclose/utils"
	jwttoken "github.com/louistwiice/go/fripclose/utils/jwt_token"
)

type controller struct {
	svc_category domain.CategoryService
	svc_user domain.UserRightService
}

func NewCategoryController(svc_cat domain.CategoryService, svc_user domain.UserRightService) *controller {
	return &controller{
		svc_category: svc_cat,
		svc_user: svc_user,
	}
}

func (c *controller) MakeCategoryHandlersWithAuth(app *gin.RouterGroup) {
	app.POST("", c.createCategory)
	app.PUT("/:id/title", c.updateTitleCategory)
	app.PUT("/:id/parent", c.updateParentCategory)
	app.DELETE("/:id", c.deleteCategory)
}

func (c *controller) MakeCategoryHandlersWithoutAuth(app *gin.RouterGroup) {
	app.GET("", c.listCategory)
	app.GET("/:id", c.getGategory)
}


/*
**
**
**
*/

func (c *controller) listCategory(ctx *gin.Context) {
	 categories, err := c.svc_category.List()
	 if err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}
	utils.ResponseJSON(ctx, http.StatusOK, http.StatusOK, "successful", categories)
}

func (c *controller) createCategory(ctx *gin.Context) {
	var data *entity.Category
	
	claim, err := jwttoken.ExtractClaimsFromAccess(ctx)
	if err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}

	is_admin, err := c.svc_user.IsAdminOrHasRight(claim["sub"].(string))
	if err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if !is_admin {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, entity.ErrUnauthorizedAction.Error(), nil)
		return
	}

	if err := ctx.ShouldBindJSON(&data); err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}

	err = c.svc_category.Create(data)
	if err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}
	utils.ResponseJSON(ctx, http.StatusOK, http.StatusCreated, "Category successfully created",data)

}

func (c *controller) getGategory(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}

	cat, err := c.svc_category.GetByID(id)
	if err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}
	utils.ResponseJSON(ctx, http.StatusOK, http.StatusFound, "successful", cat)
}

func (c *controller) updateTitleCategory(ctx *gin.Context) {
	var data *entity.Category

	claim, err := jwttoken.ExtractClaimsFromAccess(ctx)
	if err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}

	is_admin, err := c.svc_user.IsAdminOrHasRight(claim["sub"].(string))
	if err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if !is_admin {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, entity.ErrUnauthorizedAction.Error(), nil)
		return
	}

	if err := ctx.ShouldBindJSON(&data); err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}

	cat, err := c.svc_category.GetByID(id)
	if err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}

	cat.Title = data.Title
	err = c.svc_category.UpdateTitle(cat)
	if err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}
	utils.ResponseJSON(ctx, http.StatusOK, http.StatusAccepted, "successfully updated", cat)
}


func (c *controller) updateParentCategory(ctx *gin.Context) {
	var data *entity.Category

	claim, err := jwttoken.ExtractClaimsFromAccess(ctx)
	if err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}

	is_admin, err := c.svc_user.IsAdminOrHasRight(claim["sub"].(string))
	if err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if !is_admin {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, entity.ErrUnauthorizedAction.Error(), nil)
		return
	}

	if err := ctx.ShouldBindJSON(&data); err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}

	cat, err := c.svc_category.GetByID(id)
	if err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}

	cat.ParentID = data.ParentID
	err = c.svc_category.UpdateParent(cat)
	if err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}
	utils.ResponseJSON(ctx, http.StatusOK, http.StatusAccepted, "successfully updated", cat)
}

func (c *controller) deleteCategory(ctx *gin.Context) {
	id := ctx.Param("id")
	claim, err := jwttoken.ExtractClaimsFromAccess(ctx)

	if err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}

	is_admin, err := c.svc_user.IsAdminOrHasRight(claim["sub"].(string))
	if err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if !is_admin {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, entity.ErrUnauthorizedAction.Error(), nil)
		return
	}

	id_conv, err := strconv.Atoi(id)
	if err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}

	cat, err := c.svc_category.GetByID(id_conv)
	if err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}

	err = c.svc_category.Delete(cat)
	if err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}
	utils.ResponseJSON(ctx, http.StatusOK, http.StatusAccepted, "successfully delete", nil)
}