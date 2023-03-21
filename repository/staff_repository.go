package repository

import (
	"context"
	"database/sql"
	"github.com/thehaung/golang-oracle-example/domain"
	"github.com/thehaung/golang-oracle-example/internal/util/stringutil"
)

const (
	_staffTableName = "STAFF"
)

type staffRepository struct {
	oracleDB *sql.DB
}

func NewStaffRepository(oracleDB *sql.DB) domain.StaffRepository {
	return &staffRepository{
		oracleDB: oracleDB,
	}
}

func (s staffRepository) List(ctx context.Context) ([]domain.Staff, error) {
	//TODO implement me
	panic("implement me")
}

func (s staffRepository) FindById(ctx context.Context, id int64) (domain.Staff, error) {
	findByIdQuery := stringutil.BuildStringWithParams("SELECT ID, NAME, TEAM_NAME, ORGANIZATION, TITLE, ONBOARD_DATE, ACTIVE, ",
		"CREATED_AT, MODIFIED_AT, DELETED_AT FROM ", _staffTableName, " WHERE ID = :1")
	row := s.oracleDB.QueryRowContext(ctx, findByIdQuery, id)

	var staff domain.Staff

	err := row.Scan(
		&staff.ID,
		&staff.Name,
		&staff.TeamName,
		&staff.Organization,
		&staff.Title,
		&staff.OnboardDate,
		&staff.Active,
		&staff.CreatedAt,
		&staff.ModifiedAt,
		&staff.DeletedAt,
	)

	return staff, err
}

func (s staffRepository) Create(ctx context.Context, arg domain.Staff) (domain.Staff, error) {
	var staff domain.Staff
	createQuery := stringutil.BuildStringWithParams("INSERT INTO ", _staffTableName,
		" (ID, NAME, TEAM_NAME, ORGANIZATION, TITLE, ONBOARD_DATE) VALUES(:1, :2, :3, :4, :5, :6, :7)",
		" RETURNING ID, NAME, TEAM_NAME, ORGANIZATION, TITLE, ONBOARD_DATE, ACTIVE, CREATED_AT ",
		" INTO :id, :name, :team_name, :organization, :title, :onboard_date, :active, :created_at")
	row := s.oracleDB.QueryRowContext(ctx, createQuery,
		arg.ID,
		arg.Name,
		arg.TeamName,
		arg.Organization,
		arg.Title,
		arg.OnboardDate,
	)

	err := row.Scan(
		&staff.ID,
		&staff.Name,
		&staff.TeamName,
		&staff.Organization,
		&staff.Title,
		&staff.OnboardDate,
		&staff.Active,
		&staff.CreatedAt,
	)

	return staff, err
}

func (s staffRepository) Update(ctx context.Context, id int64, arg domain.Staff) (domain.Staff, error) {
	//TODO implement me
	panic("implement me")
}

func (s staffRepository) InActive(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}

func (s staffRepository) Active(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}

func (s staffRepository) Delete(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}
