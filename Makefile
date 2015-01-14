.PHONY: all run preprocess

all:
	#gofmt -w src/
	GOPATH=$(shell pwd):${GOPATH} GOBIN=$(shell pwd)/bin go install -v -gcflags -N ./...

preprocess:
	bash convert.sh
	GOPATH=$(shell pwd):${GOPATH} GOBIN=$(shell pwd)/bin go get "github.com/tchap/go-patricia/patricia"

run: all
	bin/search
