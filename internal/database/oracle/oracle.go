package oracle

import (
	"database/sql"
	"errors"
	_ "github.com/sijms/go-ora/v2"
	go_ora "github.com/sijms/go-ora/v2"
	"github.com/thehaung/golang-oracle-example/config"
	"github.com/thehaung/golang-oracle-example/internal/util/stringutil"
	"log"
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
	Disconnect() error
	GetConnection() (*sql.DB, error)
}

func NewOracleClient(conf *config.Config) Client {
	return &oracleClient{
		conf: conf,
	}
}

func (o *oracleClient) Connect() error {
	urlOptions := map[string]string{
		"FAILOVER":   "10", // the number represent how many times driver will try to reconnect before fail
		"RETRYTIME":  "1000",
		"TRACE FILE": "trace.log",
	}
	databaseUrl := go_ora.BuildUrl("localhost", 1521, o.conf.Env.DBName, o.conf.Env.DBUser, o.conf.Env.DBPass, urlOptions)
	conn, err := sql.Open(_driverName, databaseUrl)

	log.Println(databaseUrl)
	if err != nil {
		return err
	}

	if err = conn.Ping(); err != nil {
		return err
	}

	o.conn = conn
	return nil
}

func buildConnectionString(conf *config.Config) string {
	return stringutil.BuildStringWithParams(_driverName, "://", conf.Env.DBUser, ":", conf.Env.DBPass,
		"@", conf.Env.DBHost, ":", conf.Env.DBPort, "/", conf.Env.DBName, "?FAILOVER=5")
}

func (o *oracleClient) Disconnect() error {
	err := o.conn.Close()
	if err != nil {
		return err
	}

	return nil
}

func (o *oracleClient) GetConnection() (*sql.DB, error) {
	if o.conn == nil {
		return nil, errors.New("can't get oracle connection")
	}

	return o.conn, nil
}
