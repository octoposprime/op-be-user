package domain

import (
    "strings"
    "testing"

    me "github.com/octoposprime/op-be-user/internal/domain/model/entity"
    mo "github.com/octoposprime/op-be-user/internal/domain/model/object"
    "github.com/stretchr/testify/require"
)

// Service is the service layer of the domain layer.
type Service struct {
}

// NewService creates a new *Service.
func NewService() *Service {
    return &Service{}
}

// ValidateUser validates the user.
func (s *Service) ValidateUser(user *me.User) error {
    if err := user.Validate(); err != nil {
        return err
    }
    return nil
}

// ValidatePassword validates the user password.
func (s *Service) ValidatePassword(userPassword *me.UserPassword) error {
    if err := userPassword.Validate(); err != nil {
        return err
    }
    return nil
}

// ValidateToken validates the token.
func (s *Service) ValidateToken(token *mo.Token) error {
    return token.Validate()
}

// CheckUserNameRules checks the rules for the username when creating a new user.
func (s *Service) CheckUserNameRules(user *me.User) error {
    if user.UserName == "" {
        return mo.ErrorUserUsernameIsEmpty
    }
    if len(user.UserName) < 8 {
        return mo.ErrorUserUsernameIsTooShort
    }
    if len(user.UserName) > 20 {
        return mo.ErrorUserUsernameIsTooLong
    }
    // Check for spaces
    if strings.Contains(user.UserName, " ") {
        return mo.ErrorUserUsernameContainsSpace
    }
    // Check for special characters
    specialChars := "!@#$%^&*()-+=}{[]|\\:;\"'<>,.?/"
    for _, char := range specialChars {
        if strings.Contains(user.UserName, string(char)) {
            return mo.ErrorUserUsernameContainsSpecialChars
        }
    }
    // Check for valid characters (alphanumeric, underscore, period)
    validChars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789._"
    for _, char := range user.UserName {
        if !strings.Contains(validChars, string(char)) {
            return mo.ErrorUserUsernameContainsInvalidChars
        }
    }
    return nil
}

// CheckEmailRules checks the rules for the email when creating a new user.
func (s *Service) CheckEmailRules(user *me.User) error {
    if user.Email == "" {
        return mo.ErrorUserEmailIsEmpty
    }
    if !strings.Contains(user.Email, ".") {
        return mo.ErrorUserEmailIsNotValid
    }
    if !strings.Contains(user.Email, "@") {
        return mo.ErrorUserEmailIsNotValid
    }
    if strings.Contains(user.Email, " ") {
        return mo.ErrorUserEmailIsNotValid
    }
    return nil
}

// CheckPasswordRules checks the rules for the password when creating a new user.
func (s *Service) CheckPasswordRules(userPassword *me.UserPassword) error {
    if userPassword.Password == "" {
        return mo.ErrorUserPasswordIsEmpty
    }
    if len(userPassword.Password) < 8 {
        return mo.ErrorUserPasswordIsTooShort
    }
    if len(userPassword.Password) > 20 {
        return mo.ErrorUserPasswordIsTooLong
    }
    if strings.Contains(userPassword.Password, " ") {
        return mo.ErrorUserPasswordIsNotValid
    }
    return nil
}

// CheckIsAuthenticable checks the authenticity of the user.
func (s *Service) CheckIsAuthenticable(user *me.User) error {
    if user.UserStatus == mo.UserStatusINACTIVE {
        return mo.ErrorUserIsInactive
    }
    return nil
}

// TestCheckUserNameSpecialCharacters tests the CheckUserNameRules function for special characters in usernames.
func TestCheckUserNameSpecialCharacters(t *testing.T) {
    service := NewService()

    tests := []struct {
        name       string
        userName   string
        expectFail bool
    }{
        {"ContainsSpecialChars", "asd!^+.", true},
        {"NoSpecialChars", "Qwe123_", false},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            user := &me.User{UserName: tt.userName}
            err := service.CheckUserNameRules(user)
            if tt.expectFail {
                require.Error(t, err, "Expected an error for username: %s", tt.userName)
            } else {
                require.NoError(t, err, "Did not expect an error for username: %s", tt.userName)
            }
        })
    }
}
