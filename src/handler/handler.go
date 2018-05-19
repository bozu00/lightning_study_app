package handler

import (
	"github.com/labstack/echo"
	"log"
	"errors"
	"net/http"
	// "html/template"
	// "../services"
	// "io"
	// "bytes"
	// "io"
	// "fmt"
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


/*
func Render(w io.Writer, templateName string, data interface{}) error {
	f, err := Assets.Open(templateName)
	buf := bytes.NewBuffer(nil)

    io.Copy(buf, f)

	tplString := fmt.Sprintf("%s", buf)

	// funcMap := template.FuncMap {
	// 	"upper": strings.ToUpper,
	// 	"reverse": e.Reverse,
	// 	"imagePrefix": config.GetInstance().AssetConfig.GetPrefix,
	// 	// "assets" : ""
	// }

	tmpl, err := template.New(templateName).Parse(tplString)

	// Error checking elided
	err = tmpl.Execute(w, data)
	return err
}
*/

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
