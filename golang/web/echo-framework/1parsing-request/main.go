package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	r := echo.New()

	// parsing query string
	r.GET("/", func(ctx echo.Context) error {
		name := ctx.QueryParam("name")
		data := fmt.Sprintf("hello %s", name)

		return ctx.String(http.StatusOK, data)
	})

	// parsing URL path param
	r.GET("/:name", func(ctx echo.Context) error {
		name := ctx.Param("name")
		data := fmt.Sprintf("hello %s", name)

		return ctx.String(http.StatusOK, data)
	})

	// parsing URL path param and after that
	r.GET("/:name/*", func(ctx echo.Context) error {
		name := ctx.Param("name")
		message := ctx.Param("*")
		data := fmt.Sprintf("hello %s, i have message for you: %s", name, message)

		return ctx.String(http.StatusOK, data)
	})

	// parsing form data
	r.POST("/message", func(ctx echo.Context) error {
		name := ctx.FormValue("name")
		message := ctx.FormValue("message")

		data := fmt.Sprintf("hello %s, %s", name, message)

		return ctx.String(http.StatusOK, data)
	})

	r.Start(":9000")
}
