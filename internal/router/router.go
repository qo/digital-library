package router

import (
	"log/slog"

	"github.com/go-chi/chi/v5"

	"github.com/qo/digital-library/internal/handlers/api/user"
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
	r.Mount("/api", r.initApiRouter())
	// TODO: uncomment the live below
	r.Mount("/", r.initViewRouter())
}

func Init(log *slog.Logger, s Storage) *Router {
	cr := chi.NewRouter()
	r := Router{log, s, cr}
	r.initRoutes()
	return &r
}
