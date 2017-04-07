all: test readme

test: generate
	go test ./...

generate:
	go generate

readme:
	godoc2md github.com/leighmcculloch/optional > README.md

setup:
	go get github.com/ncw/gotemplate
	go get github.com/davecheney/godoc2md
