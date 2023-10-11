package router

import (
	"github.com/go-chi/chi/v5"
)

func (r Router) initViewRoutes() {
}

func (r Router) initViewRouter() *Router {
	cr := chi.NewRouter()
	ar := Router{r.logger, r.storage, cr}
	ar.initViewRoutes()
	return &ar
}
