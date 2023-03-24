package repository

import (
	"context"
	"database/sql"
	"fmt"
	go_ora "github.com/sijms/go-ora/v2"
	"github.com/thehaung/golang-oracle-example/domain"
	"github.com/thehaung/golang-oracle-example/internal/util/stringutil"
	"github.com/thehaung/golang-oracle-example/internal/util/structutil"
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
	//TODO implement me
	panic("implement me")
}

func (s staffRepository) FindById(ctx context.Context, id int64) (domain.Staff, error) {
	//findByIdQuery := stringutil.BuildStringWithParams("SELECT ID, NAME, TEAM_NAME, ORGANIZATION, TITLE, ONBOARD_DATE, ACTIVE, ",
	//	"CREATED_AT, MODIFIED_AT, DELETED_AT FROM ", _staffTableName, " WHERE ID = :1")
	//row := s.oracleDB.QueryRowContext(ctx, findByIdQuery, id)

	var staff domain.Staff

	//err := row.Scan(
	//	&staff.ID,
	//	&staff.Name,
	//	&staff.TeamName,
	//	&staff.Organization,
	//	&staff.Title,
	//	&staff.OnboardDate,
	//	&staff.Active,
	//	&staff.CreatedAt,
	//	&staff.ModifiedAt,
	//	&staff.DeletedAt,
	//)

	rows, err := s.oracleDB.Query("SELECT ID, NAME, TEAM_NAME, ORGANIZATION, TITLE, ONBOARD_DATE FROM " + _staffTableName)
	if err != nil {
		log.Println(err)
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			fmt.Println("Can't close dataset: ", err)
		}
	}()

	log.Println(rows)
	for rows.Next() {
		err = rows.Scan(&staff.ID, &staff.Name, &staff.TeamName, &staff.Organization, &staff.Title, &staff.OnboardDate)
		if err != nil {
			log.Println(err)
		}

		log.Println(structutil.ToJsonString(staff))
	}
	return staff, nil
}

func (s staffRepository) Create(ctx context.Context, arg domain.Staff) (domain.Staff, error) {
	var staff domain.Staff
	createQuery := stringutil.BuildStringWithParams(`INSERT INTO `, _staffTableName,
		` (ID, NAME, TEAM_NAME, ORGANIZATION, TITLE, ONBOARD_DATE) VALUES(:1, :2, :3, :4, :5, :6)`,
		` RETURNING ID, NAME, TEAM_NAME, ORGANIZATION, TITLE, ONBOARD_DATE, CREATED_AT`,
		` INTO :id, :name, :team_name, :organization, :title, :onboard_date, :created_at`)

	stmt, err := s.oracleDB.Prepare(createQuery)
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
