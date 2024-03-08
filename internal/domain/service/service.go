package domain

import (
    "strings"

    me "github.com/octoposprime/op-be-user/internal/domain/model/entity"
    mo "github.com/octoposprime/op-be-user/internal/domain/model/object"
)

// This is the service layer of the domain layer.
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

func (s *Service) CheckIsAuthenticable(user *me.User) error {
    if user.UserStatus == mo.UserStatusINACTIVE {
        return mo.ErrorUserIsInactive
    }
    return nil
}
