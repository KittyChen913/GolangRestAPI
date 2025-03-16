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
		context.Error(err)
		return
	}

	adminId := context.GetInt("adminId")

	err = user.Insert()
	if err != nil {
		context.Error(fmt.Errorf("[Admin : %v] create user failed", adminId))
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": fmt.Sprintf("[Admin : %v] user created", adminId), "user": user})
}

func getUsers(context *gin.Context) {
	adminId := context.GetInt("adminId")

	users, err := models.Query()
	if err != nil {
		context.Error(fmt.Errorf("[Admin : %v] could not fetch users", adminId))
		return
	}
	context.JSON(http.StatusOK, users)
}

func getUser(context *gin.Context) {
	userId, err := strconv.ParseInt(context.Param("userId"), 10, 32)
	if err != nil {
		context.Error(fmt.Errorf("could not parse user id"))
		return
	}

	adminId := context.GetInt("adminId")

	user, err := models.QueryById(int(userId))
	if err != nil {
		context.Error(fmt.Errorf("[Admin : %v] could not fetch user", adminId))
		return
	}
	context.JSON(http.StatusOK, user)
}

func updateUser(context *gin.Context) {
	userId, err := strconv.ParseInt(context.Param("userId"), 10, 32)
	if err != nil {
		context.Error(fmt.Errorf("could not parse user id"))
		return
	}

	adminId := context.GetInt("adminId")

	_, err = models.QueryById(int(userId))
	if err != nil {
		context.Error(fmt.Errorf("[Admin : %v] could not fetch the user", adminId))
		return
	}

	var updatedUser models.User
	err = context.ShouldBindJSON(&updatedUser)
	if err != nil {
		context.Error(err)
		return
	}
	updatedUser.Id = int(userId)

	err = updatedUser.Update()
	if err != nil {
		context.Error(fmt.Errorf("[Admin : %v] could not update user", adminId))
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": fmt.Sprintf("[Admin : %v] User updated.", adminId)})
}

func deleteUser(context *gin.Context) {
	userId, err := strconv.ParseInt(context.Param("userId"), 10, 32)
	if err != nil {
		context.Error(fmt.Errorf("could not parse user id"))
		return
	}

	adminId := context.GetInt("adminId")

	user, err := models.QueryById(int(userId))
	if err != nil {
		context.Error(fmt.Errorf("[Admin : %v] could not fetch user", adminId))
		return
	}
	err = user.Delete()
	if err != nil {
		context.Error(fmt.Errorf("[Admin : %v] delete user failed", adminId))
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("[Admin : %v] User deleted.", adminId)})
}
