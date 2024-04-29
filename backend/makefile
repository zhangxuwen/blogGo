.PHONY: all build run gotool clean help

BINARY="blogWeb"

all: gotool build

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BINARY}

run: 
	@go run ./main.go conf/config.yaml
	#加上@就不会在终端打印命令

gotool:
	go fmt ./
	go vet ./

clean:
	@if [ -f ${BINARY} ] ; THEN RM ${BINARY} ; fi
