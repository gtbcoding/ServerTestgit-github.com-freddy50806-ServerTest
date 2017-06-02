package main

import (
	"bufio"
	"log"
	"mime"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
func SendPic(c echo.Context) error {
	imgFile, err := os.Open("doityourself.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer imgFile.Close()
	fInfo, _ := imgFile.Stat()
	var size int64 = fInfo.Size()
	buf := make([]byte, size)
	fReader := bufio.NewReader(imgFile)
	fReader.Read(buf)
	avatar := []byte(buf)
	return c.Blob(http.StatusOK, mime.TypeByExtension(".jpg"), avatar)
}

func SendMusic(c echo.Context) error {
	return c.File("sock.jpg")
}
func AttachmentFile(c echo.Context) error {
	return c.Attachment("sock.jpg", "sockGG.jpg")
}
func main() {
	e := echo.New()
	//Home page
	e.Static("/", "view/")
	//Picture Send
	e.GET("More1", SendPic)
	e.Static("More2", "doityourself.jpg")
	//
	e.GET("More3", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	//
	e.GET("/users/:id", getUser)
	//Json
	e.GET("/GetJson", GetJson)
	e.GET("/StreamJson", StreamJson)
	e.GET("/Prettyprint", Prettyprint)
	e.POST("/Theory", Theory)

	//Sendfile
	e.GET("/SendMusic", SendMusic)
	e.GET("/SendFile", AttachmentFile)
	e.Logger.Fatal(e.Start(":1323"))

}
