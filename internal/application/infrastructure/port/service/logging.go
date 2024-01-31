package application

import (
	"context"

	pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/logging"
)

// LoggingServicePort is a port for Hexagonal Architecture Pattern.
// It is used to communicate with the other servies.
type LoggingServicePort interface {
	// Log sends the given log to the logging micro service.
	Log(ctx context.Context, logData *pb.LogData) (*pb.LoggingResult, error)
}
