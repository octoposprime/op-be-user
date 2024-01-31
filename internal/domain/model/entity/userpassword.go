package domain

import (
	"fmt"

	"github.com/google/uuid"
	mo "github.com/octoposprime/op-be-user/internal/domain/model/object"
)

// UserPassword is a struct that represents the entity of a user password values.
type UserPassword struct {
	Id              uuid.UUID `json:"id"`         // Id is the id of the password.
	UserId          uuid.UUID `json:"user_id"`    // Id is the id of the user.
	CompanyId       uuid.UUID `json:"company_id"` // Id is the id of the user's company.
	mo.UserPassword           // UserPassword is the password values of the user.
}

// NewUserPassword creates a new *UserPassword.
func NewUserPassword(id uuid.UUID,
	userId uuid.UUID,
	companyId uuid.UUID,
	userPassword mo.UserPassword) *UserPassword {
	return &UserPassword{
		Id:           id,
		UserId:       userId,
		CompanyId:    companyId,
		UserPassword: userPassword,
	}
}

// NewEmptyUserPassword creates a new *UserPassword with empty values.
func NewEmptyUserPassword() *UserPassword {
	return &UserPassword{
		Id:           uuid.UUID{},
		UserId:       uuid.UUID{},
		CompanyId:    uuid.UUID{},
		UserPassword: *mo.NewEmptyUserPassword(),
	}
}

// String returns a string representation of the UserPassword.
func (s *UserPassword) String() string {
	return fmt.Sprintf("Id: %v, "+
		"UserId: %v, "+
		"CompanyId: %v, "+
		"UserPassword: %v",
		s.Id,
		s.UserId,
		s.CompanyId,
		s.UserPassword)
}

// Equals returns true if the UserPassword is equal to the other UserPassword.
func (s *UserPassword) Equals(other *UserPassword) bool {
	if s.Id != other.Id {
		return false
	}
	if s.UserId != other.UserId {
		return false
	}
	if s.CompanyId != other.CompanyId {
		return false
	}
	if !s.UserPassword.Equals(&other.UserPassword) {
		return false
	}
	return true
}

// Clone returns a clone of the UserPassword.
func (s *UserPassword) Clone() *UserPassword {
	return &UserPassword{
		Id:           s.Id,
		UserId:       s.UserId,
		CompanyId:    s.CompanyId,
		UserPassword: *s.UserPassword.Clone(),
	}
}

// IsEmpty returns true if the UserPassword is empty.
func (s *UserPassword) IsEmpty() bool {
	if s.Id.String() != "" && s.Id != (uuid.UUID{}) {
		return false
	}
	if s.UserId.String() != "" && s.UserId != (uuid.UUID{}) {
		return false
	}
	if s.CompanyId.String() != "" && s.CompanyId != (uuid.UUID{}) {
		return false
	}
	if !s.UserPassword.IsEmpty() {
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
	s.Id = uuid.UUID{}
	s.UserId = uuid.UUID{}
	s.CompanyId = uuid.UUID{}
	s.UserPassword.Clear()
}

// Validate validates the UserPassword.
func (s *UserPassword) Validate() error {
	if s.IsEmpty() {
		return mo.ErrorUserIsEmpty
	}
	if s.UserId.String() == "" || s.UserId == (uuid.UUID{}) {
		return mo.ErrorUserIdIsEmpty
	}
	if s.CompanyId.String() == "" || s.CompanyId == (uuid.UUID{}) {
		return mo.ErrorUserCompanyIdIsEmpty
	}
	if err := s.UserPassword.Validate(); err != nil {
		return err
	}
	return nil
}
