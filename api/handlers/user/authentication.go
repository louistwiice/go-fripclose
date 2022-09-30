package handler_user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/louistwiice/go/fripclose/configs"
	"github.com/louistwiice/go/fripclose/domain"
	"github.com/louistwiice/go/fripclose/entity"
	"github.com/louistwiice/go/fripclose/utils"
	jwttoken "github.com/louistwiice/go/fripclose/utils/jwt_token"
)

type authcontroller struct {
	service domain.AuthService
}

func NewAuthController(svc domain.AuthService) *authcontroller {
	return &authcontroller{
		service: svc,
	}
}

// Register/Create a new user or account
func (c *authcontroller) register(ctx *gin.Context) {
	var user entity.UserCreateUpdate
	if err := ctx.ShouldBindJSON(&user); err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}

	err := c.service.Create(&user)
	if err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}

	response := entity.UserDisplayFormater(&user)
	go c.service.SetActivationCode(*response)

	utils.ResponseJSON(ctx, http.StatusOK, http.StatusCreated, "successful", response)
}

func (c *authcontroller) activation(ctx *gin.Context) {
	var input entity.UserActivation
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}

	resp, err := c.service.GetActivationCode(input.Username)
	if err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if resp != input.Code {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, "Wrong Token. Please try again", nil)
		return
	}

	err = c.service.ActivateUser(input.Username)
	if err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}

	user, _, err := c.service.SearchUser(input.Username)
	if err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}
	utils.ResponseJSON(ctx, http.StatusOK, http.StatusCreated, "successful activated", user)
}

// Login is used to connect to the API
func (c *authcontroller) login(ctx *gin.Context) {
	conf := configs.LoadConfigEnv()
	var input entity.UserLogin

	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}

	user, hashed_password, err := c.service.SearchUser(input.Identifier)
	if err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, entity.ErrUserNotFound.Error(), nil)
		return
	}

	if !user.IsActive {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, entity.ErrUserDeactivated.Error(), nil)
		return
	}

	err = utils.CheckHashedString(input.Password, hashed_password)
	if err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}

	tokens, err := jwttoken.GenerateToken(user.ID, user.IsStaff, user.IsSuperuser)
	if err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}

	ctx.SetCookie("access_token", tokens["access_token"], 1, "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", tokens["refresh_token"], 1, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "true", 1, "/", "localhost", false, false)

	go c.service.UpdateAuthenticationDate(user)

	utils.ResponseJSON(ctx, http.StatusOK, http.StatusOK, "Login successfully", gin.H{"token": tokens, "duration": conf.AccessTokenHourLifespan, "token_prefix": conf.TokenPrefix})
}

func (c *authcontroller) refreshToken(ctx *gin.Context) {
	type tokenReqBody struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}

	conf := configs.LoadConfigEnv()
	tokenReq := tokenReqBody{}

	if err := ctx.ShouldBindJSON(&tokenReq); err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}

	claim, err := jwttoken.ExtractClaimsFromRefresh(tokenReq.RefreshToken)
	if err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}

	user, _, err := c.service.GetByID(claim["sub"].(string))
	if err != nil || user == nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, entity.ErrUserNotFound.Error(), nil)
		return
	}

	tokens, err := jwttoken.RefreshToken(ctx, tokenReq.RefreshToken)
	if err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}

	ctx.SetCookie("access_token", tokens["access_token"], 1, "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", tokens["refresh_token"], 1, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "true", 1, "/", "localhost", false, false)

	utils.ResponseJSON(ctx, http.StatusOK, http.StatusOK, "refresh successfully", gin.H{"token": tokens, "duration": conf.AccessTokenHourLifespan, "token_prefix": conf.TokenPrefix})

}

func (c *authcontroller) logout(ctx *gin.Context) {
	ctx.SetCookie("access_token", "", -1, "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", "", -1, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "", -1, "/", "localhost", false, true)

	utils.ResponseJSON(ctx, http.StatusOK, http.StatusOK, "logout successfully", nil)

}
