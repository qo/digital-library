package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/qo/digital-library/internal/router/api"
	"github.com/qo/digital-library/internal/router/views"
	"github.com/qo/digital-library/internal/storage"
)

type Router struct {
	chi.Router
}

func New(log logger.Logger, st storage.Storage) *Router {
	cr := chi.NewRouter()
	r := Router{cr}
	r.mountRoutes(log, st)
	return &r
}

func (r Router) mountRoutes(log logger.Logger, st storage.Storage) {
	r.Mount("/api", api.New(log, st))
	r.Mount("/", views.New(log, st))
}
