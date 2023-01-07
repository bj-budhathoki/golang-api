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

// @BasePath /api/v1
// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {object} utils.Response
// @Router /auth/login [post]
func (c *authController) Login(ctx *gin.Context) {
	var loginDTO dtos.LoginDTOS
	errDTO := ctx.ShouldBind(&loginDTO)
	if errDTO != nil {
		response := utils.BuildErrorResponse("Failed to process request", errDTO.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
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

// @BasePath /api/v1
// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 201 {object} utils.Response
// @Router /auth/register [post]
func (c *authController) Register(ctx *gin.Context) {
	var resiterDTOS dtos.RegisterDTOS
	errDTO := ctx.ShouldBind(&resiterDTOS)
	if errDTO != nil {
		response := utils.BuildErrorResponse("Failed to process request", errDTO.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if !c.authService.IsDuplicateEmail(resiterDTOS.Email) {
		response := utils.BuildErrorResponse("Fail to process request", "Duplicate email", utils.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
	} else {
		newUser := c.authService.CreateUser(dtos.UserCreateDTOS(resiterDTOS))
		token := c.jwtService.GenerateToken(strconv.FormatUint(newUser.ID, 10))
		newUser.Token = token
		response := utils.BuildResponse(true, "OK", newUser)
		ctx.JSON(http.StatusCreated, response)
	}
}
