package infrastructure

import (
	"fmt"

	pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/user"
	tuuid "github.com/octoposprime/op-be-shared/tool/uuid"
	me "github.com/octoposprime/op-be-user/internal/domain/model/entity"
	mo "github.com/octoposprime/op-be-user/internal/domain/model/object"
)

// User is a struct that represents the ebus mapper of a user basic values.
type User struct {
	proto *pb.User
}

// NewUser creates a new *User.
func NewUser(pb *pb.User) *User {
	return &User{
		proto: pb,
	}
}

// String returns a string representation of the User.
func (s *User) String() string {
	return fmt.Sprintf("Id: %v, "+
		"UserName: %v, "+
		"Email: %v, "+
		"UserType: %v, "+
		"UserStatus: %v, "+
		"Tags: %v, "+
		"FirstName: %v, "+
		"LastName: %v",
		s.proto.Id,
		s.proto.Username,
		s.proto.Email,
		s.proto.UserType,
		s.proto.UserStatus,
		s.proto.Tags,
		s.proto.FirstName,
		s.proto.LastName)
}

// NewUserFromEntity creates a new *User from entity.
func NewUserFromEntity(entity me.User) *User {
	return &User{
		&pb.User{
			Id:         entity.Id.String(),
			Username:   entity.UserName,
			Email:      entity.Email,
			UserType:   pb.UserType(entity.UserType),
			UserStatus: pb.UserStatus(entity.UserStatus),
			Tags:       entity.Tags,
			FirstName:  entity.FirstName,
			LastName:   entity.LastName,
		},
	}
}

// ToPb returns a protobuf representation of the User.
func (s *User) ToPb() *pb.User {
	return s.proto
}

// ToEntity returns a entity representation of the User.
func (s *User) ToEntity() *me.User {
	return &me.User{
		Id: tuuid.FromString(s.proto.Id),
		User: mo.User{
			UserName:   s.proto.Username,
			Email:      s.proto.Email,
			UserType:   mo.UserType(s.proto.UserType),
			UserStatus: mo.UserStatus(s.proto.UserStatus),
			Tags:       s.proto.Tags,
			FirstName:  s.proto.FirstName,
			LastName:   s.proto.LastName,
		},
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

// ToPbs returns a protobuf representation of the Users.
func (s Users) ToPbs() *pb.Users {
	users := make([]*pb.User, len(s))
	for i, user := range s {
		users[i] = user.ToPb()
	}
	return &pb.Users{
		Users: users,
	}
}
