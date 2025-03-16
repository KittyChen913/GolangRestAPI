package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Not authorized."})
		return
	}

	// do some token validation tasks...

	tokenAdminId := 1
	context.Set("adminId", tokenAdminId)

	context.Next()
}
