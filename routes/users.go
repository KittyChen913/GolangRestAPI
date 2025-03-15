package routes

import (
	"golangrestapi/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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

func getUsers(context *gin.Context) {
	users, err := models.Query()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch users."})
		return
	}
	context.JSON(http.StatusOK, users)
}

func getUser(context *gin.Context) {
	userId, err := strconv.ParseInt(context.Param("userId"), 10, 32)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse user id."})
		return
	}
	user, err := models.QueryById(int(userId))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch user."})
		return
	}
	context.JSON(http.StatusOK, user)
}

func updateUser(context *gin.Context) {
	userId, err := strconv.ParseInt(context.Param("userId"), 10, 32)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse user id."})
		return
	}

	_, err = models.QueryById(int(userId))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the user."})
		return
	}

	var updatedUser models.User
	err = context.ShouldBindJSON(&updatedUser)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input parameters."})
		return
	}
	updatedUser.Id = int(userId)

	err = updatedUser.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update user."})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User updated."})
}

func deleteUser(context *gin.Context) {
	userId, err := strconv.ParseInt(context.Param("userId"), 10, 32)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse user id."})
		return
	}
	user, err := models.QueryById(int(userId))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch user."})
		return
	}
	err = user.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Delete user failed."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "User deleted."})
}
