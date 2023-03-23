package config

import (
	"fmt"
	"github.com/spf13/viper"
)

// Config App config struct
type Config struct {
	HttpServer HttpServer
	Logger     Logger
	Env        *Env
}

// HttpServer config
type HttpServer struct {
	Port string
}

// Logger config
type Logger struct {
	Development       bool
	DisableCaller     bool
	DisableStacktrace bool
	Encoding          string
	Level             string
}

func Parse() (*Config, error) {
	var conf Config

	v := viper.New()

	v.SetConfigName("./config/config")
	v.AddConfigPath(".")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, fmt.Errorf("config file not found: %w", err)
		}
		return nil, err
	}

	err := v.Unmarshal(&conf)
	if err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %w", err)
	}

	env, err := loadEnv()
	if err != nil {
		return nil, fmt.Errorf("can't load the environment config, %w", err)
	}

	conf.Env = env

	return &conf, nil
}
