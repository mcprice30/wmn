PREFIX="github.com/mcprice30/wmn"

.PHONY: sensor_hub all fmt packages

all: packages sensor_hub

packages:
	go build $(PREFIX)/data
	go build $(PREFIX)/sensor

sensor_hub:
	go build -o bin/sensor_hub sensor_hub.go

fmt:
	go fmt $(PREFIX)/data
	go fmt $(PREFIX)/sensor
	go fmt *.go

clean:
	rm bin/*
