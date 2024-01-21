package views

import (
	"github.com/go-chi/chi/v5"
	openapi_handler "github.com/qo/digital-library/internal/handlers/view/openapi"
	user_handler "github.com/qo/digital-library/internal/handlers/view/user"
	"github.com/qo/digital-library/internal/logger"
	openapi_router "github.com/qo/digital-library/internal/router/views/openapi"
	user_router "github.com/qo/digital-library/internal/router/views/user"
	"github.com/qo/digital-library/internal/storage"
)

type Router struct {
	chi.Router
}

func New(log logger.Logger, st storage.Storage) *Router {
	cr := chi.NewRouter()
	r := Router{cr}
	r.mountRoutes(log, st)
	return &r
}

func (r *Router) mountRoutes(log logger.Logger, st storage.Storage) {
	oh := openapi_handler.New(log)
	uh := user_handler.New(log, st)

	openapi_router.Init(r, oh)
	user_router.Init(r, uh)
}
