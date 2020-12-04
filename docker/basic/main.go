package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	index := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(os.Getenv("INDEX")))
	}

	http.HandleFunc("/", index)
	http.HandleFunc("/index/", index)

	http.HandleFunc("/about/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(os.Getenv("ABOUT")))
	})

	// run the server
	fmt.Println("Server running at http://localhost:9000")
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
