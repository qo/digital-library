package router

import (
	"github.com/qo/digital-library/internal/handlers/view/swagger"
)

func (r Router) initSwaggerViewRoutes() {
	r.Get("/", swagger.Get(r.logger))
}
