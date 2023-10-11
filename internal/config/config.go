package config

import (
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	EnvironmentOptions `yaml:"environment"`
	StorageOptions     `yaml:"storage"`
	HTTPServerOptions  `yaml:"http_server"`
}

type EnvironmentOptions struct {
	Env string `yaml:"env" env-default:"local"`
}

type StorageOptions struct {
	Db            string `yaml:"db"             env-required:"true"`
	MySQLOptions  `       yaml:"mysql_options"`
	SQLiteOptions `       yaml:"sqlite_options"`
}

type MySQLOptions struct {
	Name            string        `yaml:"mysql_name"`
	User            string        `yaml:"mysql_user"`
	Password        string        `yaml:"mysql_password"`
	MaxConnLifetime time.Duration `yaml:"mysql_max_conn_lifetime"`
	MaxOpenConns    int           `yaml:"mysql_max_open_conns"`
	MaxIdleConns    int           `yaml:"mysql_max_idle_conns"`
}

type SQLiteOptions struct {
	Path string `yaml:"sqlite_path"`
}

type HTTPServerOptions struct {
	Proto       string        `yaml:"proto"        env-default:"https"`
	Host        string        `yaml:"host"         env-default:"localhost"`
	Port        int           `yaml:"port"         env-default:"5454"`
	Timeout     time.Duration `yaml:"timeout"      env-default:"5s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"10s"`
}

func Load() (*Config, error) {
	const errMsg = "can't load config"

	path, err := getPath()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errMsg, err)
	}

	ex := exists(path)
	if !ex {
		return nil, fmt.Errorf("%s: file specified in config path doesn't exist", errMsg)
	}

	var cfg Config

	err = cleanenv.ReadConfig(path, &cfg)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errMsg, err)
	}

	return &cfg, nil
}
