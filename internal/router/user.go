package router

import (
	"github.com/qo/digital-library/internal/handlers/user"
)

func (r *Router) initUserRoutes() {
	r.Get("/user/{id}", user.Get(r.logger, r.storage))
	r.Put("/user", user.Put(r.logger, r.storage))
	r.Delete("/user/{id}", user.Delete(r.logger, r.storage))
}
