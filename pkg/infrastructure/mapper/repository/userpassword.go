package infrastructure

import (
	"fmt"

	"github.com/google/uuid"
	tgorm "github.com/octoposprime/op-be-shared/tool/gorm"
	me "github.com/octoposprime/op-be-user/internal/domain/model/entity"
	mo "github.com/octoposprime/op-be-user/internal/domain/model/object"
)

// UserPassword is a struct that represents the db mapper of a user password values.
type UserPassword struct {
	tgorm.Model
	UserID         uuid.UUID `json:"user_id" gorm:"type:uuid;default:uuid_nil()"` // UserId is the id of the user.
	Password       string    `json:"password" gorm:"not null"`                    // Password is the password of the user.
	PasswordStatus int       `json:"password_status" gorm:"not null;default:0"`   // PasswordStatus is the status of password of the user.
}

// NewUserPassword creates a new *UserPassword.
func NewUserPassword(id uuid.UUID,
	userId uuid.UUID,
	password string,
	paswordStatus int) *UserPassword {
	return &UserPassword{
		Model:          tgorm.Model{ID: id},
		UserID:         userId,
		Password:       password,
		PasswordStatus: paswordStatus,
	}
}

// String returns a string representation of the UserPassword.
func (s *UserPassword) String() string {
	return fmt.Sprintf("Id: %v, "+
		"UserId: %v, "+
		"Password: ***, "+
		"PasswordStatus: %v",
		s.ID,
		s.UserID,
		s.PasswordStatus)
}

// NewUserPasswordFromEntity creates a new *UserPassword from entity.
func NewUserPasswordFromEntity(entity *me.UserPassword) *UserPassword {
	return &UserPassword{
		Model:          tgorm.Model{ID: entity.Id},
		UserID:         entity.UserId,
		Password:       entity.Password,
		PasswordStatus: int(entity.PasswordStatus),
	}
}

// ToEntity returns a entity representation of the UserPassword.
func (s *UserPassword) ToEntity() *me.UserPassword {
	return &me.UserPassword{
		Id:     s.ID,
		UserId: s.UserID,
		UserPassword: mo.UserPassword{
			Password:       s.Password,
			PasswordStatus: mo.PasswordStatus(s.PasswordStatus),
		},
	}
}

type UserPasswords []*UserPassword

// NewUserPasswordsFromEntities creates a new []*UserPassword from entities.
func NewUserPasswordFromEntities(entities []me.UserPassword) UserPasswords {
	userPasswords := make([]*UserPassword, len(entities))
	for i, entity := range entities {
		userPasswords[i] = NewUserPasswordFromEntity(&entity)
	}
	return userPasswords
}

// ToEntities creates a new []me.UserPassword entity.
func (s UserPasswords) ToEntities() []me.UserPassword {
	userPasswords := make([]me.UserPassword, len(s))
	for i, userPassword := range s {
		userPasswords[i] = *userPassword.ToEntity()
	}
	return userPasswords
}
