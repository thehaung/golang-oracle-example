package grpcserver

import (
	"github.com/thehaung/golang-oracle-example/proto/pb"
	"google.golang.org/grpc"
	"time"
)

const (
	_defaultGrpcPort        = ":8082"
	_defaultShutdownTimeout = 3 * time.Second
)

type GrpcServer struct {
	server *grpc.Server
	pb.UnimplementedStaffServiceServer
}
