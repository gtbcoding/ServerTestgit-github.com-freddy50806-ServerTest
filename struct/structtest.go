package main

import (
	"net/http"

	"fmt"

	"github.com/labstack/echo"
)

type User struct {
	name  string `json:"name" xml:"name" form:"name" query:"name"`
	email string `json:"email" xml:"email" form:"email" query:"email"`
}

func main() {
	e := echo.New()
	e.GET("More3", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/users", func(c echo.Context) error {
		u := new(User)
		if err := c.Bind(&u); err != nil {
			fmt.Printf("error!")
			return err
		}
		return c.JSON(http.StatusCreated, u)
		// or
		// return c.XML(http.StatusCreated, u)
	})

	e.Logger.Fatal(e.Start(":8888"))
}
