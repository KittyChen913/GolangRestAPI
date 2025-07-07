package middlewares

import (
	"api-service/logger"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ZapLoggerMiddleware(context *gin.Context) {
	start := time.Now()

	context.Next()

	defer func() {
		if err := recover(); err != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err})
		}
	}()

	duration := time.Since(start)

	var errMsg string
	if len(context.Errors) > 0 {
		errMsg = context.Errors.Last().Err.Error()
	} else {
		errMsg = ""
	}

	logger.Log.Info(errMsg,
		zap.String("service.name", "api-service"),
		zap.String("method", context.Request.Method),
		zap.String("path", context.Request.URL.Path),
		zap.Int("status", context.Writer.Status()),
		zap.String("ip", context.ClientIP()),
		zap.Duration("latency", duration),
		zap.String("user-agent", context.Request.UserAgent()),
	)
}
