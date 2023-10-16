package router

import (
	"github.com/qo/digital-library/internal/handlers/api/author"
)

func (r *Router) initAuthorApiRoutes() {
	r.Post("/author", author.Post(r.logger, r.storage))
	r.Get("/author/{id}", author.Get(r.logger, r.storage))
	r.Put("/author", author.Put(r.logger, r.storage))
	r.Delete("/author/{id}", author.Delete(r.logger, r.storage))
}
