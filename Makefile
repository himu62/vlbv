.PHONY: run lint test\
	client/watch client/build client/test client/lint client/install\
	server/run server/build server/test server/lint server/install

run: lint test client/bulid server/run

lint: client/lint server/lint

test: client/test server/test

client/watch:
	cd client; yarn run watch

client/build:
	cd client; yarn run build

client/test:
	cd client; yarn run test

client/lint:
	cd client; yarn run lint

client/install:
	cd client; yarn install

server/run:
	go run server/vlbv.go

server/build:
	go build -o vlbv.exe server/vlbv.go

server/test:
	go test server

server/lint:
	golint server

server/install:
	cd server; dep ensure
