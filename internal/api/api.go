package api

import (
	"net/http"
)

// Auth is middleware that calls next on success otherwise redirects to
// /login if token is invalid or missing
func Auth(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		/* hack!
		token := r.URL.Query().Get("token")
		if token == "secret" {
			next.ServeHTTP(w, r)
			return
		}*/
		cookie, err := r.Cookie("token")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		if cookie.Value != "secret" {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		next.ServeHTTP(w, r)
	}
}
