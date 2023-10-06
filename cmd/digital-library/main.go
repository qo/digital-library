package main

import (
	"fmt"
	defaultLog "log"
	"net/http"

	"github.com/qo/digital-library/internal/config"
	"github.com/qo/digital-library/internal/logger"
	"github.com/qo/digital-library/internal/router"
	"github.com/qo/digital-library/internal/storage"
)

func serve(host string, port int, handler http.Handler) {
	addr := fmt.Sprintf("%s:%d", host, port)
	http.ListenAndServe(addr, handler)
}

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

	router := router.Init(log, s)

	log.Info("router started")

	log.Info("start serving")

	serve(cfg.HTTPServerOptions.Host, cfg.HTTPServerOptions.Port, router)
}
