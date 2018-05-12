package services

import (
	// "os"

	// "strings"
	// "fmt"
	"github.com/alecthomas/participle"
	// "github.com/alecthomas/participle/lexer"
	// "github.com/alecthomas/repr"
	"text/template"
	"bytes"

)

type Article struct {
	Sections []*Section `{@@}`
}
func (self *Article) ToHtml() string {
	templateName := "templates/articleParser/article.tmpl"
	tmpl, err := template.ParseFiles(templateName)
    if err != nil {
        panic(err)
    }

	var doc bytes.Buffer
    if err := tmpl.Execute(&doc, self); err != nil {
		panic(err)
    }
	return doc.String()
}

type Section struct {
	// Sections []*Section `{ @@ }`
	Header string `"[" @Ident "]"`
	CustomTags []*CustomTag `{ @@ }`
}

func (self *Section) ToHtml() string {
	templateName := "templates/articleParser/section.tmpl"
	tmpl, err := template.ParseFiles(templateName)
    if err != nil {
        panic(err)
    }

	var doc bytes.Buffer
    if err := tmpl.Execute(&doc, self); err != nil {
		panic(err)
    }
	return doc.String()
}
// type Section struct {
// 	// Header string `["[" @Ident "]"]`
// 	CustomTags []*CustomTag `{ @@ }`
// }


/*
func (self *Section) ToHtml() string {
	templateName := "templates/articleParser/section.tmpl"
	tmpl, err := template.ParseFiles(templateName)
    if err != nil {
        panic(err)
    }

	var doc bytes.Buffer
    if err := tmpl.Execute(&doc, self); err != nil {
		panic(err)
    }
	return doc.String()
}
*/

type CustomTag struct {
	Type string `@Ident`
	Value string `"{" @String "}"`
	// StartBrace string `"{"`
	// Value string `@String`
	// EndBrace string `"}"`
}

func (self *CustomTag) ToHtml() string {
	templateName := "templates/articleParser/" + self.GetTemplate()
	tmpl, err := template.ParseFiles(templateName)
    if err != nil {
        panic(err)
    }

	var doc bytes.Buffer
    if err := tmpl.Execute(&doc, self); err != nil {
		panic(err)
    }
	return doc.String()
}

func (self *CustomTag) GetTemplate() string {
	switch self.Type {
	case "P": return  "p.tmpl"
	case "H1": return "h1.tmpl"
	case "Video": return "video.tmpl"
	default:  return "p.tmpl"
	}
}


func ParseArticle(articleRawText string) string{
	// parser, err := participle.Build(&Article{}, nil)
	parser, err := participle.Build(&Article{}, nil)
	if err != nil {
		panic(err)
	}
	article := &Article{}

	err = parser.ParseString(articleRawText, article)

	if err != nil {
		panic(err)
	}
	// repr.Println(article, repr.Indent("  "), repr.OmitEmpty(true))

	return article.ToHtml()
}

