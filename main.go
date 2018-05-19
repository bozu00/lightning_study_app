package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"html/template"
	"io"
	"net/http"
	"log"
	"reflect"
    "virtualhost.local/kirakira/lightning_study_app/src/config"
    "virtualhost.local/kirakira/lightning_study_app/src/handler"
	"strings"
	"os"
	"strconv"
	"bytes"
	"fmt"
	"path"
	"regexp"

)

type Template struct {
    templates *template.Template
	echo *echo.Echo
}

// func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
//     return t.templates.ExecuteTemplate(w, name, data)
// }

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}

func getAssetFilePaths(pattern string) []string {
	dirs := TemplateAssets.Dirs
	paths := []string{}
	for dir, files := range dirs {
		for _, file := range files {
			paths = append(paths, path.Join(dir, file))
		}
	}

	res := []string{}
	// pattern := ".*.html"
	for _, file := range paths {
		bl, err := regexp.MatchString(pattern, file)
		if err != nil {
			log.Println(err.Error())
		}
		if bl {
			res = append(res, file)
		} 	
	}
	return res
}

// go:generate go-assets-builder -o templateAssets.go -p main -v TemplateAssets assets/
// go:generate go-assets-builder -o publicAssets.go -p main -v PublicAssets templates/
func main() {
	e := echo.New()

	if len(os.Args) != 2 {
		log.Println("指定された引数の数が間違っています。")
		os.Exit(1)
	}

	if err := config.InitConfig(os.Args[1]); err != nil {
		log.Printf("fail config %s", err.Error())
		panic("fail config")
		os.Exit(1)
	}

    funcMap := template.FuncMap {
        "upper": strings.ToUpper,
		"reverse": e.Reverse,
		"imagePrefix": config.GetInstance().AssetConfig.GetPrefix,
		// "assets" : ""
    }

	htmlFiles := getAssetFilePaths(".*.html")
	var temp = template.New("templates")
	for _, file := range htmlFiles {

		f, err := TemplateAssets.Open(file)
		if err != nil {
			log.Println(err)
		}

		buf := bytes.NewBuffer(nil)
		_, err = io.Copy(buf, f)
		if err != nil {
			log.Println(err)
		}

		tplString := fmt.Sprintf("%s", buf)

		temp.Funcs(funcMap).Parse(tplString)
		// temp, err = template.New("templates").Funcs(funcMap).Parse(tplString)
	}

	t := &Template{
		// templates: template.Must(template.New("templates").Funcs(funcMap).ParseGlob("templates/**/**.html")),
		templates: temp,
		echo: e,
	}
    e.Renderer = t

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())


	e.HTTPErrorHandler = func(err error, c echo.Context) {
		log.Println(reflect.TypeOf(err))
		log.Println(err.Error())
		if he, ok := err.(*echo.HTTPError); ok {
			if he.Code == 404 {
				c.Render(http.StatusNotFound,"404",nil)
			} else {
				c.Render(http.StatusInternalServerError,"500",nil)
			}
		}

		if err != nil {
				c.Render(http.StatusInternalServerError,"500",nil)
		}
	}


	// e.Static("/assets", "assets") 

	// e.GET("/assetss", http.FileServer(http.Dir("./assets"))) 

	// e.Use(middkleware.Static("/assets"))

	// e.Static("/images/slides", "upload_data")
	e.Static("/uploads/image", "uploads/image")


	// e.Static("/assets", "assets") 
	// fs := http.FileServer(http.Dir("/assets"))
	fs := http.FileServer(PublicAssets)
	e.GET("/assets/*", echo.WrapHandler(http.StripPrefix("/assets", fs))) 

	e.GET("/hello_world",         handler.HelloWorld)
	e.GET("/hello_template",      handler.HelloTemplate)

	e.GET("/article/:article_id", handler.Article).Name = "Article"


	g := e.Group("/admin")
	// g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
	// 	if username == "joe" && password == "secret" {
	// 		return true, nil
	// 	}
	// 	return false, nil
	// }))

	g.GET( "/articles/",                handler.AdminArticles).Name      = "AdminArticles"
	g.GET( "/new_article/",             handler.AdminNewArticle).Name    = "AdminArticleNew"
	g.POST("/create_article/",          handler.AdminCreateArticle).Name = "AdminArticleCreate"
	g.GET( "/edit_article/:article_id", handler.AdminEditArticle).Name   = "AdminArticle"
	g.POST("/update_article/",          handler.AdminUpdateArticle).Name = "AdminUpdate"

	// image
	g.GET( "/images",                   handler.AdminImages).Name       = "AdminImages"
	g.GET( "/image/:imageId",           handler.AdminImage).Name        = "AdminImage"
	g.GET( "/new_image/",               handler.AdminFileUploader).Name = "AdminImageNew"
	g.POST("/create_image/",            handler.AdminUploadFile).Name   = "AdminImageCreate"
	g.POST("/delete_image/",            handler.AdminDeleteImage).Name  = "AdminImageDelete"


    getAssetFilePaths(".*.html")

	port := config.GetInstance().APIConfig.Port
	e.Logger.Fatal(e.Start(":" + strconv.Itoa(port)))
	// e.Logger.Fatal(e.Start(":1323"))
}
