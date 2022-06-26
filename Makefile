.PHONY: fmt clean deps build run add-cmd run-cmd statik-generate

VERSION = 0.0.1
BUILD_FLAGS = -ldflags "\
                -X main.Version=$(VERSION) \
                "
OUTPUT_DIR = ./bin

# 引数
ARG = version

all: clean deps test build

fmt:
	go fmt ./...

run:
	go run main.go

deps:
	go get -d -v ./...

test:
	go test -v ./...

build: fmt
	go build -o $(OUTPUT_DIR)

clean:
	go clean

add-cmd:
	cobra-cli add ${ARG}

run-cmd:
	go run main.go ${ARG}

aa-generate:
	rm -r statik
	statik -src=./aa