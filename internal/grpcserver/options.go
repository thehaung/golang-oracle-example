package grpcserver

type Option func(server *GrpcServer)

func Port(port string) Option {
	return func(s *GrpcServer) {
		s.port = port
	}
}
