package domain

import (
	"fmt"
)

// Token is a struct that represents the object of a token for authentication.
type Token struct {
	AuthenticationToken string `json:"authentication_token"` // AuthenticationToken is the authentication token.
	RefreshToken        string `json:"refresh_token"`        // RefreshToken is the refresh token.
}

// NewToken creates a new *Token.
func NewToken(authenticationToken string,
	refreshToken string) *Token {
	return &Token{
		AuthenticationToken: authenticationToken,
		RefreshToken:        refreshToken,
	}
}

// NewEmptyToken creates a new *Token with empty values.
func NewEmptyToken() *Token {
	return &Token{
		AuthenticationToken: "",
		RefreshToken:        "",
	}
}

// String returns a string representation of the Token.
func (s *Token) String() string {
	return fmt.Sprintf("AuthenticationToken: %v, "+
		"RefreshToken: %v",
		s.AuthenticationToken,
		s.RefreshToken)
}

// Equals returns true if the Token is equal to the other Token.
func (s *Token) Equals(other *Token) bool {
	if s.AuthenticationToken != other.AuthenticationToken {
		return false
	}
	if s.RefreshToken != other.RefreshToken {
		return false
	}
	return true
}

// Clone returns a clone of the Token.
func (s *Token) Clone() *Token {
	return &Token{
		AuthenticationToken: s.AuthenticationToken,
		RefreshToken:        s.RefreshToken,
	}
}

// IsEmpty returns true if the Token is empty.
func (s *Token) IsEmpty() bool {
	if s.AuthenticationToken != "" {
		return false
	}
	if s.RefreshToken != "" {
		return false
	}
	return true
}

// IsNotEmpty returns true if the Token is not empty.
func (s *Token) IsNotEmpty() bool {
	return !s.IsEmpty()
}

// Clear clears the Token.
func (s *Token) Clear() {
	s.AuthenticationToken = ""
	s.RefreshToken = ""
}

// Validate validates the Token.
func (s *Token) Validate() error {
	if s.IsEmpty() {
		return ErrorUserTokenIsEmpty
	}
	if s.AuthenticationToken == "" && s.RefreshToken == "" {
		return ErrorUserAuthenticationTokenAndRefreshTokenAreEmpty
	}
	return nil
}
