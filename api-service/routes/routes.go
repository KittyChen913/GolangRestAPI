package routes

import (
	"api-service/handlers"
	"api-service/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine,
	adminHandler *handlers.AdminHandler,
	userHandler *handlers.UserHandler) {

	server.Use(middlewares.ZapLoggerMiddleware)
	server.Use(middlewares.ErrorHandler)

	server.POST("/SignUpAdmin", adminHandler.SignUpAdmin)

	amindAuthenticated := server.Group("/")
	amindAuthenticated.Use(middlewares.Authenticate)
	amindAuthenticated.POST("/CreateUser", userHandler.CreateUser)
	amindAuthenticated.GET("/GetUsers", userHandler.GetUsers)
	amindAuthenticated.GET("/GetUser/:userId", userHandler.GetUser)
	amindAuthenticated.PUT("/UpdateUser/:userId", userHandler.UpdateUser)
	amindAuthenticated.DELETE("/DeleteUser/:userId", userHandler.DeleteUser)
}
