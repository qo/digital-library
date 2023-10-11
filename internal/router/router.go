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
	proto string
	host  string
	port  int
}

type Storage interface {
	user.UserStorage
}

func (r Router) initRoutes() {
	r.Mount("/api", r.initApiRouter())
	r.Mount("/", r.initViewRouter())
}

func Init(log *slog.Logger, s Storage, proto string, host string, port int) *Router {
	cr := chi.NewRouter()
	r := Router{log, s, cr, proto, host, port}
	r.initRoutes()
	return &r
}
