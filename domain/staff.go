package domain

import (
	"context"
	"time"
)

type Staff struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	TeamName     string    `json:"team_name"`
	Organization string    `json:"organization"`
	Title        string    `json:"title"`
	OnboardDate  time.Time `json:"onboard_date"`
	Active       bool      `json:"active"`
	CreatedAt    time.Time `json:"created_at"`
	ModifiedAt   time.Time `json:"modified_at"`
	DeletedAt    time.Time `json:"deleted_at"`
}

type CreateStaffRequest struct {
	ID           int64     `json:"id" validate:"required"`
	Name         string    `json:"name" validate:"required"`
	TeamName     string    `json:"team_name" validate:"required"`
	Organization string    `json:"organization" validate:"required"`
	Title        string    `json:"title" validate:"required"`
	OnboardDate  time.Time `json:"onboard_date" validate:"required"`
}

type CreateStaffResponse struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	TeamName     string    `json:"team_name"`
	Organization string    `json:"organization"`
	Title        string    `json:"title"`
	OnboardDate  time.Time `json:"onboard_date"`
	CreatedAt    time.Time `json:"created_at"`
}

type UpdateStaffRequest struct {
	Name         string    `json:"name"`
	TeamName     string    `json:"team_name"`
	Organization string    `json:"organization"`
	Title        string    `json:"title"`
	OnboardDate  time.Time `json:"onboard_date"`
}

type UpdateStaffResponse struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	TeamName     string    `json:"team_name"`
	Organization string    `json:"organization"`
	Title        string    `json:"title"`
	OnboardDate  time.Time `json:"onboard_date"`
	Active       bool      `json:"active"`
	CreatedAt    time.Time `json:"created_at"`
	ModifiedAt   time.Time `json:"modified_at"`
}

type StaffRepository interface {
	List(ctx context.Context) ([]Staff, error)
	FindById(ctx context.Context, id int64) (Staff, error)
	Create(ctx context.Context, staff Staff) (Staff, error)
	Update(ctx context.Context, id int64, staff Staff) (Staff, error)
	InActive(ctx context.Context, id int64) error
	Active(ctx context.Context, id int64) error
	Delete(ctx context.Context, id int64) error
}

type StaffUseCase interface {
	List(ctx context.Context) ([]Staff, error)
	FindById(ctx context.Context, id int64) (Staff, error)
	Create(ctx context.Context, staff Staff) (Staff, error)
	Update(ctx context.Context, id int64, staff Staff) (Staff, error)
	InActive(ctx context.Context, id int64) error
	Active(ctx context.Context, id int64) error
	Delete(ctx context.Context, id int64) error
}
