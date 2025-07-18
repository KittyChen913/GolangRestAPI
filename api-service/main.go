package main

import (
	"api-service/db"
	_ "api-service/docs"
	"api-service/logger"
	"api-service/routes"

	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	logger.InitLogger()
	defer logger.Log.Sync()

	server := gin.Default()
	db.InitDb()

	// Swagger settings
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	routes.RegisterRoutes(server)
	server.Run()
}
