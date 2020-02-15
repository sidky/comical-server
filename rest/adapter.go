package rest

import (
	"log"
	"net/http"
)

type Adapter func(handler http.Handler) http.Handler

func Adapt(h http.Handler, adapters... Adapter) http.Handler {
	for _, adapter := range adapters {
		h = adapter(h)
	}
	return h
}

func AdaptFunc(f http.HandlerFunc, adapters... Adapter) http.Handler {
	return Adapt(f, adapters...)
}

// Common adapters
func Log() Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
			log.Printf("Request: %s %s", r.Method, r.URL.String())
			h.ServeHTTP(w, r)
		})
	}
}