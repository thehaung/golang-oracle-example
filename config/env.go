package config

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Env struct {
	DBHost string `mapstructure:"ORACLE_DB_HOST" validate:"required"`
	DBPort string `mapstructure:"ORACLE_DB_PORT" validate:"required"`
	DBUser string `mapstructure:"ORACLE_DB_USER" validate:"required"`
	DBPass string `mapstructure:"ORACLE_DB_PASS" validate:"required"`
	DBName string `mapstructure:"ORACLE_DB_NAME" validate:"required"`
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

	validate := validator.New()
	if err = validate.Struct(&env); err != nil {
		return nil, fmt.Errorf("required environment can't be loaded: %w", err)
	}

	return &env, nil
}
