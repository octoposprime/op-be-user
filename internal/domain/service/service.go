package domain

import (
	"strings"
	"unicode"

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
	if strings.Contains(user.UserName, " ") {
		return mo.ErrorUserUsernameIsNotValid
	}
	if !s.CheckUserNameSpecialCharacters(user.UserName) {
		return mo.ErrorUserUsernameIsNotValid
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

func (s *Service) CheckUserNameSpecialCharacters(username string) bool {
	for _, char := range username {
		if !unicode.IsLetter(char) && !unicode.IsNumber(char) && char != '-' && char != '_' {
			return false
		}
	}
	return true
}
