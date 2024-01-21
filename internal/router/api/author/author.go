package author

import "net/http"

type AuthorApi interface {
	Get() http.HandlerFunc
	Post() http.HandlerFunc
	Put() http.HandlerFunc
	Delete() http.HandlerFunc
}

type Router interface {
	Get(route string, handler http.HandlerFunc)
	Post(route string, handler http.HandlerFunc)
	Put(route string, handler http.HandlerFunc)
	Delete(route string, handler http.HandlerFunc)
}

func Init(r Router, a AuthorApi) {
	r.Get("/author/{id}", a.Get())
	r.Post("/author", a.Post())
	r.Put("/author", a.Put())
	r.Delete("/author/{id}", a.Delete())
}
