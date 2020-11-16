package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	http.HandleFunc("/", handleIndex)

	fmt.Println("Server running at http://localhost:9000")
	if err := http.ListenAndServe(":9000", nil); err != nil {
		fmt.Println(err.Error())
	}
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	done := make(chan bool)
	go func() {
		// do the process here
		// simulate a long-time request by putting 10 seconds sleep
		time.Sleep(10 * time.Second)

		done <- true
	}()

	select {
	case <-r.Context().Done():
		if err := r.Context().Err(); err != nil {
			if strings.Contains(strings.ToLower(err.Error()), "canceled") {
				log.Println("request canceled")
			} else {
				log.Println("unknown error occured.", err.Error())
			}
		}
	case <-done:
		log.Println("done")
	}
}

// With Request Payload
// go func() {
//     // do the process here
//     // simulate a long-time request by putting 10 seconds sleep
//     body, err := ioutil.ReadAll(r.Body)
// // ...
//     time.Sleep(10 * time.Second)
//     done <- true
// }()
