package application

import (
	"context"

	pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/logging"
)

// Log sends the given log to the logging micro service.
func (a *Service) Log(ctx context.Context, logData *pb.LogData) (*pb.LoggingResult, error) {
	return a.ServicePort.Log(ctx, logData)
}
