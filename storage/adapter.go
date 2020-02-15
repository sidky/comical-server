package storage

import (
	"comical/rest"
	"log"
	"net/http"
)

func Firestore() rest.Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
			_, err := InitStorage()
			if err != nil {
				log.Printf("Unable to initialize Firestore client: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			h.ServeHTTP(w, r)
		})
	}
}
