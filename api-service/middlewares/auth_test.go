package middlewares

import (
	"api-service/customerrors"
	"api-service/models"
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
	Authenticate(&MockAdminService{})(c)

	// Assert
	assert.True(t, c.IsAborted(), "Context should be aborted")
	assert.NotEmpty(t, c.Errors, "An error should be attached to the context")
	lastError := c.Errors.Last().Err
	var authErr *customerrors.AuthenticationError
	assert.True(t, errors.As(lastError, &authErr), "The error should be an AuthenticationError")
	assert.Equal(t, "AuthenticationError: Not authorized.", lastError.Error())
}

// TestAuthenticate_Success_WithAuthHeader tests that the request should succeed and set adminName when an Authorization header is provided.
func TestAuthenticate_Success_WithAuthHeader(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Arrange
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("POST", "/GetUsers", nil)
	c.Request.Header.Set("Authorization", "Bearer some-valid-token")

	// Act
	Authenticate(&MockAdminService{})(c)

	// Assert
	assert.False(t, c.IsAborted(), "Context should not be aborted")
	adminName, exists := c.Get("adminName")
	assert.True(t, exists, "adminName should exist in context")
	assert.Equal(t, "mockadmin", adminName)
}

// MockAdminService implements AdminService for testing
type MockAdminService struct{}

func (m *MockAdminService) SignUpAdmin(admin *models.Admin) error {
	return nil
}

func (m *MockAdminService) QueryAdmin(id int) (*models.Admin, error) {
	return &models.Admin{Name: "mockadmin"}, nil
}
