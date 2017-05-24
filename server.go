package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Static("/", "view/")
	e.Static("More1", "avatar.png")
	e.Static("More2", "avatar.png")
	e.Logger.Fatal(e.Start(":1323"))

}
