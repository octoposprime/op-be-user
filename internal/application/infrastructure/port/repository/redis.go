package application

import (
	"context"

	"github.com/google/uuid"
	pb_logging "github.com/octoposprime/op-be-shared/pkg/proto/pb/logging"
	me "github.com/octoposprime/op-be-user/internal/domain/model/entity"
)

// RedisPort is a port for Hexagonal Architecture Pattern.
// It is used to communicate with the redis.
type RedisPort interface {
	// SetLogger sets logging call-back function
	SetLogger(LogFunc func(ctx context.Context, logData *pb_logging.LogData) (*pb_logging.LoggingResult, error))

	// DeleteUserPasswordByUserId hard-deletes the given user id in the redis.
	DeleteUserPasswordByUserId(ctx context.Context, userId uuid.UUID) error

	// GetUserPasswordByUserId returns active password of the given user.
	GetUserPasswordByUserId(ctx context.Context, userId uuid.UUID) (me.UserPassword, error)

	// ChangePassword changes the poassword of the given user in the redis.
	ChangePassword(ctx context.Context, userPassword me.UserPassword) error
}
