package router

import (
	"log/slog"

	"github.com/go-chi/chi/v5"
	"github.com/qo/digital-library/internal/handlers/user"
)

type Router struct {
	logger  *slog.Logger
	storage Storage
	chi.Router
}

type Storage interface {
	user.UserStorage
}

func (r Router) initRoutes() {
	r.initUserRoutes()
}

func Init(log *slog.Logger, s Storage) *Router {
	cr := chi.NewRouter()
	r := Router{log, s, cr}
	r.initRoutes()
	return &r
}
