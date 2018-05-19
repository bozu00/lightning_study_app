

main: templateAssets.go publicAssets.go
	go build -o main main.go publicAssets.go templateAssets.go

templateAssets.go:
	go-assets-builder -o templateAssets.go -p main -v TemplateAssets templates/

publicAssets.go:
	go-assets-builder -o publicAssets.go -p main -v PublicAssets assets/

