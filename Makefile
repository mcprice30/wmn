SETPATH=GOPATH="$(shell pwd)"
GOCMD=gobin/go/bin/go
GO=$(SETPATH) $(GOCMD)

PREFIX="github.com/mcprice30/wmn"

GO_DOWNLOAD_SRC="https://storage.googleapis.com/golang/go1.8.linux-amd64.tar.gz"
GO_TARBALL=godownload.tar.gz
GROOT=$(shell pwd)/gobin/go
export GOROOT:=$(GROOT)

.PHONY: sensor_hub all fmt packages display_hub manet_node data_source config

all: packages sensor_hub display_hub manet_node data_source

packages:
	$(GO) build $(PREFIX)/data
	$(GO) build $(PREFIX)/sensor
	$(GO) build $(PREFIX)/network
	$(GO) build $(PREFIX)/transport
	$(GO) build $(PREFIX)/chief
	$(GO) build $(PREFIX)/config

sensor_hub:
	$(GO) build -o bin/sensor_hub sensor_hub.go

display_hub:
	$(GO) build -o bin/display_hub display_hub.go

manet_node:
	$(GO) build -o bin/manet_node manet_node.go

data_source:
	$(GO) build -o bin/data_source data_source.go

fmt:
	$(GO) fmt $(PREFIX)/data
	$(GO) fmt $(PREFIX)/sensor
	$(GO) fmt $(PREFIX)/network
	$(GO) fmt $(PREFIX)/transport
	$(GO) fmt $(PREFIX)/chief
	$(GO) fmt $(PREFIX)/config
	$(GO) fmt *.go

test:
	$(GO) test $(PREFIX)/data

clean:
	rm bin/*

start: start_test

start_test:
	bin/manet_node Node1 &
	bin/display_hub Display 12345 &
	bin/sensor_hub Sensor 5001 &
	bin/data_source heartrate 5002 5001 > /dev/null &
	bin/data_source location 5003 5001 > /dev/null &
	bin/data_source oxygen 5004 5001 > /dev/null &
	bin/data_source gas 5005 5001 > /dev/null &

kill:
	pidof `ls ./bin` | xargs kill

setup:
	@mkdir -p bin
	@mkdir -p gobin
	@echo "Downloading Go..."
	@curl $(GO_DOWNLOAD_SRC) > gobin/$(GO_TARBALL) && \
	echo "Extracting Go..." && \
	tar -xzf gobin/$(GO_TARBALL) -C gobin && \
	echo "Go installed successfully!"
	@echo "Installing dependencies!"
	@$(GO) get golang.org/x/net/websocket
	@echo "Setup complete!"
