package book

import "net/http"

type BookApi interface {
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

func Init(r Router, a BookApi) {
	r.Get("/book/{id}", a.Get)
	r.Post("/book", a.Post)
	r.Put("/book", a.Put)
	r.Delete("/book/{id}", a.Delete)
}
