package application

import (
	"context"

	me "github.com/octoposprime/op-be-user/internal/domain/model/entity"
)

// CreateUser sends the given user to the application layer for creating a new user.
func (a CommandAdapter) CreateUser(ctx context.Context, user me.User) (me.User, error) {
	return a.Service.CreateUser(ctx, user)
}

// UpdateUserBase sends the given base values of the user to the repository of the infrastructure layer for updating base values of user data.
func (a CommandAdapter) UpdateUserBase(ctx context.Context, user me.User) (me.User, error) {
	return a.Service.UpdateUserBase(ctx, user)
}

// UpdateUserStatus sends the given status value of the user to the repository of the infrastructure layer for updating status of user data.
func (a CommandAdapter) UpdateUserStatus(ctx context.Context, user me.User) (me.User, error) {
	return a.Service.UpdateUserStatus(ctx, user)
}

// UpdateUserRole sends the given type value of the user to the repository of the infrastructure layer for updating role of user data.
func (a CommandAdapter) UpdateUserRole(ctx context.Context, user me.User) (me.User, error) {
	return a.Service.UpdateUserRole(ctx, user)
}

// DeleteUser sends the given user to the application layer for deleting data.
func (a CommandAdapter) DeleteUser(ctx context.Context, user me.User) (me.User, error) {
	return a.Service.DeleteUser(ctx, user)
}

// ChangePassword sends the given user password to the application layer for changing user password.
func (a CommandAdapter) ChangePassword(ctx context.Context, userPassword me.UserPassword) error {
	return a.Service.ChangePassword(ctx, userPassword)
}
