package formModel

import (

)

type NewArticle struct {
  Title              string `json:"title" form:"title"` 
  Description        string `json:"description" form:"description"` 
  EyecatchImageId    int    `json:"eyecatch" form:"eyecatch"` 
  TableOfContents    string `json:"toc" form:"toc"`
  Content            string `json:"content" form:"content"`
  IsPublish          string `json:"is_publish" form:"is_publish"`
  // AuthorId           string `json:"title" form:"title"`
}


type UpdateArticle struct {
  Id		         int64  `json:"id" form:"id"`
  Title              string `json:"title" form:"title"` 
  Description        string `json:"description" form:"description"` 
  EyecatchImageId    int    `json:"eyecatch" form:"eyecatch"` 
  TableOfContents    string `json:"toc" form:"toc"`
  Content            string `json:"content" form:"content"`
  IsPublish          string `json:"is_publish" form:"is_publish"`
  // AuthorId           string `json:"title" form:"title"`
}


