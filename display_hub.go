// Package sensor represents a sensor that is sending data to the communication
// hub.
package main

import (
	"github.com/mcprice30/wmn/chief"
	"github.com/mcprice30/wmn/config"
)

func main() {

	config.LoadConfig("config_test.txt", "Display")
	chief.RunListener()
}
