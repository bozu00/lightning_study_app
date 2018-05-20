package services

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _ArticlePartsAssets779d21283d57064ffb69819bcecdd2af1e6b218f = "<div class=\"articleSection__video\">\n    <img src=\"/uploads/image/$1\">\n</div>\n"
var _ArticlePartsAssetsd3396e55840b4cd6ee1ec4d5a8720777680bf8ee = "<div class=\"articleSection__description\"><p>$1</p></div>\n"
var _ArticlePartsAssets539c0bd7123bc711d43bd31c98ef36df04dec943 = "<div class=\"articleSection__video\"><iframe src=\"https://www.youtube.com/embed/$1?rel=0&playsinline=1&showinfo=0&autoplay=1&color=white&controls=1&modestbranding=1&start=20&end=30\" frameborder=\"0\" allow=\"autoplay; encrypted-media\" allowfullscreen></iframe></div>\n"
var _ArticlePartsAssets79d1b284286558b7f3a2052f5dfd07c01d62d635 = "<h1 class=\"articleSection__title\">$1</h1>\n"

// ArticlePartsAssets returns go-assets FileSystem
var ArticlePartsAssets = assets.NewFileSystem(map[string][]string{"/": []string{"templates"}, "/templates": []string{"parserItem"}, "/templates/parserItem": []string{"h1.tmpl", "image.tmpl", "p.tmpl", "video.tmpl"}}, map[string]*assets.File{
	"/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1526819442, 1526819442000000000),
		Data:     nil,
	}, "/templates": &assets.File{
		Path:     "/templates",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1526308817, 1526308817000000000),
		Data:     nil,
	}, "/templates/parserItem": &assets.File{
		Path:     "/templates/parserItem",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1526308817, 1526308817000000000),
		Data:     nil,
	}, "/templates/parserItem/h1.tmpl": &assets.File{
		Path:     "/templates/parserItem/h1.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1526308817, 1526308817000000000),
		Data:     []byte(_ArticlePartsAssets79d1b284286558b7f3a2052f5dfd07c01d62d635),
	}, "/templates/parserItem/image.tmpl": &assets.File{
		Path:     "/templates/parserItem/image.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1526308817, 1526308817000000000),
		Data:     []byte(_ArticlePartsAssets779d21283d57064ffb69819bcecdd2af1e6b218f),
	}, "/templates/parserItem/p.tmpl": &assets.File{
		Path:     "/templates/parserItem/p.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1526308817, 1526308817000000000),
		Data:     []byte(_ArticlePartsAssetsd3396e55840b4cd6ee1ec4d5a8720777680bf8ee),
	}, "/templates/parserItem/video.tmpl": &assets.File{
		Path:     "/templates/parserItem/video.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1526308817, 1526308817000000000),
		Data:     []byte(_ArticlePartsAssets539c0bd7123bc711d43bd31c98ef36df04dec943),
	}}, "")
