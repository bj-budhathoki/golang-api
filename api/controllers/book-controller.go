package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/bj-budhathoki/golang-api/api/services"
	"github.com/bj-budhathoki/golang-api/dtos"
	"github.com/bj-budhathoki/golang-api/model"
	"github.com/bj-budhathoki/golang-api/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type BookControlller interface {
	CreateBook(ctx *gin.Context)
	UpdateBook(ctx *gin.Context)
	AllBook(ctx *gin.Context)
	DeleteBook(ctx *gin.Context)
	FindBookById(ctx *gin.Context)
}

type bookController struct {
	bookService services.BookService
	jwtService  services.JWTService
}

func NewBookController(service services.BookService, jwtServ services.JWTService) BookControlller {
	return &bookController{
		bookService: service,
		jwtService:  jwtServ,
	}
}

func (c *bookController) CreateBook(ctx *gin.Context) {
	var bookDTOS dtos.BookCreateDTOS
	errDTO := ctx.ShouldBind(&bookDTOS)
	if errDTO != nil {
		response := utils.BuildErrorResponse("Failed to process request", errDTO.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	} else {
		authHeader := ctx.GetHeader("Authorization")
		userID := c.getUserIDByToken(authHeader)
		convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			bookDTOS.UserID = convertedUserID
		}
		result := c.bookService.CreateBook(bookDTOS)
		response := utils.BuildResponse(true, "OK", result)
		ctx.JSON(http.StatusCreated, response)
	}

}

func (c *bookController) UpdateBook(ctx *gin.Context) {
	var bookUpdateDTO dtos.BookUpdateDTOS
	errDTO := ctx.ShouldBind(&bookUpdateDTO)
	if errDTO != nil {
		res := utils.BuildErrorResponse("Failed to process request", errDTO.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	id, errID := strconv.ParseUint(userID, 10, 64)
	if c.bookService.IsAllowedToEdit(userID, bookUpdateDTO.ID) {
		if errID == nil {
			bookUpdateDTO.UserID = id
		}
		result := c.bookService.UpdateBook(bookUpdateDTO)
		response := utils.BuildResponse(true, "OK", result)
		ctx.JSON(http.StatusOK, response)
	} else {
		response := utils.BuildErrorResponse("You dont have permission", "You are not the owner", utils.EmptyObj{})
		ctx.JSON(http.StatusForbidden, response)
	}
}

// @BasePath /api/
// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {array} []model.Book
// @Router /api/v1/books [get]
func (c *bookController) AllBook(ctx *gin.Context) {
	var books []model.Book = c.bookService.AllBook()
	res := utils.BuildResponse(true, "OK", books)
	ctx.JSON(http.StatusOK, res)
}
func (c *bookController) FindBookById(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		res := utils.BuildErrorResponse("No param was found", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	var book model.Book = c.bookService.FindBookById(id)
	if (book == model.Book{}) {
		res := utils.BuildErrorResponse("Data not found", "No data with given id", utils.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
	} else {
		res := utils.BuildResponse(true, "Ok", book)
		ctx.JSON(http.StatusOK, res)
	}
}
func (c *bookController) DeleteBook(context *gin.Context) {
	var book model.Book
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := utils.BuildErrorResponse("Failed tou get id", "No param id were found", utils.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	book.ID = id
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.bookService.IsAllowedToEdit(userID, book.ID) {
		c.bookService.DeleteBook(book)
		res := utils.BuildResponse(true, "Deleted", utils.EmptyObj{})
		context.JSON(http.StatusOK, res)
	} else {
		response := utils.BuildErrorResponse("You dont have permission", "You are not the owner", utils.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *bookController) getUserIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
