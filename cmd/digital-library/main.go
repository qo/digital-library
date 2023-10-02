package main

import (
	defaultLog "log"

	"github.com/qo/digital-library/internal/config"
	"github.com/qo/digital-library/internal/logger"
	"github.com/qo/digital-library/internal/storage"
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

	s, err := storage.Init(cfg.StorageOptions)
	if err != nil {
		log.Error(err.Error())
		return
	}

	log.Info("storage loaded")

	user, err := s.GetUser(1)
	if err != nil {
		log.Error(err.Error())
	} else {
		log.Info("user pulled", "user", user)
	}
}
