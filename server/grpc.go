package server

import (
	"fmt"
	"net"

	"github.com/imhshekhar47/go-api-core/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GrpcRegistrarFunc func(grpc.ServiceRegistrar, *any)

var (
	grpcServerLogger = logger.GetLogger("core/server/grpc")
)

type GrpcServer struct {
	port   uint16
	server *grpc.Server
}

func NewGrpcServer(port uint16) *GrpcServer {
	return &GrpcServer{
		port:   port,
		server: grpc.NewServer(),
	}
}

func (s *GrpcServer) Get() *grpc.Server {
	return s.server
}

func (s *GrpcServer) Run() error {
	grpcServerLogger.Debugln("entry: Run()")
	reflection.Register(s.server)

	address := fmt.Sprintf("0.0.0.0:%d", s.port)
	listener, err := net.Listen("tcp", address)

	if err != nil {
		panic(err)
	}

	grpcServerLogger.Debugf("exit: Run(), gRPC listening on %s", listener.Addr())
	return s.server.Serve(listener)
}
