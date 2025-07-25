package customerrors

import "fmt"

type AuthenticationError struct {
	message string
}

func (e *AuthenticationError) Error() string {
	return fmt.Sprintf("AuthenticationError: %s", e.message)
}

func NewAuthenticationError(errorMessage string) error {
	return &AuthenticationError{
		message: errorMessage,
	}
}
