package handler

import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	// "log"

    "virtualhost.local/kirakira/lightning_study_app/src/viewModel"
    "virtualhost.local/kirakira/lightning_study_app/src/customError"
    "github.com/Shaked/gomobiledetect"
)


func Article(c echo.Context) error {
	detect := mobiledetect.NewMobileDetect(c.Request(), nil)

    if	detect.IsMobile() || detect.IsTablet() {
		articleId, err := strconv.Atoi(c.Param("article_id"))
		if err != nil {
			return customError.NoResource
		}

		article, err := viewModel.GetArticle(articleId)
		if err != nil {
			return customError.NoResource
		}

		strct := &struct {
			Article viewModel.ArticleViewModel 
			Msg string
		}{ article, "msg"}
		return c.Render(http.StatusOK, "article", strct)
	} else {
		return c.String(http.StatusOK, "hello")
	}
}

func ArticleApp(c echo.Context) error {
	detect := mobiledetect.NewMobileDetect(c.Request(), nil)

    if	detect.IsMobile() || detect.IsTablet() {
		articleId, err := strconv.Atoi(c.Param("article_id"))
		if err != nil {
			return customError.NoResource
		}

		article, err := viewModel.GetArticle(articleId)
		if err != nil {
			return customError.NoResource
		}

		strct := &struct {
			Article viewModel.ArticleViewModel 
			Msg string
		}{ article, "msg"}
		return c.Render(http.StatusOK, "article", strct)
	} else {
		return c.String(http.StatusOK, "hello")
	}
}

