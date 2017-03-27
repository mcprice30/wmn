// Package sensor represents a sensor that is sending data to the communication
// hub.
package main

import (
	"github.com/mcprice30/wmn/network"
	"github.com/mcprice30/wmn/sensor"
)

func main() {

	network.SetMyAddress(0x0001)

	// Listen for data from the sensors.
	hub := sensor.CreateSensorHub("localhost:5001", 0x0001, 0x0002)
	hub.Listen()
}
