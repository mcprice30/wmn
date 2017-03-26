// Package sensor represents a sensor that is sending data to the communication
// hub.
package main

import (
	"github.com/mcprice30/wmn/network"
	"github.com/mcprice30/wmn/transport"
)

func main() {

	network.SetMyAddress(0x0002)
	go transport.RunEcho(0x0002)

	// Spin forever.
	for {
	}
}
