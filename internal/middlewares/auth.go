package middlewares

import (
	"net/http"

	"github.com/Kwynto/gosession"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session := gosession.Start(&w, r)

		if session.Get("is_login") == nil {
			http.Redirect(w, r, "/", 303)
		}

		next.ServeHTTP(w, r)
	})
}

func GuestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session := gosession.Start(&w, r)

		if session.Get("is_login") == true {
			http.Redirect(w, r, "/hello", 303)
		}

		next.ServeHTTP(w, r)
	})
}
