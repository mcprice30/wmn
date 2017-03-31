PREFIX="github.com/mcprice30/wmn"

.PHONY: sensor_hub all fmt packages display_hub manet_node data_source config

all: packages sensor_hub display_hub manet_node data_source

packages:
	go build $(PREFIX)/data
	go build $(PREFIX)/sensor
	go build $(PREFIX)/network
	go build $(PREFIX)/transport
	go build $(PREFIX)/chief
	go build $(PREFIX)/config

sensor_hub:
	go build -o bin/sensor_hub sensor_hub.go

display_hub:
	go build -o bin/display_hub display_hub.go

manet_node:
	go build -o bin/manet_node manet_node.go

data_source:
	go build -o bin/data_source data_source.go

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

start:
	bin/manet_node  &
	bin/display_hub &
	bin/sensor_hub &
	bin/data_source heartrate > /dev/null &
	bin/data_source location > /dev/null &
	bin/data_source oxygen > /dev/null &
	bin/data_source gas > /dev/null &

kill:
	pidof `ls ./bin` | xargs kill
