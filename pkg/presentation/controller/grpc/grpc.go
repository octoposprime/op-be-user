package presentation

import (
	"net"

	pb_authentication "github.com/octoposprime/op-be-shared/pkg/proto/pb/authentication"
	pb_error "github.com/octoposprime/op-be-shared/pkg/proto/pb/error"
	pb_user "github.com/octoposprime/op-be-shared/pkg/proto/pb/user"
	tgrpc "github.com/octoposprime/op-be-shared/tool/grpc"
	pp_command "github.com/octoposprime/op-be-user/internal/application/presentation/port/command"
	pp_query "github.com/octoposprime/op-be-user/internal/application/presentation/port/query"
	"google.golang.org/grpc"
)

// Grpc is the gRPC API for the application
type Grpc struct {
	pb_error.UnimplementedErorrSvcServer
	pb_user.UnimplementedUserSvcServer
	pb_authentication.UnimplementedAuthenticationSvcServer
	queryHandler   pp_query.QueryPort
	commandHandler pp_command.CommandPort
}

// NewGrpc creates a new instance of Grpc
func NewGrpc(qh pp_query.QueryPort, ch pp_command.CommandPort) *Grpc {
	api := &Grpc{
		queryHandler:   qh,
		commandHandler: ch,
	}
	return api
}

// Serve starts the API server
func (a *Grpc) Serve(port string) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer(
		grpc.UnaryInterceptor(tgrpc.Interceptor),
	)
	pb_error.RegisterErorrSvcServer(s, a)
	pb_user.RegisterUserSvcServer(s, a)
	pb_authentication.RegisterAuthenticationSvcServer(s, a)
	if err := s.Serve(listener); err != nil {
		panic(err)
	}
}
