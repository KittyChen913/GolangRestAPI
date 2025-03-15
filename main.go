package main

import (
	"golangrestapi/db"
	"golangrestapi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	db.InitDb()

	server.POST("/CreateUser", createUser)
	server.Run()
}

func createUser(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input parameters."})
		return
	}
	err = user.Insert()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Create user failed."})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User created.", "user": user})
}
