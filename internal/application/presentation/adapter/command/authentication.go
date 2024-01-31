package application

import (
	"context"

	mo "github.com/octoposprime/op-be-user/internal/domain/model/object"
)

// Login generates an authentication token if the given users are valid.
func (a CommandAdapter) Login(ctx context.Context, loginRequest mo.LoginRequest) (mo.Token, error) {
	return a.Service.Login(ctx, loginRequest)
}

// Refresh regenerate an authentication token.
func (a CommandAdapter) Refresh(ctx context.Context, token mo.Token) (mo.Token, error) {
	return a.Service.Refresh(ctx, token)
}

// Logout clears some footprints for the user.
func (a CommandAdapter) Logout(ctx context.Context, token mo.Token) error {
	return a.Service.Logout(ctx, token)
}
