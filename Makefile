
# GOSRCS                   := $(shell find . -type f -name '*.go' -not -iwholename './vendor/*' -not -name '*articlePartsAssets.go')
GOSRCS                   := $(shell find . -type f -name '*.go' -not -iwholename './vendor/*')
TEMPLATE_SRCS            := $(shell find ./templates -type f -name '*.html')
ARTICLE_PARTS_ASSET_SRCS := $(shell find ./templates/parserItem -type f -name '*.tmpl'  -not -iwholename '*/.DS_Store')
PUBLIC_ASSET_SRCS        := $(shell find ./assets -type f -name '*'  -not -iwholename '*/.DS_Store')

# main: templateAssets.go publicAssets.go src/services/articlePartsAssets.go $(GOSRCS)
main: templateAssets.go publicAssets.go $(GOSRCS)
	go build -o main main.go publicAssets.go templateAssets.go

# .PHONY: templateAssets.go
templateAssets.go: $(TEMPLATE_SRCS)
	go-assets-builder -o templateAssets.go -p main -v TemplateAssets templates/

# .PHONY: publicAssets.go
publicAssets.go: $(PUBLIC_ASSET_SRCS)
	go-assets-builder -o publicAssets.go -p main -v PublicAssets assets/

src/services/articlePartsAssets.go: $(ARTICLE_PARTS_ASSET_SRCS)
	go-assets-builder -o src/services/articlePartsAssets.go -p services -v ArticlePartsAssets templates/parserItem/

development: templateAssets.go $(GOSRCS)
	go build -o main main.go publicAssets.go templateAssets.go
