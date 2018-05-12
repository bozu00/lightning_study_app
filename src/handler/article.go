package handler

import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"log"
	//"reflect"
	// "../models"
	"../viewModel"
	// "../responses"
	// "html/template"

)


func Article(c echo.Context) error {
	articleId, err := strconv.Atoi(c.Param("article_id"))
	log.Println("article id : " + strconv.Itoa(articleId))
	if err != nil {
		articleId = 1
	}
	article, err := viewModel.GetArticle(articleId)
	
	if errPage := checkErrPage(err, c); errPage != nil {
		// エラーページを返す必要があるなら返す
		log.Println(errPage)
		return errPage
	}

	strct := &struct {
		Article viewModel.ArticleViewModel 
		Msg string
	}{ article, "msg"}
    return c.Render(http.StatusOK, "article", strct)
}

