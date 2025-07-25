package handlers

import (
	"api-service/models"
	"api-service/services"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// createUser 建立新使用者
// @Summary 建立新使用者
// @Description 由管理員建立新使用者
// @Tags User
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param user body models.User true "使用者資料"
// @Success 201 {object} models.User
// @Failure 400 {object} error
// @Router /CreateUser [post]
func (handler *UserHandler) CreateUser(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.Error(err)
		return
	}
	adminId := context.GetInt("adminId")
	user.CreateDateTime = time.Now()

	err = handler.userService.CreateUser(&user)
	if err != nil {
		context.Error(fmt.Errorf("[Admin : %v] create user failed. [error] : %v", adminId, err))
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": fmt.Sprintf("[Admin : %v] user created", adminId), "user": user})
}

// getUsers 取得所有使用者
// @Summary 取得所有使用者
// @Description 由管理員查詢所有使用者
// @Tags User
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Success 200 {array} models.User
// @Failure 400 {object} error
// @Router /GetUsers [get]
func (handler *UserHandler) GetUsers(context *gin.Context) {
	adminId := context.GetInt("adminId")
	users, err := handler.userService.GetUsers()
	if err != nil {
		context.Error(fmt.Errorf("[Admin : %v] could not fetch users. [error] : %v", adminId, err))
		return
	}
	context.JSON(http.StatusOK, users)
}

// getUser 取得單一使用者
// @Summary 取得單一使用者
// @Description 由管理員查詢指定 ID 的使用者
// @Tags User
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param userId path int true "使用者ID"
// @Success 200 {object} models.User
// @Failure 400 {object} error
// @Router /GetUser/{userId} [get]
func (handler *UserHandler) GetUser(context *gin.Context) {
	userId, err := strconv.ParseInt(context.Param("userId"), 10, 32)
	if err != nil {
		context.Error(fmt.Errorf("could not parse user id"))
		return
	}

	adminId := context.GetInt("adminId")

	user, err := handler.userService.GetUser(int(userId))
	if err != nil {
		context.Error(fmt.Errorf("[Admin : %v] could not fetch user. [error] : %v", adminId, err))
		return
	}
	context.JSON(http.StatusOK, user)
}

// updateUser 更新使用者
// @Summary 更新使用者
// @Description 由管理員更新指定 ID 的使用者資料
// @Tags User
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param userId path int true "使用者ID"
// @Param user body models.User true "使用者資料"
// @Success 201 {object} string
// @Failure 400 {object} error
// @Router /UpdateUser/{userId} [put]
func (handler *UserHandler) UpdateUser(context *gin.Context) {
	userId, err := strconv.ParseInt(context.Param("userId"), 10, 32)
	if err != nil {
		context.Error(fmt.Errorf("could not parse user id. [error] : %v", err))
		return
	}

	adminId := context.GetInt("adminId")

	_, err = handler.userService.GetUser(int(userId))
	if err != nil {
		context.Error(fmt.Errorf("[Admin : %v] could not fetch the user. [error] : %v", adminId, err))
		return
	}

	var updatedUser models.User
	err = context.ShouldBindJSON(&updatedUser)
	if err != nil {
		context.Error(err)
		return
	}
	updatedUser.Id = int(userId)

	err = handler.userService.UpdateUser(int(userId), &updatedUser)
	if err != nil {
		context.Error(fmt.Errorf("[Admin : %v] could not update user. [error] : %v", adminId, err))
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": fmt.Sprintf("[Admin : %v] User updated.", adminId)})
}

// deleteUser 刪除使用者
// @Summary 刪除使用者
// @Description 由管理員刪除指定 ID 的使用者
// @Tags User
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param userId path int true "使用者ID"
// @Success 200 {object} string
// @Failure 400 {object} error
// @Router /DeleteUser/{userId} [delete]
func (handler *UserHandler) DeleteUser(context *gin.Context) {
	userId, err := strconv.ParseInt(context.Param("userId"), 10, 32)
	if err != nil {
		context.Error(fmt.Errorf("could not parse user id"))
		return
	}

	adminId := context.GetInt("adminId")

	_, err = handler.userService.GetUser(int(userId))
	if err != nil {
		context.Error(fmt.Errorf("[Admin : %v] could not fetch user. [error] : %v", adminId, err))
		return
	}
	err = handler.userService.DeleteUser(int(userId))
	if err != nil {
		context.Error(fmt.Errorf("[Admin : %v] delete user failed. [error] : %v", adminId, err))
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("[Admin : %v] User deleted.", adminId)})
}
