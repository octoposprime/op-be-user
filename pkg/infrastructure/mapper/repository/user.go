package infrastructure

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/lib/pq"
	tgorm "github.com/octoposprime/op-be-shared/tool/gorm"
	me "github.com/octoposprime/op-be-user/internal/domain/model/entity"
	mo "github.com/octoposprime/op-be-user/internal/domain/model/object"
)

// User is a struct that represents the db mapper of a user basic values.
type User struct {
	tgorm.Model

	UserName   string         `json:"user_name" gorm:"not null;default:''"`  // UserName is the user name of the user.
	Email      string         `json:"email" gorm:"not null;default:''"`      // Email is the email address of the user.
	Role       string         `json:"role" gorm:"not null;default:''"`       // Role is the role of the user.
	UserType   int            `json:"user_type" gorm:"not null;default:0"`   // UserType is the type of the user.
	UserStatus int            `json:"user_status" gorm:"not null;default:0"` // UserStatus is the status of the user.
	Tags       pq.StringArray `json:"tags" gorm:"type:text[]"`               // Tags is the tags of the user.
	FirstName  string         `json:"first_name" gorm:"not null;default:''"` // FirstName is the first name of the user.
	LastName   string         `json:"last_name" gorm:"not null;default:''"`  // FirstName is the last name of the user.
}

// NewUser creates a new *User.
func NewUser(id uuid.UUID,
	userName string,
	email string,
	role string,
	userType int,
	userStatus int,
	tags pq.StringArray,
	firstName string,
	lastName string) *User {
	return &User{
		Model:      tgorm.Model{ID: id},
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

// String returns a string representation of the User.
func (s *User) String() string {
	return fmt.Sprintf("Id: %v, "+
		"UserName: %v, "+
		"Email: %v, "+
		"Role: %v, "+
		"UserType: %v, "+
		"UserStatus: %v, "+
		"Tags: %v, "+
		"FirstName: %v, "+
		"LastName: %v",
		s.ID,
		s.UserName,
		s.Email,
		s.Role,
		s.UserType,
		s.UserStatus,
		s.Tags,
		s.FirstName,
		s.LastName)
}

// NewUserFromEntity creates a new *User from entity.
func NewUserFromEntity(entity me.User) *User {
	return &User{
		Model:      tgorm.Model{ID: entity.Id},
		UserName:   entity.UserName,
		Email:      entity.Email,
		Role:       entity.Role,
		UserType:   int(entity.UserType),
		UserStatus: int(entity.UserStatus),
		Tags:       entity.Tags,
		FirstName:  entity.FirstName,
		LastName:   entity.LastName,
	}
}

// ToEntity returns a entity representation of the User.
func (s *User) ToEntity() *me.User {
	return &me.User{
		Id: s.ID,
		User: mo.User{
			UserName:   s.UserName,
			Email:      s.Email,
			Role:       s.Role,
			UserType:   mo.UserType(s.UserType),
			UserStatus: mo.UserStatus(s.UserStatus),
			Tags:       s.Tags,
			FirstName:  s.FirstName,
			LastName:   s.LastName,
		},
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
	}
}

type Users []*User

// NewUsersFromEntities creates a new []*User from entities.
func NewUserFromEntities(entities []me.User) Users {
	users := make([]*User, len(entities))
	for i, entity := range entities {
		users[i] = NewUserFromEntity(entity)
	}
	return users
}

// ToEntities creates a new []me.User entity.
func (s Users) ToEntities() []me.User {
	users := make([]me.User, len(s))
	for i, user := range s {
		users[i] = *user.ToEntity()
	}
	return users
}
