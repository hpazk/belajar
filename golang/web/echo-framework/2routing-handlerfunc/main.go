package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	var index = echo.WrapHandler(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("Index Page"))
			},
		),
	)

	r := echo.New()

	r.GET("/", index)
	r.GET("/about", echo.WrapHandler(http.HandlerFunc(AboutHandler)))

	r.Start(":9000")
}

// AboutHandler is...
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About Page"))
}
