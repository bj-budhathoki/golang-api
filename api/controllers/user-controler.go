package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/bj-budhathoki/golang-api/api/services"
	"github.com/bj-budhathoki/golang-api/dtos"
	"github.com/bj-budhathoki/golang-api/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type UserController interface {
	Update(ctx *gin.Context)
	Profile(ctx *gin.Context)
}
type userController struct {
	userService services.UserService
	jwtService  services.JWTService
}

func NewUserContoller(userService services.UserService, jwtService services.JWTService) UserController {
	return &userController{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (c *userController) Update(ctx *gin.Context) {
	var UserUpdateDTOS dtos.UserUpdateDTOS
	errDTO := ctx.ShouldBind(&UserUpdateDTOS)
	if errDTO != nil {
		res := utils.BuildErrorResponse("Failed to process request", errDTO.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	authHeadr := ctx.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeadr)
	if errToken != nil {
		panic(errToken)
	}
	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	UserUpdateDTOS.ID = id
	u := c.userService.Update((UserUpdateDTOS))
	res := utils.BuildResponse(true, "OK", u)
	ctx.JSON(http.StatusOK, res)
}

func (c *userController) Profile(ctx *gin.Context) {
	authHeadr := ctx.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeadr)
	if errToken != nil {
		panic(errToken)
	}
	claims := token.Claims.(jwt.MapClaims)
	user := c.userService.Profile(fmt.Sprintf("%v", claims["user_id"]))
	res := utils.BuildResponse(true, "OK", user)
	ctx.JSON(http.StatusOK, res)
}
