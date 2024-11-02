package main

import (
	"net/http"
)

// AdminHandler untuk halaman admin
func AdminHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to Admin page"))
	}
}

// RequestMethodGetMiddleware memastikan hanya metode GET yang diizinkan
func RequestMethodGetMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method is not allowed"))
			return
		}
		next.ServeHTTP(w, r)
	})
}

// AdminMiddleware memastikan bahwa hanya pengguna dengan role ADMIN yang diizinkan
func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role := r.Header.Get("role")
		if role != "ADMIN" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Role not authorized"))
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	// Menetapkan chaining middleware pada endpoint /admin
	http.Handle("/admin", RequestMethodGetMiddleware(AdminMiddleware(AdminHandler())))

	http.ListenAndServe("localhost:8080", nil)
}
