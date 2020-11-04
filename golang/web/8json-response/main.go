package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := []struct {
			Name string
			Age  int
		}{
			{"ltg", 25},
			{"hpazk", 26},
			{"zkh", 27},
		}

		jsonInByte, err := json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonInByte)

	})

	fmt.Println("Server start a http://localhost:9000")
	if err := http.ListenAndServe(":9000", nil); err != nil {
		fmt.Println(err.Error())
	}
}
