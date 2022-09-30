package handler_user

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/louistwiice/go/fripclose/domain"
	"github.com/louistwiice/go/fripclose/entity"
	"github.com/louistwiice/go/fripclose/utils"
	jwttoken "github.com/louistwiice/go/fripclose/utils/jwt_token"
)

type controller struct {
	service domain.UserService
}

func NewUserController(svc domain.UserService) *controller {
	return &controller{
		service: svc,
	}
}

func (c *controller) listUsers(ctx *gin.Context) {
	users, err := c.service.List()
	if err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}
	utils.ResponseJSON(ctx, http.StatusOK, http.StatusOK, "successful", users)
}

func (c *controller) getUser(ctx *gin.Context) {
	id := ctx.Param("id")

	user, _, err := c.service.GetByID(id)
	if err != nil || user == nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}
	utils.ResponseJSON(ctx, http.StatusOK, http.StatusFound, "successful", user)
}

func (c *controller) connectedUser(ctx *gin.Context) {
	claim, err := jwttoken.ExtractClaimsFromAccess(ctx)
	if err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}

	// Only the Owner or admin is allowed to delete an account
	user, _, err := c.service.GetByID(claim["sub"].(string))
	if err != nil || user == nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}
	utils.ResponseJSON(ctx, http.StatusOK, http.StatusFound, "connected user found", user)
}

func (c *controller) updateUser(ctx *gin.Context) {
	var data *entity.UserCreateUpdate
	var id = ctx.Param("id")

	if err := ctx.ShouldBindJSON(&data); err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}

	user, _, err := c.service.GetByID(id)
	if err != nil || user == nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, entity.ErrNotFound.Error(), nil)
		return
	}

	data = entity.ValidateUpdate(data, user)

	err = c.service.UpdateUser(data)
	if err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}
	response := entity.UserDisplayFormater(data)
	utils.ResponseJSON(ctx, http.StatusOK, http.StatusAccepted, "successful", response)
}

// Update user password. Only the owner can update his password
func (c *controller) updatePassword(ctx *gin.Context) {
	var data entity.ChangePassword
	var id = ctx.Param("id")

	claim, err := jwttoken.ExtractClaimsFromAccess(ctx)
	if err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if err := ctx.ShouldBindJSON(&data); err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}

	user, password, err := c.service.GetByID(id)
	if err != nil || user == nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, entity.ErrNotFound.Error(), nil)
		return
	}

	user_claim, _, _ := c.service.GetByID(claim["sub"].(string))
	if user_claim.ID != user.ID {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, entity.ErrUnauthorizedAction.Error(), nil)
		return
	}

	err = utils.CheckHashedString(data.OldPassword, password)
	if err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, "old password does not match", err.Error())
		return
	}

	udapte_user := &entity.UserCreateUpdate{
		UserDisplay: *user,
		Password:    data.NewPassword,
	}

	err = c.service.UpdatePassword(udapte_user)
	if err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}
	utils.ResponseJSON(ctx, http.StatusOK, http.StatusOK, "successful", "Password reset successfully")
}

func (c *controller) deleteUser(ctx *gin.Context) {
	var data *entity.UserCreateUpdate
	var id = ctx.Param("id")

	if err := ctx.ShouldBindJSON(&data); err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}

	claim, err := jwttoken.ExtractClaimsFromAccess(ctx)
	if err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}

	// Only the Owner or admin is allowed to delete an account
	user_claim, _, _ := c.service.GetByID(claim["sub"].(string))
	is_admin, _ := c.service.IsAdminOrHasRight(claim["sub"].(string))
	if user_claim.ID != id && !is_admin {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, entity.ErrUnauthorizedAction.Error(), nil)
		return
	}

	err = c.service.Delete(id)
	if err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}
	utils.ResponseJSON(ctx, http.StatusOK, http.StatusAccepted, "successfully deleted", nil)
}

func (c *controller) uploadProfileImage(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}

	claim, err := jwttoken.ExtractClaimsFromAccess(ctx)
	if err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}

	user, _, err := c.service.GetByID(claim["sub"].(string))
	if err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}

	hashed, err := utils.HashString(file.Filename)
	if err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}

	clean_hash := strings.ReplaceAll(hashed[12:30], ".", "")
	clean_hash = strings.ReplaceAll(clean_hash, "/", "")
	clean_hash = strings.ReplaceAll(clean_hash, "\\", "")

	save_to := fmt.Sprintf("/media/users/%s__%s", clean_hash, file.Filename)
	user.Picture = save_to
	err = c.service.UploadPicture(user)
	if err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}

	ctx.SaveUploadedFile(file, "."+save_to)

	utils.ResponseJSON(ctx, http.StatusOK, http.StatusAccepted, "Picture upload successfully", user)
}
