.PHONY: build run

APPLICATION_NAME=tt-apiserver

build:
	go build -o bin/${APPLICATION_NAME} ./cmd/${APPLICATION_NAME}

run: build
	./bin/${APPLICATION_NAME}
