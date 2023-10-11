package router

import (
	"github.com/go-chi/chi/v5"
)

func (r Router) initApiRoutes() {
	r.initUserApiRoutes()
}

func (r Router) initApiRouter() *Router {
	cr := chi.NewRouter()
	ar := Router{r.logger, r.storage, cr}
	ar.initApiRoutes()
	return &ar
}
