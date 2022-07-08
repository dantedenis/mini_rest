package web

import (
	"log"
	"net/http"
)

func midddleLogger(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request: Method " + r.Method + ", url " + r.RequestURI)
		next.ServeHTTP(w,r)
	}
}