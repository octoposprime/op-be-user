package application

import (
	"context"

	pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/logging"
)

// CommandPort is a port for Hexagonal Architecture Pattern.
// It is used to communicate with the application layer.
type LoggingCommandPort interface {
	// Log sends the given log to the logging micro service.
	Log(ctx context.Context, logData *pb.LogData) (*pb.LoggingResult, error)
}
