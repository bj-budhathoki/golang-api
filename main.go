package main

import (
	"net/http"
	"time"

	"github.com/bj-budhathoki/golang-api/api/controllers"
	"github.com/bj-budhathoki/golang-api/api/repository"
	"github.com/bj-budhathoki/golang-api/api/services"
	docs "github.com/bj-budhathoki/golang-api/docs"
	"github.com/bj-budhathoki/golang-api/infrastructure"
	"github.com/bj-budhathoki/golang-api/middlewares"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                    = infrastructure.SetupDatabase()
	userRepository repository.UserRepository   = repository.NewUserRepository(db)
	bookRepsitory  repository.BookRepository   = repository.NewBookRepository(db)
	jwtService     services.JWTService         = services.NewJWTService()
	authService    services.AuthService        = services.NewAuthService(userRepository)
	userService    services.UserService        = services.NewUserService(userRepository)
	bookService    services.BookService        = services.NewBookService(bookRepsitory)
	authController controllers.AuthController  = controllers.NewAuthController(authService, jwtService)
	userController controllers.UserController  = controllers.NewUserContoller(userService, jwtService)
	bookController controllers.BookControlller = controllers.NewBookController(bookService, jwtService)
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/
// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host      localhost:8080
// @BasePath  /api
// @securityDefinitions.basic  BasicAuth

func main() {
	defer infrastructure.CloseDatabaseConnection(db)
	startTime := time.Now()
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		v1.GET("health-check", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"status": "OK",
				"uptime": time.Since(startTime),
			})
		})
		authRoutes := v1.Group("/auth")
		{
			authRoutes.POST("/login", authController.Login)
			authRoutes.POST("/register", authController.Register)
		}
		userRouter := v1.Group("user", middlewares.AuthorizeJWT(jwtService))
		{
			userRouter.PUT("/update", userController.Update)
			userRouter.GET("/profile", userController.Profile)
		}
		bookRouter := v1.Group("/books", middlewares.AuthorizeJWT(jwtService))
		{
			bookRouter.GET("/", bookController.AllBook)
			bookRouter.POST("/:id", bookController.FindBookById)
			bookRouter.POST("/", bookController.CreateBook)
			bookRouter.PUT("/:id", bookController.UpdateBook)
			bookRouter.DELETE("/:id", bookController.DeleteBook)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run()
}
