package grpctransport

import (
	"github.com/thehaung/golang-oracle-example/domain"
	"github.com/thehaung/golang-oracle-example/proto/pb"
)

type StaffService struct {
	staffUseCase domain.StaffUseCase
	pb.UnimplementedStaffServiceServer
}

func NewStaffService(useCase domain.StaffUseCase) *StaffService {
	return &StaffService{
		staffUseCase: useCase,
	}
}
