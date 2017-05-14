.PHONY: all run build test install

all: run

run:
	go run vlbv.go

build:
	go build vlbv.go

test:
	go test -v ./...

install:
	go get -u -v github.com/golang/dep/cmd/dep
	dep ensure
