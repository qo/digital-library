package router

import (
	"github.com/qo/digital-library/internal/handlers/view/user"
)

func (r *Router) initUserViewRoutes() {
	r.Get("/user/{id}", user.Get(r.logger, r.proto, r.host, r.port))
}
