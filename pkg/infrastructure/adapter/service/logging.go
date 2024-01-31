package infrastructure

import (
	"context"
	"fmt"

	pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/logging"
	tconfig "github.com/octoposprime/op-be-shared/tool/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Log sends the given log to the logging micro service.
func (a ServiceAdapter) Log(ctx context.Context, logData *pb.LogData) (*pb.LoggingResult, error) {
	conn, err := grpc.Dial(tconfig.GetInternalConfigInstance().Grpc.LoggerHost+":"+tconfig.GetInternalConfigInstance().Grpc.LoggerPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(logData.String())
		fmt.Println(err)
		return &pb.LoggingResult{}, err
	}

	pbResult, err := pb.NewLoggingSvcClient(conn).Log(ctx, logData)
	if err != nil {
		fmt.Println(logData.String())
		fmt.Println(err)
		return &pb.LoggingResult{}, err
	}
	return pbResult, nil
}
