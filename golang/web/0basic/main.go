package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/index/", index)
	http.HandleFunc("/about/", about)

	var address = "localhost:9000"
	fmt.Printf("Server running at %s\n", address)

	// err := http.ListenAndServe(address, nil)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// Memanfaatkan struct http.Server
	server := new(http.Server)
	server.Addr = address
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	var message = "Main Page"
	w.Write([]byte(message))
}

func about(w http.ResponseWriter, r *http.Request) {
	var message = "About Page"
	w.Write([]byte(message))
}
