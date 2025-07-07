package routes

import (
	"api-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SignUpAdmin godoc
// @Summary Admin 註冊
// @Description 註冊新的管理員帳號
// @Tags admin
// @Accept json
// @Produce json
// @Param admin body models.Admin true "管理員資訊"
// @Success 201 {object} map[string]interface{} "註冊成功"
// @Failure 400 {object} map[string]string "參數錯誤"
// @Failure 500 {object} map[string]string "註冊失敗"
// @Router /SignUpAdmin [post]
func signUpAdmin(context *gin.Context) {
	var admin models.Admin
	err := context.ShouldBindJSON(&admin)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input parameters."})
		return
	}
	err = admin.Insert()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Admin registration failed."})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Admin registration successfully.", "admin": admin})
}
