package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Env struct {
	OracleDB Oracle
}

type Oracle struct {
	DBHost string `mapstrucuture:"ORACLE_DB_HOST"`
	DBPort string `mapstrucuture:"ORACLE_DB_PORT"`
	DBUser string `mapstrucuture:"ORACLE_DB_USER"`
	DBPass string `mapstrucuture:"ORACLE_DB_PASS"`
	DBName string `mapstrucuture:"ORACLE_DB_NAME"`
}

func loadEnv() (*Env, error) {
	var env Env
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("can't find the file .env: %w", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		return nil, fmt.Errorf("environment can't be loaded: %w", err)
	}

	return &env, nil
}
