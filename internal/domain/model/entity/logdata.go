package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	smodel "github.com/octoposprime/op-be-shared/pkg/model"
	pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/logging"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// LogData is a struct that represents the entity of a log.
type LogData struct {
	Id            uuid.UUID `json:"id"` // Id is the id of the log.
	*pb.LogHeader           // LogHeader is the header of the log.
	*pb.LogBody             // LogBody is the body of the log.
}

// NewLogData creates a new *LogData.
func NewLogData() *LogData {
	return &LogData{
		Id:        uuid.UUID{},
		LogHeader: &pb.LogHeader{},
		LogBody:   &pb.LogBody{},
	}
}

// GenerateLogData generated a new *pb.LogData
func (s *LogData) GenerateLogData(logType pb.LogType,
	path string,
	userId string,
	message string) *pb.LogData {

	return &pb.LogData{
		Id: uuid.UUID{}.String(),
		Header: &pb.LogHeader{
			EventDate:   timestamppb.New(time.Now()),
			LogType:     logType,
			ServiceName: smodel.ServiceUser,
			Path:        path,
			UserId:      userId,
		},
		Body: &pb.LogBody{
			Message: message,
		},
	}
}

// String returns a string representation of the LogData.
func (s *LogData) String() string {
	return fmt.Sprintf("Id: %v, "+
		"LogHeader: %v, "+
		"LogBody: %v",
		s.Id,
		s.LogHeader,
		s.LogBody)
}
