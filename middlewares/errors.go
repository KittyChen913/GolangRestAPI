package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(context *gin.Context) {
	context.Next()

	if len(context.Errors) > 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": strings.Join(context.Errors.Errors(), " | ")})
	}
}
