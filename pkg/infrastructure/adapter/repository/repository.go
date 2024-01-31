package infrastructure

import (
	"context"

	pb_logging "github.com/octoposprime/op-be-shared/pkg/proto/pb/logging"
)

// Log is the default log function
func Log(ctx context.Context, logData *pb_logging.LogData) (*pb_logging.LoggingResult, error) {
	return &pb_logging.LoggingResult{}, nil
}
