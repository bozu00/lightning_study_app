package handler

import (
	"github.com/labstack/echo"
	"log"
	"errors"
	"net/http"
	// "../services"
)

func checkErr(err error, msg string) bool {
	if err != nil {
		log.Println(msg, err)
		return false
	}
	return true
}


func checkErrPage(err error, c echo.Context) error {
	if err != nil {
		c.String(http.StatusNotFound, "not found page")
		return errors.New("error")
	}
	return nil
}


// Handler

func HelloWorld(c echo.Context) error {
	// test := services.ParseArticle("hello")
	return c.String(http.StatusOK, "hello")
}

func HelloTemplate(c echo.Context) error {
    return c.Render(http.StatusOK, "hello", "this is message")
}

func LectureArticle(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
