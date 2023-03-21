package oracle

import (
	"database/sql"
	"errors"
	_ "github.com/sijms/go-ora/v2"
	"github.com/thehaung/golang-oracle-example/config"
	"github.com/thehaung/golang-oracle-example/internal/util/stringutil"
)

const (
	_driverName = "oracle"
)

type oracleClient struct {
	conf *config.Config
	conn *sql.DB
}

type Client interface {
	Connect() error
	Ping() error
	Disconnect() error
	GetConnection() (*sql.DB, error)
}

func NewOracleClient(conf *config.Config) Client {
	return &oracleClient{
		conf: conf,
	}
}

func (o oracleClient) Connect() error {
	conn, err := sql.Open(_driverName, buildConnectionString(o.conf))
	if err != nil {
		return err
	}

	if err = o.Ping(); err != nil {
		return err
	}

	o.conn = conn
	return nil
}

func buildConnectionString(conf *config.Config) string {
	return stringutil.BuildStringWithParams(_driverName, "://", conf.Env.OracleDB.DBUser, ":", conf.Env.OracleDB.DBPass,
		"@", conf.Env.OracleDB.DBHost, "/", conf.Env.OracleDB.DBName)
}

func (o oracleClient) Ping() error {
	return o.conn.Ping()
}

func (o oracleClient) Disconnect() error {
	err := o.conn.Close()
	if err != nil {
		return err
	}

	return nil
}

func (o oracleClient) GetConnection() (*sql.DB, error) {
	if o.conn == nil {
		return nil, errors.New("can't get oracle connection")
	}

	return o.conn, nil
}