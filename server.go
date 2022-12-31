package main

import (
	"github.com/bj-budhathoki/golang-api/api/controllers"
	"github.com/bj-budhathoki/golang-api/api/repository"
	"github.com/bj-budhathoki/golang-api/api/services"
	"github.com/bj-budhathoki/golang-api/infrastructure"
	"github.com/bj-budhathoki/golang-api/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                   = infrastructure.SetupDatabase()
	userRepository repository.UserRepository  = repository.NewUserRepository(db)
	jwtService     services.JWTService        = services.NewJWTService()
	authService    services.AuthService       = services.NewAuthService(userRepository)
	userService    services.UserService       = services.NewUserService(userRepository)
	authController controllers.AuthController = controllers.NewAuthController(authService, jwtService)
	userController controllers.UserController = controllers.NewUserContoller(userService, jwtService)
)

func main() {
	defer infrastructure.CloseDatabaseConnection(db)
	r := gin.Default()
	authRoutes := r.Group("/api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}
	userRouter := r.Group("/api/user", middlewares.AuthorizeJWT(jwtService))
	{
		userRouter.PUT("/update", userController.Update)
		userRouter.GET("/profile", userController.Profile)
	}
	r.Run()
}
