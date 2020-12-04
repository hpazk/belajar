package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/save", saveProcess)

	fmt.Println("Server running as http://localhost:9000")
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func saveProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only accept POST method request", http.StatusBadRequest)
		return
	}

	// // Header - Content-Type: application/json
	decoder := json.NewDecoder(r.Body)
	payload := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{}

	if err := decoder.Decode(&payload); err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	message := fmt.Sprintf(
		"your name is %s and your age is %d",
		payload.Name,
		payload.Age,
	)

	w.Write([]byte(message))

}
