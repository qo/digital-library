package main

import (
	defaultLog "log"

	"github.com/qo/digital-library/internal/config"
	"github.com/qo/digital-library/internal/logger"
)

func main() {
  cfg, err := config.Load()
  if err != nil {
    defaultLog.Fatal(err)
  }
  
  log, err := logger.Init(cfg.Env)
  if err != nil {
    defaultLog.Fatal(err)
  }

  log.Info("config loaded")

  log.Info("config", "cfg", cfg)

  log.Info("logger loaded")

  log.Info("starting server")
} 
