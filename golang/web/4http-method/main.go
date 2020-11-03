package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			w.Write([]byte("post"))
		case "GET":
			w.Write([]byte("get"))
		case "PUT":
			w.Write([]byte("put"))
		case "DELETE":
			w.Write([]byte("delete"))
		default:
			http.Error(w, "", http.StatusBadRequest)
		}
	})

	fmt.Println("servert running at http://localhost:9000/")
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
