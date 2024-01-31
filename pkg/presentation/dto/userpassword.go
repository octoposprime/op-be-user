package presentation

import (
	"fmt"

	pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/user"
	tuuid "github.com/octoposprime/op-be-shared/tool/uuid"
	me "github.com/octoposprime/op-be-user/internal/domain/model/entity"
	mo "github.com/octoposprime/op-be-user/internal/domain/model/object"
)

// UserPassword is a struct that represents the dto of a user password values.
type UserPassword struct {
	proto *pb.UserPassword
}

// NewUserPassword creates a new *UserPassword.
func NewUserPassword(pb *pb.UserPassword) *UserPassword {
	return &UserPassword{
		proto: pb,
	}
}

// String returns a string representation of the User.
func (s *UserPassword) String() string {
	return fmt.Sprintf("Id: %v, "+
		"UserId: %v, "+
		"Password: ***, "+
		"PasswordStatus: %v",
		s.proto.Id,
		s.proto.UserId,
		s.proto.PasswordStatus)
}

// NewUserPasswordFromEntity creates a new *UserPassword from entity.
func NewUserPasswordFromEntity(entity *me.UserPassword) *UserPassword {
	return &UserPassword{
		&pb.UserPassword{
			Id:             entity.Id.String(),
			UserId:         entity.UserId.String(),
			Password:       entity.Password,
			PasswordStatus: pb.PasswordStatus(entity.PasswordStatus),
		},
	}
}

// ToPb returns a protobuf representation of the UserPassword.
func (s *UserPassword) ToPb() *pb.UserPassword {
	return s.proto
}

// ToEntity returns a entity representation of the UserPassword.
func (s *UserPassword) ToEntity() *me.UserPassword {
	return &me.UserPassword{
		Id:     tuuid.FromString(s.proto.Id),
		UserId: tuuid.FromString(s.proto.UserId),
		UserPassword: mo.UserPassword{
			Password:       s.proto.Password,
			PasswordStatus: mo.PasswordStatus(s.proto.PasswordStatus),
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

// ToPbs returns a protobuf representation of the UserPasswords.
func (s UserPasswords) ToPbs() *pb.UserPasswords {
	userPasswords := make([]*pb.UserPassword, len(s))
	for i, userPassword := range s {
		userPasswords[i] = userPassword.proto
	}
	return &pb.UserPasswords{
		UserPasswords: userPasswords,
	}
}
