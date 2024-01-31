package presentation

import (
	"fmt"

	pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/authentication"
	mo "github.com/octoposprime/op-be-user/internal/domain/model/object"
)

// Token is a struct that represents the dto of a token for authentication.
type Token struct {
	proto *pb.Token
}

// NewToken creates a new *Token.
func NewToken(pb *pb.Token) *Token {
	return &Token{
		proto: pb,
	}
}

// String returns a string representation of the Token.
func (s *Token) String() string {
	return fmt.Sprintf("AuthenticationToken: %v, "+
		"RefreshToken: %v",
		s.proto.AuthenticationToken,
		s.proto.RefreshToken)
}

// NewTokenFromObject creates a new *Token from object.
func NewTokenFromObject(object *mo.Token) *Token {
	return &Token{
		&pb.Token{
			AuthenticationToken: object.AuthenticationToken,
			RefreshToken:        object.RefreshToken,
		},
	}
}

// ToPb returns a protobuf representation of the Token.
func (s *Token) ToPb() *pb.Token {
	return s.proto
}

// ToEntity returns a object representation of the Token.
func (s *Token) ToObject() *mo.Token {
	return &mo.Token{
		AuthenticationToken: s.proto.AuthenticationToken,
		RefreshToken:        s.proto.RefreshToken,
	}
}
