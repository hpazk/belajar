package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/student", ActionStudent)

	server := new(http.Server)
	server.Addr = ":9000"

	fmt.Println("server started at http://localhost:9000")
	server.ListenAndServe()
}

// ActionStudent is...
func ActionStudent(w http.ResponseWriter, r *http.Request) {
	if !Auth(w, r) {
		return
	}
	if !AllowOnlyGET(w, r) {
		return
	}

	if id := r.URL.Query().Get("id"); id != "" {
		OutputJSON(w, SelectStudent(id))
		return
	}

	OutputJSON(w, GetStudents())
}

// OutputJSON is...
func OutputJSON(w http.ResponseWriter, o interface{}) {
	res, err := json.Marshal(o)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

var students = []*Student{}

// Student is...
type Student struct {
	ID    string
	Name  string
	Grade int32
}

// GetStudents is...
func GetStudents() []*Student {
	return students
}

// SelectStudent is...
func SelectStudent(id string) *Student {
	for _, each := range students {
		if each.ID == id {
			return each
		}
	}

	return nil
}

func init() {
	students = append(students, &Student{ID: "s001", Name: "bourne", Grade: 2})
	students = append(students, &Student{ID: "s002", Name: "ethan", Grade: 2})
	students = append(students, &Student{ID: "s003", Name: "wick", Grade: 3})
}

// USERNAME is...
const USERNAME = "hpazk"

// PASSWORD is...
const PASSWORD = "123#"

// Auth is...
func Auth(w http.ResponseWriter, r *http.Request) bool {
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

// AllowOnlyGET is...
func AllowOnlyGET(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != "GET" {
		w.Write([]byte("Only GET is allowed"))
		return false
	}

	return true
}
