package grpcserver

import (
	"context"
	"github.com/thehaung/golang-oracle-example/config"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

const (
	_defaultGrpcPort        = ":8082"
	_defaultShutdownTimeout = 3 * time.Second
)

type GrpcServer struct {
	server          *grpc.Server
	notify          chan error
	port            string
	shutdownTimeout time.Duration
}

func NewGrpcServer(conf *config.Config, options ...Option) *GrpcServer {
	grpcServer := grpc.NewServer()
	s := &GrpcServer{
		server:          grpcServer,
		notify:          make(chan error, 1),
		port:            _defaultGrpcPort,
		shutdownTimeout: _defaultShutdownTimeout,
	}

	for _, opt := range options {
		opt(s)
	}

	s.start()

	return s
}

func (s *GrpcServer) start() {
	go func() {
		lis, err := net.Listen("tcp", s.port)
		if err != nil {
			log.Fatalln(err)
			return
		}

		log.Printf("GrpcServer serve at tcp port %s", s.port)
		s.notify <- s.server.Serve(lis)
		close(s.notify)
	}()
}

func (s *GrpcServer) Server() *grpc.Server {
	return s.server
}

// Notify -.
func (s *GrpcServer) Notify() <-chan error {
	return s.notify
}

// Shutdown -.
func (s *GrpcServer) Shutdown() {
	_, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	s.server.GracefulStop()
}
