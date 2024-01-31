package application

import (
	"context"

	"github.com/google/uuid"
	pb_logging "github.com/octoposprime/op-be-shared/pkg/proto/pb/logging"
	me "github.com/octoposprime/op-be-user/internal/domain/model/entity"
)

// DbPort is a port for Hexagonal Architecture Pattern.
// It is used to communicate with the database.
type DbPort interface {
	// SetLogger sets logging call-back function
	SetLogger(LogFunc func(ctx context.Context, logData *pb_logging.LogData) (*pb_logging.LoggingResult, error))

	// GetUsersByFilter returns the users that match the given filter.
	GetUsersByFilter(ctx context.Context, userFilter me.UserFilter) (me.Users, error)

	// SaveUser insert a new user or update the existing one in the database.
	SaveUser(ctx context.Context, user me.User) (me.User, error)

	// DeleteUser soft-deletes the given user in the database.
	DeleteUser(ctx context.Context, user me.User) (me.User, error)

	// GetUserPasswordByUserId returns active password of the given user.
	GetUserPasswordByUserId(ctx context.Context, userId uuid.UUID) (me.UserPassword, error)

	// ChangePassword changes the poassword of the given user in the database.
	ChangePassword(ctx context.Context, userPassword me.UserPassword) (me.UserPassword, error)
}
