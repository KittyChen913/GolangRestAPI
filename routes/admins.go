package routes

import (
	"golangrestapi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
