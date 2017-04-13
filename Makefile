all: test readme

test: generate
	go test ./...

generate:
	rm -f *_generated.go
	go generate
	rename 's/gotemplate_([A-Za-z0-9]+)\.go/\L$$1_generated.go/' gotemplate_*.go

readme:
	godocdown github.com/leighmcculloch/go-optional > README.md

setup:
	go get github.com/ncw/gotemplate
	go get github.com/davecheney/godoc2md
