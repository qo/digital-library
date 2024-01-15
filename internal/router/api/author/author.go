package author

import "net/http"

type AuthorApi interface {
	Get() http.HandlerFunc
	Post() http.HandlerFunc
	Put() http.HandlerFunc
	Delete() http.HandlerFunc
}

type Router interface {
	Get(route string, handler func() http.HandlerFunc)
	Post(route string, handler func() http.HandlerFunc)
	Put(route string, handler func() http.HandlerFunc)
	Delete(route string, handler func() http.HandlerFunc)
}

func Init(r Router, a AuthorApi) {
	r.Get("/author/{id}", a.Get)
	r.Post("/author", a.Post)
	r.Put("/author", a.Put)
	r.Delete("/author/{id}", a.Delete)
}
