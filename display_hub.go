// Package sensor represents a sensor that is sending data to the communication
// hub.
package main

import (
	"github.com/mcprice30/wmn/network"
	"github.com/mcprice30/wmn/chief"
)

func main() {

	network.SetMyAddress(0x0002)
	chief.RunListener(0x0002)
}
