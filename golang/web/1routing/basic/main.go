package main

import (
	"fmt"
	"net/http"
)

func main() {
	index := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Index Page"))
	}

	http.HandleFunc("/", index)
	http.HandleFunc("/index/", index)

	http.HandleFunc("/about/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("About Page"))
	})

	// run the server
	fmt.Println("Server running at localhost:9000")
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println(err.Error())
	}

}
