package controllers

import (
	"net/http"
	"strconv"

	"github.com/bj-budhathoki/golang-api/api/services"
	"github.com/bj-budhathoki/golang-api/dtos"
	"github.com/bj-budhathoki/golang-api/model"
	"github.com/bj-budhathoki/golang-api/utils"
	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}
type authController struct {
	authService services.AuthService
	jwtService  services.JWTService
}

func NewAuthController(authService services.AuthService, jwtService services.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}
func (c *authController) Login(ctx *gin.Context) {
	var loginDTO dtos.LoginDTOS
	errDTO := ctx.ShouldBind(&loginDTO)
	if errDTO != nil {
		response := utils.BuildErrorResponse("Failed to process request", errDTO.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)

	}
	authResult := c.authService.VerfiyCredential(loginDTO.Email, loginDTO.Password)
	if v, ok := authResult.(model.User); ok {
		generatedToken := c.jwtService.GenerateToken(strconv.FormatUint(v.ID, 10))
		v.Token = generatedToken
		response := utils.BuildResponse(true, "OK", v)
		ctx.JSON(http.StatusOK, response)
		return
	}
	response := utils.BuildErrorResponse("Please check your credentail", "Invalid credentail", utils.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (c *authController) Register(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "hello resister",
	})
}
