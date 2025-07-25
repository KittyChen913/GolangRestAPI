package routes

import (
	"api-service/handlers"
	"api-service/middlewares"
	"api-service/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine,
	adminHandler *handlers.AdminHandler,
	userHandler *handlers.UserHandler,
	adminService services.AdminService) {

	server.Use(middlewares.ZapLoggerMiddleware)
	server.Use(middlewares.ErrorHandler)

	server.POST("/SignUpAdmin", adminHandler.SignUpAdmin)

	amindAuthenticated := server.Group("/")
	amindAuthenticated.Use(middlewares.Authenticate(adminService))
	amindAuthenticated.POST("/CreateUser", userHandler.CreateUser)
	amindAuthenticated.GET("/GetUsers", userHandler.GetUsers)
	amindAuthenticated.GET("/GetUser/:userId", userHandler.GetUser)
	amindAuthenticated.PUT("/UpdateUser/:userId", userHandler.UpdateUser)
	amindAuthenticated.DELETE("/DeleteUser/:userId", userHandler.DeleteUser)
}
