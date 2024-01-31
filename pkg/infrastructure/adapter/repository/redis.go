package infrastructure

import (
	"context"

	"github.com/google/uuid"
	smodel "github.com/octoposprime/op-be-shared/pkg/model"
	pb_logging "github.com/octoposprime/op-be-shared/pkg/proto/pb/logging"
	tredis "github.com/octoposprime/op-be-shared/tool/redis"
	me "github.com/octoposprime/op-be-user/internal/domain/model/entity"
	mo "github.com/octoposprime/op-be-user/internal/domain/model/object"
	map_repo "github.com/octoposprime/op-be-user/pkg/infrastructure/mapper/repository"
)

type RedisAdapter struct {
	*tredis.RedisClient
	Log func(ctx context.Context, logData *pb_logging.LogData) (*pb_logging.LoggingResult, error)
}

func NewRedisAdapter(redisClient *tredis.RedisClient) RedisAdapter {
	adapter := RedisAdapter{
		redisClient,
		Log,
	}

	return adapter
}

// SetLogger sets logging call-back function
func (a *RedisAdapter) SetLogger(LoggerFunc func(ctx context.Context, logData *pb_logging.LogData) (*pb_logging.LoggingResult, error)) {
	a.Log = LoggerFunc
}

// DeleteUserPasswordByUserId hard-deletes the given user id in the redis.
func (a RedisAdapter) DeleteUserPasswordByUserId(ctx context.Context, userId uuid.UUID) error {
	err := a.RedisClient.DeleteHKey(ctx, "USERPASSWORD", userId.String())
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "DeleteUserPasswordByUserId", userId, err.Error()))
		return err
	}
	return nil
}

// GetUserPasswordByUserId returns active password of the given user.
func (a RedisAdapter) GetUserPasswordByUserId(ctx context.Context, userId uuid.UUID) (me.UserPassword, error) {
	var userPasswordsDbMapper map_repo.UserPassword
	data, err := a.RedisClient.ReadHKey(ctx, "USERPASSWORD", userId.String(), userPasswordsDbMapper)
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "GetUsersByFilter", userId, err.Error()))
		return me.UserPassword{}, err
	}
	if result, ok := data.(map_repo.UserPassword); ok {
		return *result.ToEntity(), nil
	} else {
		return me.UserPassword{}, mo.ErrorUserPasswordNotFound
	}
}

// ChangePassword changes the given user password in the redis.
func (a RedisAdapter) ChangePassword(ctx context.Context, userPassword me.UserPassword) error {
	err := a.RedisClient.WriteHKey(ctx, "USERPASSWORD", userPassword.UserId.String(), userPassword)
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "ChangePassword", userId, err.Error()))
		return err
	}
	return nil
}
