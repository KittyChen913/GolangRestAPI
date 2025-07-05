package main

import (
	"golangrestapi/db"
	_ "golangrestapi/docs"
	"golangrestapi/routes"

	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	server := gin.Default()
	db.InitDb()

	// Swagger settings
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	routes.RegisterRoutes(server)
	server.Run()
}
