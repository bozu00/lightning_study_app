
package services

import (
	// "os"

	// "github.com/alecthomas/participle"
	// "github.com/alecthomas/participle/lexer"
	// "github.com/alecthomas/repr"
	// "fmt"
	"regexp"
	//"strings"
	// "io/ioutil"
	"log"
	"bytes"
	"io"
	"fmt"
)



func Parse(articleRawText string) string{


	templateName := "/templates/parserItem/"
	// tmplH1, err := ioutil.ReadFile( templateName + "h1.tmpl")
    tmplH1, err := readTmplFile(templateName + "h1.tmpl")
    if err != nil {
        panic(err)
    }

    tmplP, err := readTmplFile(templateName + "p.tmpl")
    // tmplP, err := ioutil.ReadFile( templateName + "p.tmpl")
    if err != nil {
        panic(err)
    }

    tmplVideo, err := readTmplFile(templateName + "video.tmpl")
    // tmplVideo, err := ioutil.ReadFile( templateName + "video.tmpl")
    if err != nil {
        panic(err)
    }

    tmplImage, err := readTmplFile(templateName + "image.tmpl")
    // tmplImage, err := ioutil.ReadFile( templateName + "image.tmpl")
    if err != nil {
        panic(err)
    }

	res := articleRawText
	
	// セクション開始
	re, _ := regexp.Compile(`\[([^/]+?)\]`)
	res = re.ReplaceAllString(res, "<div class='articleSection'>")

	// セクション終了
	re, _ = regexp.Compile(`\[\/(.+?)\]`)
	res = re.ReplaceAllString(res, "</div>")
	// h1

	re, _ = regexp.Compile(`H1\s*\{([^}]+)\}`)
	res = re.ReplaceAllString(res, string(tmplH1))

	// p
	re, _ = regexp.Compile(`P \s*\{([^}]+)\}`)
	res = re.ReplaceAllString(res, string(tmplP))

	// video
	re, _ = regexp.Compile(`Video \s*\{([^}]+)\}`)
	res = re.ReplaceAllString(res, string(tmplVideo))

	// video
	re, _ = regexp.Compile(`Image \s*\{([^}]+)\}`)
	res = re.ReplaceAllString(res, string(tmplImage))


	// re, _ = regexp.Compile(`H1\s*\{(.+?)\}`)
	// src = re.ReplaceAllString(src, "$1")

	return res 
}

func readTmplFile(filename string) (string, error) {
	f, err := ArticlePartsAssets.Open(filename)
	if err != nil {
		log.Println(err)
	}

	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, f)
	if err != nil {
		log.Println(err)
	}

	tplString := fmt.Sprintf("%s", buf)
	return tplString, nil
}
