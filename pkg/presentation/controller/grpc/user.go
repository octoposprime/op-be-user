package presentation

import (
	"context"

	pb_user "github.com/octoposprime/op-be-shared/pkg/proto/pb/user"
	dto "github.com/octoposprime/op-be-user/pkg/presentation/dto"
)

// GetUsersByFilter returns the users that match the given filter.
func (a *Grpc) GetUsersByFilter(ctx context.Context, filter *pb_user.UserFilter) (*pb_user.Users, error) {
	users, err := a.queryHandler.GetUsersByFilter(ctx, *dto.NewUserFilter(filter).ToEntity())
	return dto.NewUserFromEntities(users).ToPbs(), err
}

// CreateUser sends the given user to the application layer for creating new user.
func (a *Grpc) CreateUser(ctx context.Context, user *pb_user.User) (*pb_user.User, error) {
	data, err := a.commandHandler.CreateUser(ctx, *dto.NewUser(user).ToEntity())
	return dto.NewUserFromEntity(data).ToPb(), err
}

// UpdateUserRole sends the given user to the application layer for updating user role.
func (a *Grpc) UpdateUserRole(ctx context.Context, user *pb_user.User) (*pb_user.User, error) {
	data, err := a.commandHandler.UpdateUserRole(ctx, *dto.NewUser(user).ToEntity())
	return dto.NewUserFromEntity(data).ToPb(), err
}

// UpdateUserBase sends the given user to the application layer for updating user's base values.
func (a *Grpc) UpdateUserBase(ctx context.Context, user *pb_user.User) (*pb_user.User, error) {
	data, err := a.commandHandler.UpdateUserBase(ctx, *dto.NewUser(user).ToEntity())
	return dto.NewUserFromEntity(data).ToPb(), err
}

// UpdateUserStatus sends the given user to the application layer for updating user status.
func (a *Grpc) UpdateUserStatus(ctx context.Context, user *pb_user.User) (*pb_user.User, error) {
	data, err := a.commandHandler.UpdateUserStatus(ctx, *dto.NewUser(user).ToEntity())
	return dto.NewUserFromEntity(data).ToPb(), err
}

// DeleteUser sends the given user to the application layer for deleting data.
func (a *Grpc) DeleteUser(ctx context.Context, user *pb_user.User) (*pb_user.User, error) {
	data, err := a.commandHandler.DeleteUser(ctx, *dto.NewUser(user).ToEntity())
	return dto.NewUserFromEntity(data).ToPb(), err
}

// ChangePassword sends the given user password to the application layer for changing user password.
func (a *Grpc) ChangePassword(ctx context.Context, userPassword *pb_user.UserPassword) (*pb_user.UserPasswordResult, error) {
	err := a.commandHandler.ChangePassword(ctx, *dto.NewUserPassword(userPassword).ToEntity())
	return &pb_user.UserPasswordResult{}, err
}
