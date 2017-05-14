.PHONY: all run build test install update

all: run

run: build
	vlbv.exe

build:
	rsrc -manifest vlbv.manifest -o rsrc.syso
	go build -o vlbv.exe -ldflags="-H windowsgui"

test:
	go test -v ./...

install:
	go get -u -v github.com/golang/dep/cmd/dep
	go get -u -v github.com/akavel/rsrc
	dep ensure -v -update

update:
	dep ensure -v -update
