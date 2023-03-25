package grpctransport

import (
	"github.com/thehaung/golang-oracle-example/domain"
	"github.com/thehaung/golang-oracle-example/proto/pb"
)

type Server struct {
	staffUseCase domain.StaffUseCase
	pb.UnimplementedStaffServiceServer
}

func NewTransport(useCase domain.StaffUseCase) *Server {
	return &Server{
		staffUseCase: useCase,
	}
}
