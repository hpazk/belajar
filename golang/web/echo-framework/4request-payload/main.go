package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

// User struct is...
type User struct {
	Name  string
	Email string
}

func main() {
	r := echo.New()

	r.Any("/", func(c echo.Context) (err error) {
		u := new(User)
		if err = c.Bind(u); err != nil {
			return
		}

		return c.JSON(http.StatusOK, u)

	})

	fmt.Println("Server start at http://localhost:9000/")
	r.Start(":9000")
}

/*
%Form Data
curl -X POST http://localhost:9000 \
	-d 'name=Hpazk' \
	-d 'email=hpazk@refinerydev.com'

%JSON Payload
curl -X POST http://localhost:9000 \
	-H 'Content-Type: application/json' \
	-d '{"name": "Hpazk", "email": "hpazk@refinerydev.com"}'

%XML Payload
curl -X POST http://localhost:9000 \
	-H 'Content-Type: application/xml' \
	-d '<?xml version=1.0?>
		<Data>\
			<Name>Hpazk</Name>\
			<Email>hpazk@refinerydev.com</Email>/
		</Data>'

%Query String
curl -X GET http://localhost:9000/?name=Hpazk&email=hpazk%40yahoo.com

*/
