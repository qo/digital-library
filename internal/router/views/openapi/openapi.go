package openapi

import "net/http"

type openapiApi interface {
	Get() http.HandlerFunc
}

type router interface {
	Get(route string, handler func() http.HandlerFunc)
	Post(route string, handler func() http.HandlerFunc)
	Put(route string, handler func() http.HandlerFunc)
	Delete(route string, handler func() http.HandlerFunc)
}

func Init(r router, a openapiApi) {
	r.Get("/openapi", a.Get)
}
