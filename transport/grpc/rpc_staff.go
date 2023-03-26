package grpctransport

import (
	"context"
	"github.com/thehaung/golang-oracle-example/domain"
	"github.com/thehaung/golang-oracle-example/internal/util/timeutil"
	"github.com/thehaung/golang-oracle-example/proto/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func (s *StaffService) ListStaff(ctx context.Context, _ *emptypb.Empty) (*pb.ListStaffResponse, error) {
	list, err := s.staffUseCase.List(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to fetch list staff: %s", err)
	}

	if len(list) == 0 {
		return nil, status.Errorf(codes.NotFound, "failed to fetch list staff")
	}

	var staffs []*pb.Staff
	for _, staff := range list {
		staffs = append(staffs, convertStaff(staff))
	}
	return &pb.ListStaffResponse{Staffs: staffs}, nil
}

func (s *StaffService) CreateStaff(ctx context.Context, req *pb.CreateStaffRequest) (*pb.CreateStaffResponse, error) {
	staff := domain.Staff{
		ID:           req.GetId(),
		Name:         req.GetName(),
		TeamName:     req.GetTeamName(),
		Organization: req.GetOrganization(),
		Title:        req.GetTitle(),
		OnboardDate:  req.GetOnboardDate().AsTime(),
	}

	createdStaff, err := s.staffUseCase.Create(ctx, staff)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create staff: %s", err)
	}

	response := &pb.CreateStaffResponse{
		Id:           createdStaff.ID,
		Name:         createdStaff.Name,
		TeamName:     createdStaff.TeamName,
		Organization: createdStaff.Organization,
		Title:        createdStaff.Title,
		OnboardDate:  timestamppb.New(createdStaff.OnboardDate),
		CreatedAt:    timestamppb.New(createdStaff.CreatedAt),
	}
	return response, nil
}

func (s *StaffService) UpdateStaff(ctx context.Context, req *pb.UpdateStaffRequest) (*pb.UpdateStaffResponse, error) {
	return nil, nil
}

func (s *StaffService) FindStaffById(ctx context.Context, req *pb.StaffId) (*pb.Staff, error) {
	defer timeutil.TimeTrack(time.Now())

	staff, err := s.staffUseCase.FindById(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "can't find the staff: %s", err)
	}

	return convertStaff(staff), nil
}

func (s *StaffService) DeleteStaff(ctx context.Context, req *pb.StaffId) (*emptypb.Empty, error) {
	return nil, nil
}
