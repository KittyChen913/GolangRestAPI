package middlewares

import (
	"api-service/customerrors"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(context *gin.Context) {
	context.Next()

	if len(context.Errors) > 0 {
		err := context.Errors.Last().Err

		switch {
		case errors.As(err, &customerrors.AuthenticationError{}):
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		default:
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		}
	}
}
