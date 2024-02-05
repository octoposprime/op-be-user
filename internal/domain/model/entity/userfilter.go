package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	mo "github.com/octoposprime/op-be-user/internal/domain/model/object"
)

// UserFilter is a struct that represents the filter of a user.
type UserFilter struct {
	Id         uuid.UUID     `json:"id"`          // Id is the id of the user.
	UserName   string        `json:"user_name"`   // UserName is the user name of the user.
	Email      string        `json:"email"`       // Email is the email address of the user.
	UserType   mo.UserType   `json:"user_type"`   // UserType is the type of the user.
	UserStatus mo.UserStatus `json:"user_status"` // UserStatus is the status of the user.
	Tags       []string      `json:"tags"`        // Tags is the tags of the user.
	FirstName  string        `json:"first_name"`  // FirstName is the first name of the user.
	LastName   string        `json:"last_name"`   // LastName is the last name of the user.

	CreatedAtFrom time.Time `json:"created_at_from"` // CreatedAt is in the between of CreatedAtFrom and CreatedAtTo.
	CreatedAtTo   time.Time `json:"created_at_to"`   // CreatedAt is in the between of CreatedAtFrom and CreatedAtTo.
	UpdatedAtFrom time.Time `json:"updated_at_from"` // UpdatedAt is in the between of UpdatedAtFrom and UpdatedAtTo.
	UpdatedAtTo   time.Time `json:"updated_at_to"`   // UpdatedAt is in the between of UpdatedAtFrom and UpdatedAtTo.

	SearchText string           `json:"search_text"` // SearchText is the full-text search value.
	SortType   string           `json:"sort_type"`   // SortType is the sorting type (ASC,DESC).
	SortField  mo.UserSortField `json:"sort_field"`  // SortField is the sorting field of the user.

	Limit  int `json:"limit"`  // Limit provides to limitation row size.
	Offset int `json:"offset"` // Offset provides a starting row number of the limitation.
}

// NewUserFilter creates a new *UserFilter.
func NewUserFilter(id uuid.UUID,
	compnayId uuid.UUID,
	userName string,
	email string,
	userType mo.UserType,
	userStatus mo.UserStatus,
	tags []string,
	firstName string,
	lastName string,
	createdAtFrom time.Time,
	createdAtTo time.Time,
	updatedAtFrom time.Time,
	updatedAtTo time.Time,
	searchText string,
	sortType string,
	sortField mo.UserSortField,
	limit int,
	offset int) *UserFilter {
	return &UserFilter{
		Id:            id,
		UserName:      userName,
		Email:         email,
		UserType:      userType,
		UserStatus:    userStatus,
		Tags:          tags,
		FirstName:     firstName,
		LastName:      lastName,
		CreatedAtFrom: createdAtFrom,
		CreatedAtTo:   createdAtTo,
		UpdatedAtFrom: updatedAtFrom,
		UpdatedAtTo:   updatedAtTo,
		SearchText:    searchText,
		SortType:      sortType,
		SortField:     sortField,
		Limit:         limit,
		Offset:        offset,
	}
}

// NewEmptyUserFilter creates a new *UserFilter with empty values.
func NewEmptyUserFilter() *UserFilter {
	return &UserFilter{
		Id:            uuid.UUID{},
		UserName:      "",
		Email:         "",
		UserType:      mo.UserTypeNONE,
		UserStatus:    mo.UserStatusNONE,
		Tags:          []string{},
		FirstName:     "",
		LastName:      "",
		CreatedAtFrom: time.Time{},
		CreatedAtTo:   time.Time{},
		UpdatedAtFrom: time.Time{},
		UpdatedAtTo:   time.Time{},
		SearchText:    "",
		SortType:      "",
		SortField:     mo.UserSortFieldNONE,
		Limit:         0,
		Offset:        0,
	}
}

// String returns a string representation of the UserFilter.
func (s *UserFilter) String() string {
	return fmt.Sprintf("Id: %v, "+
		"UserName: %v, "+
		"Email: %v, "+
		"UserType: %v, "+
		"UserStatus: %v, "+
		"Tags: %v, "+
		"FirstName: %v, "+
		"LastName: %v, "+
		"CreatedAtFrom: %v, "+
		"CreatedAtTo: %v, "+
		"UpdatedAtFrom: %v, "+
		"UpdatedAtTo: %v, "+
		"SearchText: %v, "+
		"SortType: %v, "+
		"SortField: %v, "+
		"Limit: %v, "+
		"Offset: %v",
		s.Id,
		s.UserName,
		s.Email,
		s.UserType,
		s.UserStatus,
		s.Tags,
		s.FirstName,
		s.LastName,
		s.CreatedAtFrom,
		s.CreatedAtTo,
		s.UpdatedAtFrom,
		s.UpdatedAtTo,
		s.SearchText,
		s.SortType,
		s.SortField,
		s.Limit,
		s.Offset)
}
