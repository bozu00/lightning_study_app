
GOSRCS               := $(shell find . -type f -name '*.go' -not -iwholename './vendor/*')
TEMPLATE_SRCS        := $(shell find ./templates -type f -name '*.html' -o -name '*.tmpl')
PUBLIC_ASSET_SRCS    := $(shell find ./assets -type f -name '.'  -not -iwholename '*/.DS_Store')

main: templateAssets.go publicAssets.go $(GOSRCS)
	go build -o main main.go publicAssets.go templateAssets.go

# .PHONY: templateAssets.go
templateAssets.go: $(TEMPLATE_SRCS)
	go-assets-builder -o templateAssets.go -p main -v TemplateAssets templates/

# .PHONY: publicAssets.go
publicAssets.go: $(PUBLIC_ASSET_SRCS)
	go-assets-builder -o publicAssets.go -p main -v PublicAssets assets/

