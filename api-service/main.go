package main

import (
	"api-service/db"
	_ "api-service/docs"
	"api-service/handlers"
	"api-service/logger"
	"api-service/repositories"
	"api-service/routes"
	"api-service/services"

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

	adminRepo := repositories.NewAdminRepository(db.Db)
	adminService := services.NewAdminService(adminRepo)
	adminHandler := handlers.NewAdminHandler(adminService)

	userRepo := repositories.NewUserRepository(db.Db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	routes.RegisterRoutes(server, adminHandler, userHandler)
	server.Run()
}
