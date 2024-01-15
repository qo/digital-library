package views

import (
	"github.com/go-chi/chi/v5"
	"github.com/qo/digital-library/internal/router/views/openapi"
	"github.com/qo/digital-library/internal/router/views/user"
	"github.com/qo/digital-library/internal/storage"
)

type Router struct {
	chi.Router
}

func New(log logger.Logger, st storage.Storage) *Router {
	cr := chi.NewRouter()
	r := Router{cr}
	r.mountRoutes()
	return &r
}

func (r *Router) mountRoutes(log logger.Logger, st storage.Storage) {
	openapi.Init(r)
	user.Init(r)
}
