package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", getData)
	http.HandleFunc("/submit", submitData)

	// run the server
	fmt.Println("Server running at localhost:9000")
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func getData(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var tmpl = template.Must(template.New("form").ParseFiles("view.html"))
		var err = tmpl.Execute(w, nil)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	http.Error(w, "", http.StatusBadRequest)
}

func submitData(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var tmpl = template.Must(template.New("result").ParseFiles("view.html"))

		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		// Header - Content-Type: application/x-www-form-urlencoded
		var name = r.FormValue("name")
		var message = r.FormValue("message")

		var data = map[string]string{"name": name, "message": message}

		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, "", http.StatusInternalServerError)
		}
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}
