PREFIX="github.com/mcprice30/wmn"

.PHONY: sensor_hub all fmt packages display_hub

all: packages sensor_hub display_hub

packages:
	go build $(PREFIX)/data
	go build $(PREFIX)/sensor
	go build $(PREFIX)/network
	go build $(PREFIX)/transport

sensor_hub:
	go build -o bin/sensor_hub sensor_hub.go

display_hub:
	go build -o bin/display_hub display_hub.go

fmt:
	go fmt $(PREFIX)/data
	go fmt $(PREFIX)/sensor
	go fmt $(PREFIX)/network
	go fmt $(PREFIX)/transport
	go fmt *.go

test:
	go test $(PREFIX)/data

clean:
	rm bin/*
