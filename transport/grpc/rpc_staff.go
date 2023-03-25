package grpctransport

import (
	"context"
	"github.com/thehaung/golang-oracle-example/proto/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *StaffService) ListStaff(context.Context, *emptypb.Empty) (*pb.ListStaffResponse, error) {
	return nil, nil
}

func (s *StaffService) CreateStaff(context.Context, *pb.CreateStaffRequest) (*pb.CreateStaffResponse, error) {
	return nil, nil
}

func (s *StaffService) UpdateStaff(context.Context, *pb.UpdateStaffRequest) (*pb.UpdateStaffResponse, error) {
	return nil, nil
}

func (s *StaffService) FindStaffById(context.Context, *pb.StaffId) (*pb.Staff, error) {
	return nil, nil
}

func (s *StaffService) DeleteStaff(context.Context, *pb.StaffId) (*emptypb.Empty, error) {
	return nil, nil
}
