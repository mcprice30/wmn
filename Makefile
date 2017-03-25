PREFIX="github.com/mcprice30/wmn"

.PHONY: stream all

all:
	go build $(PREFIX)/data
	go build $(PREFIX)/sensor

stream:
	go build -o bin/stream stream.go

