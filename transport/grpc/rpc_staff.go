package grpctransport

import (
	"context"
	"github.com/thehaung/golang-oracle-example/proto/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) ListStaff(context.Context, *emptypb.Empty) (*pb.ListStaffResponse, error) {
	return nil, nil
}

func (s *Server) CreateStaff(context.Context, *pb.CreateStaffRequest) (*pb.CreateStaffResponse, error) {
	return nil, nil
}

func (s *Server) UpdateStaff(context.Context, *pb.UpdateStaffRequest) (*pb.UpdateStaffResponse, error) {
	return nil, nil
}

func (s *Server) FindStaffById(context.Context, *pb.StaffId) (*pb.Staff, error) {
	return nil, nil
}

func (s *Server) DeleteStaff(context.Context, *pb.StaffId) (*emptypb.Empty, error) {
	return nil, nil
}
