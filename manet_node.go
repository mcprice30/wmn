// Package sensor represents a sensor that is sending data to the communication
// hub.
package main

import (
	"fmt"

	"github.com/mcprice30/wmn/config"
	//"github.com/mcprice30/wmn/data"
	"github.com/mcprice30/wmn/network"
)

func main() {

	config.LoadConfig("config_test.txt", "Node1")
	conn := network.CreateManet()
	for {
		fmt.Println(conn.Receive())
	}
}
