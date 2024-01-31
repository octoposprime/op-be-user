package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	mo "github.com/octoposprime/op-be-user/internal/domain/model/object"
)

// User is a struct that represents the entity of a user basic values.
type User struct {
	Id        uuid.UUID `json:"id"`         // Id is the id of the user.
	CompanyId uuid.UUID `json:"company_id"` // CompanyId is the id of the user's company.
	mo.User             // User is the basic values of the user.

	// Only for view
	CreatedAt time.Time `json:"created_at"` // CreatedAt is the create time.
	UpdatedAt time.Time `json:"updated_at"` // UpdatedAt is the update time.
}

// NewUser creates a new *User.
func NewUser(id uuid.UUID,
	companyId uuid.UUID,
	user mo.User) *User {
	return &User{
		Id:        id,
		CompanyId: companyId,
		User:      user,
	}
}

// NewEmptyUser creates a new *User with empty values.
func NewEmptyUser() *User {
	return &User{
		Id:        uuid.UUID{},
		CompanyId: uuid.UUID{},
		User:      *mo.NewEmptyUser(),
	}
}

// String returns a string representation of the User.
func (s *User) String() string {
	return fmt.Sprintf("Id: %v, "+
		"CompanyId: %v, "+
		"User: %v",
		s.Id,
		s.CompanyId,
		s.User)
}

// Equals returns true if the User is equal to the other User.
func (s *User) Equals(other *User) bool {
	if s.Id != other.Id {
		return false
	}
	if s.CompanyId != other.CompanyId {
		return false
	}
	if !s.User.Equals(&other.User) {
		return false
	}
	return true
}

// Clone returns a clone of the User.
func (s *User) Clone() *User {
	return &User{
		Id:        s.Id,
		CompanyId: s.CompanyId,
		User:      *s.User.Clone(),
	}
}

// IsEmpty returns true if the User is empty.
func (s *User) IsEmpty() bool {
	if s.Id.String() != "" && s.Id != (uuid.UUID{}) {
		return false
	}
	if s.CompanyId.String() != "" && s.CompanyId != (uuid.UUID{}) {
		return false
	}
	if !s.User.IsEmpty() {
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
	s.Id = uuid.UUID{}
	s.CompanyId = uuid.UUID{}
	s.User.Clear()
}

// Validate validates the User.
func (s *User) Validate() error {
	if s.IsEmpty() {
		return mo.ErrorUserIsEmpty
	}
	if s.CompanyId.String() == "" || s.CompanyId != (uuid.UUID{}) {
		return mo.ErrorUserCompanyIdIsEmpty
	}
	if err := s.User.Validate(); err != nil {
		return err
	}
	return nil
}

// Users contains a slice of *User and total number of users.
type Users struct {
	Users     []User `json:"users"`      // Users is the slice of *User.
	TotalRows int64  `json:"total_rows"` // TotalRows is the total number of rows.
}
