package routes

import (
	"api-service/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.Use(middlewares.ZapLoggerMiddleware)
	server.Use(middlewares.ErrorHandler)

	server.POST("/SignUpAdmin", signUpAdmin)

	amindAuthenticated := server.Group("/")
	amindAuthenticated.Use(middlewares.Authenticate)
	amindAuthenticated.POST("/CreateUser", createUser)
	amindAuthenticated.GET("/GetUsers", getUsers)
	amindAuthenticated.GET("/GetUser/:userId", getUser)
	amindAuthenticated.PUT("/UpdateUser/:userId", updateUser)
	amindAuthenticated.DELETE("/DeleteUser/:userId", deleteUser)
}
