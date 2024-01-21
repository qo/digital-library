package openapi

import "net/http"

type openapiApi interface {
	Get() http.HandlerFunc
}

type router interface {
	Get(route string, handler http.HandlerFunc)
	Post(route string, handler http.HandlerFunc)
	Put(route string, handler http.HandlerFunc)
	Delete(route string, handler http.HandlerFunc)
}

func Init(r router, a openapiApi) {
	r.Get("/openapi", a.Get())
}
