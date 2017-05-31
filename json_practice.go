package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type User struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}
type user struct {
	Id    uint   `json:"id"`
	Title string `json:"title"`
}

func Theory(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		fmt.Printf("error!")
		return err
	}
	return c.JSON(http.StatusOK, u)
	// or
	// return c.XML(http.StatusCreated, u)
}

// Handler
func GetJson(c echo.Context) error {
	u := &User{
		Name:  "Jon",
		Email: "jon@labstack.com",
	}
	return c.JSON(http.StatusOK, u)
}

func StreamJson(c echo.Context) error {
	u := &User{
		Name:  "Jon",
		Email: "jon@labstack.com",
	}
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(u)
}
