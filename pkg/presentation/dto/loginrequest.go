package presentation

import (
	"fmt"

	pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/authentication"
	mo "github.com/octoposprime/op-be-user/internal/domain/model/object"
)

// LoginRequest is a struct that represents the dto of the required values of login for authentication.
type LoginRequest struct {
	proto *pb.LoginRequest
}

// NewLoginRequest creates a new *LoginRequest.
func NewLoginRequest(pb *pb.LoginRequest) *LoginRequest {
	return &LoginRequest{
		proto: pb,
	}
}

// String returns a string representation of the LoginRequest.
func (s *LoginRequest) String() string {
	return fmt.Sprintf("UserName: %v, "+
		"Email: %v, "+
		"Password: ***",
		s.proto.Username,
		s.proto.Email)
}

// NewLoginRequestFromObject creates a new *LoginRequest from object.
func NewLoginRequestFromObject(object *mo.LoginRequest) *LoginRequest {
	return &LoginRequest{
		&pb.LoginRequest{
			Username: object.UserName,
			Email:    object.Email,
			Password: object.Password,
		},
	}
}

// ToPb returns a protobuf representation of the LoginRequest.
func (s *LoginRequest) ToPb() *pb.LoginRequest {
	return s.proto
}

// ToEntity returns a object representation of the LoginRequest.
func (s *LoginRequest) ToObject() *mo.LoginRequest {
	return &mo.LoginRequest{
		UserName: s.proto.Username,
		Email:    s.proto.Email,
		Password: s.proto.Password,
	}
}
