package main

import (
	routes "github.com/bj-budhathoki/golang-api/api/routes"
	"github.com/bj-budhathoki/golang-api/infrastructure"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = infrastructure.SetupDatabase()
)

func main() {
	defer infrastructure.CloseDatabaseConnection(db)
	r := gin.Default()
	routes.AuthRoutes(r)
	r.Run()
}
