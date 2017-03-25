PREFIX="github.com/mcprice30/wmn"

.PHONY: stream all

all:
	go build $(PREFIX)/data
	go build $(PREFIX)/sensor

stream:
	go build -o bin/stream stream.go

fmt:
	go fmt $(PREFIX)/data
	go fmt $(PREFIX)/sensor
	go fmt *.go

