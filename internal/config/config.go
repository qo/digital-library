package config

import (
  "fmt"
  "time"
  "github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
  EnvironmentConfig `yaml:"environment"`
  StorageConfig `yaml:"storage"` 
  HTTPServerConfig `yaml:"http_server"`
}

type EnvironmentConfig struct {
  Env string `yaml:"env" env-default="local"`
}

type StorageConfig struct {
  StoragePath string `yaml:"storage_path" env-required:"true"`
}

type HTTPServerConfig struct {
  Host string `yaml:"host" env-default="localhost"`
  Port string `yaml:"port" env-default="5454"`
  Timeout time.Duration `yaml:"timeout" env-default="5s"`
  IdleTimeout time.Duration `yaml:"idle_timeout" env-default="10s"`
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
