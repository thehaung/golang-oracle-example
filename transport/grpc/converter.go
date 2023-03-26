package grpctransport

import (
	"github.com/thehaung/golang-oracle-example/domain"
	"github.com/thehaung/golang-oracle-example/proto/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertStaff(staff domain.Staff) *pb.Staff {
	return &pb.Staff{
		Id:           staff.ID,
		Name:         staff.Name,
		TeamName:     staff.TeamName,
		Organization: staff.Organization,
		Title:        staff.Title,
		OnboardDate:  timestamppb.New(staff.OnboardDate),
		Active:       staff.Active,
		CreatedAt:    timestamppb.New(staff.CreatedAt),
		ModifiedAt:   timestamppb.New(staff.ModifiedAt),
		DeletedAt:    timestamppb.New(staff.DeletedAt),
	}
}
