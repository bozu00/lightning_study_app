package handler

import (
	"github.com/labstack/echo"
	"net/http"
	// "../services"
	// "../formModel"
	// "../models"

    "virtualhost.local/kirakira/lightning_study_app/src/models"
    "virtualhost.local/kirakira/lightning_study_app/src/formModel"
    "virtualhost.local/kirakira/lightning_study_app/src/services"
    // "virtualhost.local/kirakira/lightning_study_app/src/customError"
	"log"
	"os"
	"fmt"
	"strconv"
	//"../customError"
)


func AdminNewArticle(c echo.Context) error {
    return c.Render(http.StatusOK, "newArticle", "this is message")
}



func AdminCreateArticle(c echo.Context) error {
	articleForm := new(formModel.NewArticle)
	if err := c.Bind(articleForm); err != nil {
		log.Println(articleForm)
	}

	articleId, err := models.CreateArticle(*articleForm)
	if err != nil {
		return err
	}

	return c.Redirect(http.StatusMovedPermanently, "/admin/edit_article/" + strconv.FormatInt(articleId, 10))
}


func AdminArticles(c echo.Context) error {
	offset, err := strconv.Atoi(c.Param("offset"))
	if err != nil {
		offset = 0
	}
	limit, err  := strconv.Atoi(c.Param("limit"))
	if err != nil {
		limit = 20
	}

	articleModels, err := models.GetArticles(limit, offset)

	strct := &struct {
		Articles []models.Article
		Msg string
	}{ articleModels, "msg"}

	// return Render(http.StatusOK, "admin/articles.html", strct)
    return c.Render(http.StatusOK, "adminArticles", strct)
}

func AdminEditArticle(c echo.Context) error {
	articleId, err := strconv.Atoi(c.Param("article_id"))
	if err != nil {
		return err
	}

	articleModel, err := models.GetArticle(articleId)

	strct := &struct {
		ArticleModel models.Article
		Msg string
	}{ articleModel, "msg"}

    return c.Render(http.StatusOK, "editArticle", strct)
}


func AdminUpdateArticle(c echo.Context) error {
	// articleId, err := strconv.Atoi(c.Param("article_id"))

	articleId, err := strconv.Atoi(c.FormValue("id"))
	articleModel, err := models.GetArticle(articleId)
	log.Println(err)
	log.Println(articleModel)

	articleForm := new(formModel.UpdateArticle)
	if err := c.Bind(articleForm); err != nil {
		log.Println(articleForm)
	}

	models.UpdateArticle(*articleForm)

	return c.Redirect(http.StatusMovedPermanently, "/admin/edit_article/" + c.FormValue("id"))
}

func AdminFileUploader(c echo.Context) error {
    return c.Render(http.StatusOK, "fileUploader", "")
}


func AdminUploadFile(c echo.Context) error {
	file, err := c.FormFile("image")
	if err != nil {
		return err
	}

	fileName := services.GetFileUploader().FileSave(file)
	id, err := models.CreateImage(fileName)

	return c.Redirect(http.StatusMovedPermanently, "/admin/image/" + fmt.Sprint(id))
}


func AdminImage(c echo.Context) error {
	imageIdString := c.Param("imageId")
	imageId, err := strconv.Atoi(imageIdString)
	if err != nil {
		return err
	}

	//  idからファイルnameを引いてくる
	image, err := models.GetImage(imageId)

	// 自作ORMの実験
	// var image models.Image
	// image := models.NewImage()
	// err = models.Get(&image, imageId)
		
	
	if err != nil {
		return err
	}

	strct := &struct {
		Image models.Image
		Msg string
	}{ image, "msg"}

	// return customError.NoResource
    return c.Render(http.StatusOK, "adminImage", strct)
}

func AdminDeleteImage(c echo.Context) error {
	imageIdString := c.FormValue("imageId")
	imageId, err := strconv.Atoi(imageIdString)
	if err != nil {
		return err
	}

	// idからファイルnameを引いてくる
	image, err := models.GetImage(imageId)
	if err := os.Remove("uploads/image/" + image.Name); err != nil {
    }
	models.DeleteImage(image.Id)

	return c.Redirect(http.StatusMovedPermanently, "/admin/images")
}

func AdminImages(c echo.Context) error {

	images, err := models.GetImages(10, 0)
	if err != nil {
		return err
	}

	strct := &struct {
		Images []models.Image
		Msg string
	}{ images, "msg"}

	// return customError.TestError
    // return c.Render(http.StatusOK, "/templates/admin/images.html", strct)
    return c.Render(http.StatusOK, "adminImages", strct)
}

