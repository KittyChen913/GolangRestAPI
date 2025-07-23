package middlewares

import (
	"api-service/customerrors"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TestAuthenticate_Failure_NoAuthHeader tests that the request should be aborted and an error returned when no Authorization header is provided.
func TestAuthenticate_Failure_NoAuthHeader(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Arrange
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("POST", "/GetUsers", nil)

	// Act
	Authenticate(c)

	// Assert
	assert.True(t, c.IsAborted(), "Context should be aborted")
	assert.NotEmpty(t, c.Errors, "An error should be attached to the context")
	lastError := c.Errors.Last().Err
	var authErr *customerrors.AuthenticationError
	assert.True(t, errors.As(lastError, &authErr), "The error should be an AuthenticationError")
	assert.Equal(t, "AuthenticationError: Not authorized.", lastError.Error())
}

// TestAuthenticate_Success_WithAuthHeader tests that the request should succeed and set adminId when an Authorization header is provided.
func TestAuthenticate_Success_WithAuthHeader(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Arrange
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("POST", "/GetUsers", nil)
	c.Request.Header.Set("Authorization", "Bearer some-valid-token")

	// Act
	Authenticate(c)

	// Assert
	assert.False(t, c.IsAborted(), "Context should not be aborted")
	adminId, exists := c.Get("adminId")
	assert.True(t, exists, "adminId should exist in context")
	assert.GreaterOrEqual(t, adminId, 0, "adminId should not be less than 0")
}
