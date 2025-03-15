package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {

	server.POST("/CreateUser", createUser)
	server.GET("/GetUsers", getUsers)
	server.GET("/GetUser/:userId", getUser)
	server.PUT("/UpdateUser/:userId", updateUser)
}
