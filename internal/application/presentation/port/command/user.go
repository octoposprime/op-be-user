package application

import (
	"context"

	me "github.com/octoposprime/op-be-user/internal/domain/model/entity"
)

// CommandPort is a port for Hexagonal Architecture Pattern.
// It is used to communicate with the application layer.
type UserCommandPort interface {
	// CreateUser sends the given user to the application layer for creating a new user.
	CreateUser(ctx context.Context, user me.User) (me.User, error)

	// UpdateUserBase sends the given base values of the user to the repository of the infrastructure layer for updating base values of user data.
	UpdateUserBase(ctx context.Context, user me.User) (me.User, error)

	// UpdateUserStatus sends the given status value of the user to the repository of the infrastructure layer for updating status of user data.
	UpdateUserStatus(ctx context.Context, user me.User) (me.User, error)

	// UpdateUserRole sends the given type value of the user to the repository of the infrastructure layer for updating role of user data.
	UpdateUserRole(ctx context.Context, user me.User) (me.User, error)

	// DeleteUser sends the given user to the application layer for deleting data.
	DeleteUser(ctx context.Context, user me.User) (me.User, error)

	// ChangePassword sends the given user password to the application layer for changing user password.
	ChangePassword(ctx context.Context, userPassword me.UserPassword) error
}
