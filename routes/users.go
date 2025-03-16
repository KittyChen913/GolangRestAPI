package routes

import (
	"fmt"
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

	adminId := context.GetInt("adminId")

	err = user.Insert()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("[Admin : %v] Create user failed.", adminId)})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": fmt.Sprintf("[Admin : %v] User created.", adminId), "user": user})
}

func getUsers(context *gin.Context) {
	adminId := context.GetInt("adminId")

	users, err := models.Query()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("[Admin : %v] Could not fetch users.", adminId)})
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

	adminId := context.GetInt("adminId")

	user, err := models.QueryById(int(userId))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("[Admin : %v] Could not fetch user.", adminId)})
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

	adminId := context.GetInt("adminId")

	_, err = models.QueryById(int(userId))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("[Admin : %v] Could not fetch the user.", adminId)})
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
		context.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("[Admin : %v] Could not update user.", adminId)})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": fmt.Sprintf("[Admin : %v] User updated.", adminId)})
}

func deleteUser(context *gin.Context) {
	userId, err := strconv.ParseInt(context.Param("userId"), 10, 32)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse user id."})
		return
	}

	adminId := context.GetInt("adminId")

	user, err := models.QueryById(int(userId))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("[Admin : %v] Could not fetch user.", adminId)})
		return
	}
	err = user.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("[Admin : %v] Delete user failed.", adminId)})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("[Admin : %v] User deleted.", adminId)})
}
