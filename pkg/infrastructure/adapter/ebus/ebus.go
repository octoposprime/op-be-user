package infrastructure

import (
	"context"
	"fmt"
	"time"

	pb_logging "github.com/octoposprime/op-be-shared/pkg/proto/pb/logging"
	pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/user"
	tredis "github.com/octoposprime/op-be-shared/tool/redis"
	tserialize "github.com/octoposprime/op-be-shared/tool/serialize"
	me "github.com/octoposprime/op-be-user/internal/domain/model/entity"
	map_ebus "github.com/octoposprime/op-be-user/pkg/infrastructure/mapper/ebus"
)

type EBusAdapter struct {
	redisClient *tredis.RedisClient
	Log         func(ctx context.Context, logData *pb_logging.LogData) (*pb_logging.LoggingResult, error)
}

func NewEBusAdapter(redisClient *tredis.RedisClient) EBusAdapter {
	adapter := EBusAdapter{
		redisClient: redisClient,
		Log:         Log,
	}
	return adapter
}

// SetLogger sets logging call-back function
func (a *EBusAdapter) SetLogger(LoggerFunc func(ctx context.Context, logData *pb_logging.LogData) (*pb_logging.LoggingResult, error)) {
	a.Log = LoggerFunc
}

// Log is the default log function
func Log(ctx context.Context, logData *pb_logging.LogData) (*pb_logging.LoggingResult, error) {
	fmt.Println(logData)
	return &pb_logging.LoggingResult{}, nil
}

// Listen listens to the redis messaging queue and calls the given callBack function for each received user.
func (a EBusAdapter) Listen(ctx context.Context, channelName string, callBack func(channelName string, user me.User)) {
	for {
		result, err := a.redisClient.BLPop(ctx, 0*time.Second, channelName).Result()
		if err != nil {
			continue
		}
		inChannelName := result[0]
		user := tserialize.NewSerializer(me.User{}).FormJson(result[1]).(*pb.User)
		go callBack(inChannelName, *map_ebus.NewUser(user).ToEntity())
	}
}
