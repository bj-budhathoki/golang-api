package routes

import (
	"github.com/bj-budhathoki/golang-api/api/controllers"
	"github.com/gin-gonic/gin"
)

var (
	authController controllers.AuthController = controllers.NewAuthController()
)

func AuthRoutes(r *gin.Engine) {
	authRoutes := r.Group("/api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}
}
