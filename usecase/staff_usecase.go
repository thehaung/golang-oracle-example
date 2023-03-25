package usecase

import (
	"context"
	"github.com/thehaung/golang-oracle-example/domain"
)

type staffUseCase struct {
	staffRepository domain.StaffRepository
}

func NewStaffUseCase(staffRepository domain.StaffRepository) domain.StaffUseCase {
	return &staffUseCase{
		staffRepository: staffRepository,
	}
}

func (s staffUseCase) List(ctx context.Context) ([]domain.Staff, error) {
	return s.staffRepository.List(ctx)
}

func (s staffUseCase) FindById(ctx context.Context, id int64) (domain.Staff, error) {
	return s.staffRepository.FindById(ctx, id)
}

func (s staffUseCase) Create(ctx context.Context, staff domain.Staff) (domain.Staff, error) {
	return s.staffRepository.Create(ctx, staff)
}

func (s staffUseCase) Update(ctx context.Context, staff domain.Staff) (domain.Staff, error) {
	return s.staffRepository.Update(ctx, staff)
}

func (s staffUseCase) Delete(ctx context.Context, id int64) error {
	return s.staffRepository.Delete(ctx, id)
}
