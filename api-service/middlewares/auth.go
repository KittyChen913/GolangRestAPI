package middlewares

import (
	"golangrestapi/customerrors"

	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.Error(customerrors.NewAuthenticationError("Not authorized."))
		context.Abort()
		return
	}

	// do some token validation tasks...

	tokenAdminId := 1
	context.Set("adminId", tokenAdminId)

	context.Next()
}
