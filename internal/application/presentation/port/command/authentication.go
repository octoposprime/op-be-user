package application

import (
	"context"

	mo "github.com/octoposprime/op-be-user/internal/domain/model/object"
)

// CommandPort is a port for Hexagonal Architecture Pattern.
// It is used to communicate with the application layer.
type AuthenticationCommandPort interface {
	// Login generates an authentication token if the given login request values are valid.
	Login(ctx context.Context, loginRequest mo.LoginRequest) (mo.Token, error)

	// Refresh regenerate an authentication token.
	Refresh(ctx context.Context, token mo.Token) (mo.Token, error)

	// Logout clears some footprints for the user.
	Logout(ctx context.Context, token mo.Token) error
}
