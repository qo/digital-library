package router

import (
	"github.com/go-chi/chi/v5"
)

func (r Router) initApiRoutes() {
	r.initSwaggerViewRoutes()
	r.initUserApiRoutes()
	r.initBookApiRoutes()
	r.initAuthorApiRoutes()
}

func (r Router) initApiRouter() *Router {
	cr := chi.NewRouter()
	ar := Router{r.logger, r.storage, cr, r.proto, r.host, r.port}
	ar.initApiRoutes()
	return &ar
}
