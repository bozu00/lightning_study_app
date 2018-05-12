package viewModel

import (
	"../models"
	"../services"
	"html/template"
	"strings"
	// "fmt"
	// "log"
)

type ArticleViewModel struct {
	Id                int64  
	Title             string 
	EyecatchImagePath string 
	Description		  template.HTML
	TableOfContents   []string
	Content           template.HTML
	AuthorId          int64  
	CreatedAt         string 
	UpdatedAt         string 
}

func GetArticle(article_id int) (ArticleViewModel, error) {
	articleModel, err := models.GetArticle(article_id)
	article := makeArticleViewModel(articleModel)

	return article, err
}


func makeArticleViewModel(model models.Article) ArticleViewModel{
	content := template.HTML(services.Parse(model.Content))
	description := template.HTML(model.Description)
	tableOfContents := strings.Split(model.TableOfContents, ",")
	eyecatchImage, err := models.GetImage(model.EyeCatchImageID)
	eyecatchImagePath := eyecatchImage.Name
	if err != nil {
		eyecatchImagePath = ""
	}

	
	article := ArticleViewModel {
		Id: model.Id, 
		Title: model.Title,
		Description: description,
		EyecatchImagePath: eyecatchImagePath,
		TableOfContents: tableOfContents,
		Content: content,
		AuthorId: model.AuthorId,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt}

	return article
}
