package main

import "net/http"

// USERNAME is...
const USERNAME = "hpazk"

// PASSWORD is...
const PASSWORD = "123#"

// MiddlewareAuth is...
func MiddlewareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			w.Write([]byte(`something went wrong`))
			return
		}
		isValid := (username == USERNAME) && (password == PASSWORD)
		if !isValid {
			w.Write([]byte(`wrong username/password`))
			return
		}
		next.ServeHTTP(w, r)
	})
}

// MiddlewareAllowOnlyGET is...
func MiddlewareAllowOnlyGET(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.Write([]byte("Only GET is allowed"))
			return
		}
		next.ServeHTTP(w, r)
	})
}
