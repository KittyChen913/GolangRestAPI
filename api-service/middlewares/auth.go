package middlewares

import (
	"api-service/customerrors"
	"api-service/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate(adminService services.AdminService) gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.Request.Header.Get("Authorization")
		if token == "" {
			context.Error(customerrors.NewAuthenticationError("Not authorized."))
			context.Abort()
			return
		}

		// do some token validation tasks...

		tokenAdminId := 1
		adminInfo, err := adminService.QueryAdmin(tokenAdminId)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Admin validation failed. [error] : " + err.Error()})
			context.Abort()
			return
		}
		context.Set("adminName", adminInfo.Name)
		context.Next()
	}
}
