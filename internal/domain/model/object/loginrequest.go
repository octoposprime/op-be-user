package domain

import (
	"fmt"
)

// LoginRequest is a struct that represents the required values of login.
type LoginRequest struct {
	UserName string `json:"user_name"` // UserName is the user name of the user.
	Email    string `json:"email"`     // Email is the email address of the user.
	Password string `json:"password"`  // Password is the password of the user.
}

// NewLoginRequest creates a new *LoginRequest.
func NewLoginRequest(userName string,
	email string,
	password string) *LoginRequest {
	return &LoginRequest{
		UserName: userName,
		Email:    email,
		Password: password,
	}
}

// NewEmptyLoginRequest creates a new *LoginRequest with empty values.
func NewEmptyLoginRequest() *LoginRequest {
	return &LoginRequest{
		UserName: "",
		Email:    "",
		Password: "",
	}
}

// String returns a string representation of the reqiured values of login.
func (s *LoginRequest) String() string {
	return fmt.Sprintf("UserName: %v, "+
		"Email: %v, "+
		"Password: ***",
		s.UserName,
		s.Email)
}

// Equals returns true if the login values is equal to the other login values.
func (s *LoginRequest) Equals(other *LoginRequest) bool {
	if s.UserName != other.UserName {
		return false
	}
	if s.Email != other.Email {
		return false
	}
	if s.Password != other.Password {
		return false
	}
	return true
}

// Clone returns a clone of the required values of login.
func (s *LoginRequest) Clone() *LoginRequest {
	return &LoginRequest{
		UserName: s.UserName,
		Email:    s.Email,
		Password: s.Password,
	}
}

// IsEmpty returns true if the login values are empty.
func (s *LoginRequest) IsEmpty() bool {
	if s.UserName != "" {
		return false
	}
	if s.Email != "" {
		return false
	}
	if s.Password != "" {
		return false
	}
	return true
}

// IsNotEmpty returns true if the login values are not empty.
func (s *LoginRequest) IsNotEmpty() bool {
	return !s.IsEmpty()
}

// Clear clears the LoginRequest.
func (s *LoginRequest) Clear() {
	s.UserName = ""
	s.Email = ""
	s.Password = ""
}

// Validate validates the LoginRequest.
func (s *LoginRequest) Validate() error {
	if s.IsEmpty() {
		return ErrorUserIsEmpty
	}
	if s.UserName == "" && s.Email == "" {
		return ErrorUserUsernameOrEmailAreEmpty
	}
	if s.Password == "" {
		return ErrorUserPasswordIsEmpty
	}
	return nil
}
