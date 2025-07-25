package routes

import (
	"api-service/handlers"
	"api-service/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine, handler *handlers.UserHandler) {
	server.Use(middlewares.ZapLoggerMiddleware)
	server.Use(middlewares.ErrorHandler)

	server.POST("/SignUpAdmin", signUpAdmin)

	amindAuthenticated := server.Group("/")
	amindAuthenticated.Use(middlewares.Authenticate)
	amindAuthenticated.POST("/CreateUser", handler.CreateUser)
	amindAuthenticated.GET("/GetUsers", handler.GetUsers)
	amindAuthenticated.GET("/GetUser/:userId", handler.GetUser)
	amindAuthenticated.PUT("/UpdateUser/:userId", handler.UpdateUser)
	amindAuthenticated.DELETE("/DeleteUser/:userId", handler.DeleteUser)
}
