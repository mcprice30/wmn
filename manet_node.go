// Package sensor represents a sensor that is sending data to the communication
// hub.
package main

import (
	"fmt"

	"github.com/mcprice30/wmn/data"
	"github.com/mcprice30/wmn/network"
)

func main() {

	network.SetMyAddress(0x0003)
	conn := network.BindManet()
	conn.SetNeighbors([]data.ManetAddr{0x0001, 0x0002})
	// Spin forever.
	for {
		fmt.Println(conn.Receive())
	}
}
