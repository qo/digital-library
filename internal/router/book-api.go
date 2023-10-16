package router

import (
	"github.com/qo/digital-library/internal/handlers/api/book"
)

func (r *Router) initBookApiRoutes() {
	r.Post("/book", book.Post(r.logger, r.storage))
	r.Get("/book/{id}", book.Get(r.logger, r.storage))
	r.Put("/book", book.Put(r.logger, r.storage))
	r.Delete("/book/{id}", book.Delete(r.logger, r.storage))
}
