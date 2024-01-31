package domain

import (
	"fmt"
)

// User is a struct that represents the object of a user basic values.
type User struct {
	UserName   string     `json:"user_name"`   // UserName is the user name of the user.
	Email      string     `json:"email"`       // Email is the email address of the user.
	Role       string     `json:"role"`        // Role is the role of the user.
	UserType   UserType   `json:"user_type"`   // UserType is the type of the user.
	UserStatus UserStatus `json:"user_status"` // UserStatus is the status of the user.
	Tags       []string   `json:"tags"`        // Tags is the tags of the user.
	FirstName  string     `json:"first_name"`  // FirstName is the first name of the user.
	LastName   string     `json:"last_name"`   // LastName is the last name of the user.
}

// NewUser creates a new *User.
func NewUser(userName string,
	email string,
	role string,
	userType UserType,
	userStatus UserStatus,
	tags []string,
	firstName string,
	lastName string) *User {
	return &User{
		UserName:   userName,
		Email:      email,
		Role:       role,
		UserType:   userType,
		UserStatus: userStatus,
		Tags:       tags,
		FirstName:  firstName,
		LastName:   lastName,
	}
}

// NewEmptyUser creates a new *User with empty values.
func NewEmptyUser() *User {
	return &User{
		UserName:   "",
		Email:      "",
		Role:       "",
		UserType:   UserTypeNONE,
		UserStatus: UserStatusNONE,
		Tags:       []string{},
		FirstName:  "",
		LastName:   "",
	}
}

// String returns a string representation of the User.
func (s *User) String() string {
	return fmt.Sprintf("UserName: %v, "+
		"Email: %v, "+
		"Role: %v, "+
		"UserType: %v, "+
		"UserStatus: %v, "+
		"Tags: %v, "+
		"FirstName: %v, "+
		"LastName: %v",
		s.UserName,
		s.Email,
		s.Role,
		s.UserType,
		s.UserStatus,
		s.Tags,
		s.FirstName,
		s.LastName)
}

// Equals returns true if the User is equal to the other User.
func (s *User) Equals(other *User) bool {
	if s.UserName != other.UserName {
		return false
	}
	if s.Email != other.Email {
		return false
	}
	if s.Role != other.Role {
		return false
	}
	if s.UserType != other.UserType {
		return false
	}
	if s.UserStatus != other.UserStatus {
		return false
	}
	for i := range s.Tags {
		if s.Tags[i] != other.Tags[i] {
			return false
		}
	}
	if s.FirstName != other.FirstName {
		return false
	}
	if s.LastName != other.LastName {
		return false
	}
	return true
}

// Clone returns a clone of the User.
func (s *User) Clone() *User {
	return &User{
		UserName:   s.UserName,
		Email:      s.Email,
		Role:       s.Role,
		UserType:   s.UserType,
		UserStatus: s.UserStatus,
		Tags:       s.Tags,
		FirstName:  s.FirstName,
		LastName:   s.LastName,
	}
}

// IsEmpty returns true if the User is empty.
func (s *User) IsEmpty() bool {
	if s.UserName != "" {
		return false
	}
	if s.Email != "" {
		return false
	}
	if s.Role != "" {
		return false
	}
	if s.UserType != UserTypeNONE {
		return false
	}
	if s.UserStatus != UserStatusNONE {
		return false
	}
	if len(s.Tags) != 0 {
		return false
	}
	if s.FirstName != "" {
		return false
	}
	if s.LastName != "" {
		return false
	}
	return true
}

// IsNotEmpty returns true if the User is not empty.
func (s *User) IsNotEmpty() bool {
	return !s.IsEmpty()
}

// Clear clears the User.
func (s *User) Clear() {
	s.UserName = ""
	s.Email = ""
	s.Role = ""
	s.UserType = UserTypeNONE
	s.UserStatus = UserStatusNONE
	s.Tags = []string{}
	s.FirstName = ""
	s.LastName = ""
}

// Validate validates the User.
func (s *User) Validate() error {
	if s.IsEmpty() {
		return ErrorUserIsEmpty
	}
	if s.UserName == "" {
		return ErrorUserUsernameIsEmpty
	}
	if s.Email == "" {
		return ErrorUserEmailIsEmpty
	}
	if s.Role == "" {
		return ErrorUserRoleIsEmpty
	}
	return nil
}
