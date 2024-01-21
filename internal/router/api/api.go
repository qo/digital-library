package api

import (
	"github.com/go-chi/chi/v5"
	author_handler "github.com/qo/digital-library/internal/handlers/api/author"
	book_handler "github.com/qo/digital-library/internal/handlers/api/book"
	user_handler "github.com/qo/digital-library/internal/handlers/api/user"
	"github.com/qo/digital-library/internal/logger"
	author_router "github.com/qo/digital-library/internal/router/api/author"
	book_router "github.com/qo/digital-library/internal/router/api/book"
	user_router "github.com/qo/digital-library/internal/router/api/user"
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
	ah := author_handler.New(log, st)
	bh := book_handler.New(log, st)
	uh := user_handler.New(log, st)

	author_router.Init(r, ah)
	book_router.Init(r, bh)
	user_router.Init(r, uh)
}
