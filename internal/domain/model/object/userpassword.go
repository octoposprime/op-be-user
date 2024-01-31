package domain

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// UserPassword is a struct that represents the object of a user password values.
type UserPassword struct {
	Password       string         `json:"password"`        // Password is the password of the user.
	PasswordStatus PasswordStatus `json:"password_status"` // PasswordStatus is the status of password of the user.
}

// NewUserPassword creates a new *UserPassword.
func NewUserPassword(password string,
	passwordStatus PasswordStatus) *UserPassword {
	return &UserPassword{
		Password:       password,
		PasswordStatus: passwordStatus,
	}
}

// NewEmptyUserPassword creates a new *UserPassword with empty values.
func NewEmptyUserPassword() *UserPassword {
	return &UserPassword{
		Password:       "",
		PasswordStatus: PasswordStatusNONE,
	}
}

// String returns a string representation of the UserPassword.
func (s *UserPassword) String() string {
	return fmt.Sprintf("Password: ***, "+
		"PasswordStatus: %v",
		s.PasswordStatus)
}

// Equals returns true if the User is equal to the other UserPassword.
func (s *UserPassword) Equals(other *UserPassword) bool {
	if s.Password != other.Password {
		return false
	}
	if s.PasswordStatus != PasswordStatusNONE {
		return false
	}
	return true
}

// ClonePassword returns a clone of the UserPassword.
func (s *UserPassword) Clone() *UserPassword {
	return &UserPassword{
		Password:       s.Password,
		PasswordStatus: s.PasswordStatus,
	}
}

// IsEmpty returns true if the UserPassword is empty.
func (s *UserPassword) IsEmpty() bool {
	if s.Password != "" {
		return false
	}
	if s.PasswordStatus != PasswordStatusNONE {
		return false
	}
	return true
}

// IsNotEmpty returns true if the UserPassword is not empty.
func (s *UserPassword) IsNotEmpty() bool {
	return !s.IsEmpty()
}

// Clear clears the UserPassword.
func (s *UserPassword) Clear() {
	s.Password = ""
	s.PasswordStatus = PasswordStatusNONE
}

// Validate validates the UserPassword.
func (s *UserPassword) Validate() error {
	if s.IsEmpty() {
		return ErrorUserPasswordIsEmpty
	}
	return nil
}

// ComparePass compares hashes of the passwords
func (s *UserPassword) ComparePass(pass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(s.Password), []byte(pass)) == nil
}
