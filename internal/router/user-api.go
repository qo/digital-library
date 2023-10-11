package router

import (
	"github.com/qo/digital-library/internal/handlers/api/user"
)

func (r *Router) initUserApiRoutes() {
	r.Get("/user/{id}", user.Get(r.logger, r.storage))
	r.Put("/user", user.Put(r.logger, r.storage))
	r.Delete("/user/{id}", user.Delete(r.logger, r.storage))
}
