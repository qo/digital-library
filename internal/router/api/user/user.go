package author

import "net/http"

type UserApi interface {
	Get() http.HandlerFunc
	Post() http.HandlerFunc
	Put() http.HandlerFunc
	Delete() http.HandlerFunc
	GetFavoriteBooks() http.HandlerFunc
	GetFavoriteAuthors() http.HandlerFunc
}

type Router interface {
	Get(route string, handler func() http.HandlerFunc)
	Post(route string, handler func() http.HandlerFunc)
	Put(route string, handler func() http.HandlerFunc)
	Delete(route string, handler func() http.HandlerFunc)
}

func Init(r Router, a UserApi) {
	r.Get("/user/{id}", a.Get)
	r.Post("/user", a.Post)
	r.Put("/user", a.Put)
	r.Delete("/user/{id}", a.Delete)
	r.Get("/user/{id}/books", a.GetFavoriteBooks)
	r.Get("/user/{id}/author", a.GetFavoriteAuthors)
}
