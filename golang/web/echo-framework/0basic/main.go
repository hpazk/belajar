package main

import (
	"net/http"

	"github.com/labstack/echo"
)

// M is...
type M map[string]interface{}

func main() {
	r := echo.New()

	r.GET("/", func(ctx echo.Context) error {
		data := "Index Page"
		return ctx.String(http.StatusOK, data)
	})

	r.GET("/html", func(ctx echo.Context) error {
		data := "<h1 style=\"color: red\">Index Page</h1>"
		return ctx.HTML(http.StatusOK, data)
	})

	r.GET("/json", func(ctx echo.Context) error {
		data := M{"message": "Success"}
		return ctx.JSON(http.StatusOK, data)
	})

	r.GET("/redirect", func(ctx echo.Context) error {
		return ctx.Redirect(http.StatusTemporaryRedirect, "/")
	})

	r.Start(":9000")
}
