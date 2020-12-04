package main

import "net/http"

// USERNAME is...
const USERNAME = "hpazk"

// PASSWORD is...
const PASSWORD = "123#"

// MiddlewareAuth is...
func MiddlewareAuth(w http.ResponseWriter, r *http.Request) bool {
	username, password, ok := r.BasicAuth()
	if !ok {
		w.Write([]byte(`something went wrong`))
		return false
	}

	isValid := (username == USERNAME) && (password == PASSWORD)
	if !isValid {
		w.Write([]byte(`wrong username/password`))
		return false
	}

	return true
}

// MiddlewareAllowOnlyGET is...
func MiddlewareAllowOnlyGET(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != "GET" {
		w.Write([]byte("Only GET is allowed"))
		return false
	}

	return true
}
