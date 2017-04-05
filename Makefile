# Constants for running go
SETPATH=GOPATH="$(shell pwd)"
GOCMD=gobin/go/bin/go
GO=$(SETPATH) $(GOCMD)

# Package location
PREFIX="github.com/mcprice30/wmn"

# Constants for installing go.
GO_DOWNLOAD_SRC="https://storage.googleapis.com/golang/go1.8.linux-amd64.tar.gz"
GO_TARBALL=godownload.tar.gz
GROOT=$(shell pwd)/gobin/go
export GOROOT:=$(GROOT)

.PHONY: sensor_hub all fmt packages display_hub manet_node data_source config

# All will compile all packages and binaries.
all: packages sensor_hub display_hub manet_node data_source

# Packages separately compiles each package
packages:
	$(GO) build $(PREFIX)/data
	$(GO) build $(PREFIX)/sensor
	$(GO) build $(PREFIX)/network
	$(GO) build $(PREFIX)/transport
	$(GO) build $(PREFIX)/chief
	$(GO) build $(PREFIX)/config

# Rules for building packages
sensor_hub:
	$(GO) build -o bin/sensor_hub sensor_hub.go

display_hub:
	$(GO) build -o bin/display_hub display_hub.go

manet_node:
	$(GO) build -o bin/manet_node manet_node.go

data_source:
	$(GO) build -o bin/data_source data_source.go

# Format all go code
fmt:
	$(GO) fmt $(PREFIX)/data
	$(GO) fmt $(PREFIX)/sensor
	$(GO) fmt $(PREFIX)/network
	$(GO) fmt $(PREFIX)/transport
	$(GO) fmt $(PREFIX)/chief
	$(GO) fmt $(PREFIX)/config
	$(GO) fmt *.go

# Run all unit tests
test:
	$(GO) test $(PREFIX)/data

# Get rid of old binaries.
clean:
	rm bin/*

# Default to the test run.
start: start_test

start_test:
	bin/manet_node config_test.txt Node1 &
	bin/display_hub config_test.txt error_test.txt Display 10109 &
	bin/sensor_hub config_test.txt error_test.txt Sensor 10100 &
	bin/data_source heartrate 10108 10100 &
	bin/data_source location 10107 10100 &
	bin/data_source oxygen 10106 10100 &
	bin/data_source gas 10105 10100 &

# MUST BE CALLED ON TUX 203
start_demo:
	ssh mcp0038@tux204 'cd Lab2/wmn && bin/manet_node config_test.txt Node1 &'
	ssh mcp0038@tux204 'cd Lab2/wmn && bin/manet_node config_test.txt Node2 &'
	ssh mcp0038@tux204 'cd Lab2/wmn && bin/manet_node config_test.txt Node3 &'
	ssh mcp0038@tux204 'cd Lab2/wmn && bin/manet_node config_test.txt Node4 &'
	ssh mcp0038@tux204 'cd Lab2/wmn && bin/manet_node config_test.txt Node5 &'
	ssh mcp0038@tux204 'cd Lab2/wmn && bin/manet_node config_test.txt Node6 &'
	ssh mcp0038@tux204 'cd Lab2/wmn && bin/manet_node config_test.txt Node7 &'
	ssh mcp0038@tux204 'cd Lab2/wmn && bin/manet_node config_test.txt Node8 &'
	ssh mcp0038@tux204 'cd Lab2/wmn && bin/manet_node config_test.txt Node9 &'
	ssh mcp0038@tux204 'cd Lab2/wmn && bin/manet_node config_test.txt Node10 &'
	ssh mcp0038@tux205 'cd Lab2/wmn && bin/manet_node config_test.txt Node11 &'
	ssh mcp0038@tux205 'cd Lab2/wmn && bin/manet_node config_test.txt Node12 &'
	ssh mcp0038@tux205 'cd Lab2/wmn && bin/manet_node config_test.txt Node13 &'
	ssh mcp0038@tux205 'cd Lab2/wmn && bin/display_hub config_test.txt error_test.txt Display 10109 &'
	bin/sensor_hub config_test.txt error_test.txt Sensor 10100 &
	bin/data_source heartrate 10108 10100 &
	bin/data_source location 10107 10100 &
	bin/data_source oxygen 10106 10100 &
	bin/data_source gas 10105 10100 &

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
