package repository

import (
	"context"
	"database/sql"
	"fmt"
	go_ora "github.com/sijms/go-ora/v2"
	"github.com/thehaung/golang-oracle-example/domain"
	"github.com/thehaung/golang-oracle-example/internal/util/stringutil"
	"log"
)

const (
	_staffTableName = "DB_USER1.STAFF"
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
	rows, err := s.oracleDB.Query("SELECT ID, NAME, TEAM_NAME, ORGANIZATION, TITLE, ONBOARD_DATE, ACTIVE, CREATED_AT, MODIFIED_AT FROM " +
		_staffTableName + " WHERE ACTIVE = 1 ORDER BY CREATED_AT DESC")
	if err != nil {
		return nil, err
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			fmt.Println("Can't close dataset: ", err)
		}
	}()

	var staffs []domain.Staff
	for rows.Next() {
		var staff domain.Staff
		err = rows.Scan(
			&staff.ID,
			&staff.Name,
			&staff.TeamName,
			&staff.Organization,
			&staff.Title,
			&staff.OnboardDate,
			&staff.Active,
			&staff.CreatedAt,
			&staff.ModifiedAt,
		)

		if err != nil {
			log.Println(err)
			continue
		}

		staffs = append(staffs, staff)
	}

	return staffs, err
}

func (s staffRepository) FindById(ctx context.Context, id int64) (domain.Staff, error) {
	findByIdQuery := stringutil.BuildStringWithParams("SELECT ID, NAME, TEAM_NAME, ORGANIZATION, TITLE, ONBOARD_DATE, ACTIVE, ",
		"CREATED_AT, MODIFIED_AT, DELETED_AT FROM ", _staffTableName, " WHERE ID = :1")
	row := s.oracleDB.QueryRowContext(ctx, findByIdQuery, id)

	var staff domain.Staff
	var deletedAt sql.NullTime
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
		&deletedAt,
	)

	if deletedAt.Valid {
		staff.DeletedAt = &deletedAt.Time
	}
	return staff, err
}

func (s staffRepository) Create(ctx context.Context, arg domain.Staff) (domain.Staff, error) {
	var staff domain.Staff
	createQuery := stringutil.BuildStringWithParams(`INSERT INTO `, _staffTableName,
		` (ID, NAME, TEAM_NAME, ORGANIZATION, TITLE, ONBOARD_DATE) VALUES(:1, :2, :3, :4, :5, :6)`,
		` RETURNING ID, NAME, TEAM_NAME, ORGANIZATION, TITLE, ONBOARD_DATE, CREATED_AT`,
		` INTO :id, :name, :team_name, :organization, :title, :onboard_date, :created_at`)

	stmt, err := s.oracleDB.Prepare(createQuery)
	if err != nil {
		return staff, err
	}

	defer func() {
		_ = stmt.Close()
	}()

	_, err = stmt.ExecContext(ctx,
		arg.ID,
		arg.Name,
		arg.TeamName,
		arg.Organization,
		arg.Title,
		arg.OnboardDate,
		sql.Out{Dest: &staff.ID},
		go_ora.Out{Dest: &staff.Name, Size: 100},
		go_ora.Out{Dest: &staff.TeamName, Size: 100},
		go_ora.Out{Dest: &staff.Organization, Size: 100},
		go_ora.Out{Dest: &staff.Title, Size: 100},
		sql.Out{Dest: &staff.OnboardDate},
		sql.Out{Dest: &staff.CreatedAt},
	)

	return staff, err
}

func (s staffRepository) Update(ctx context.Context, arg domain.Staff) (domain.Staff, error) {
	var staff domain.Staff
	updateQuery := stringutil.BuildStringWithParams(`UPDATE `, _staffTableName,
		` SET NAME = :1, TEAM_NAME = :2, ORGANIZATION = :3, TITLE = :4, ONBOARD_DATE = :5, MODIFIED_AT = CURRENT_DATE WHERE id = :6`,
		` RETURNING ID, NAME, TEAM_NAME, ORGANIZATION, TITLE, ONBOARD_DATE, CREATED_AT, MODIFIED_AT`,
		` INTO :id, :name, :team_name, :organization, :title, :onboard_date, :created_at, :modified_at`)

	stmt, err := s.oracleDB.Prepare(updateQuery)
	if err != nil {
		return staff, err
	}

	defer func() {
		_ = stmt.Close()
	}()

	_, err = stmt.ExecContext(ctx,
		arg.Name,
		arg.TeamName,
		arg.Organization,
		arg.Title,
		arg.OnboardDate,
		arg.ID,
		sql.Out{Dest: &staff.ID},
		go_ora.Out{Dest: &staff.Name, Size: 100},
		go_ora.Out{Dest: &staff.TeamName, Size: 100},
		go_ora.Out{Dest: &staff.Organization, Size: 100},
		go_ora.Out{Dest: &staff.Title, Size: 100},
		sql.Out{Dest: &staff.OnboardDate},
		sql.Out{Dest: &staff.CreatedAt},
		sql.Out{Dest: &staff.ModifiedAt},
	)

	return staff, err
}

func (s staffRepository) InActive(ctx context.Context, id int64) error {
	inActiveQuery := stringutil.BuildStringWithParams(`UPDATE `, _staffTableName, ` SET ACTIVE = 0,`,
		` MODIFIED_AT = CURRENT_DATE, DELETED_AT = CURRENT_DATE WHERE :id = 1`)
	stmt, err := s.oracleDB.Prepare(inActiveQuery)
	if err != nil {
		return err
	}

	defer func() {
		_ = stmt.Close()
	}()

	_, err = stmt.ExecContext(ctx, id)
	return err
}

func (s staffRepository) Active(ctx context.Context, id int64) error {
	activeQuery := stringutil.BuildStringWithParams(`UPDATE `, _staffTableName, `SET ACTIVE = 1,`,
		` MODIFIED_AT = CURRENT_DATE WHERE :id = 1`)
	stmt, err := s.oracleDB.Prepare(activeQuery)
	if err != nil {
		return err
	}

	defer func() {
		_ = stmt.Close()
	}()

	_, err = stmt.ExecContext(ctx, id)
	return err
}

func (s staffRepository) Delete(ctx context.Context, id int64) error {
	return s.InActive(ctx, id)
}
