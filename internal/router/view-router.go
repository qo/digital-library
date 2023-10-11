package router

import (
	"github.com/go-chi/chi/v5"
)

func (r Router) initViewRoutes() {
	r.initUserViewRoutes()
}

func (r Router) initViewRouter() *Router {
	cr := chi.NewRouter()
	ar := Router{r.logger, r.storage, cr, r.proto, r.host, r.port}
	ar.initViewRoutes()
	return &ar
}
