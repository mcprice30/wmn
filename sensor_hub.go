// Package sensor represents a sensor that is sending data to the communication
// hub.
package main

import (
	"github.com/mcprice30/wmn/config"
	"github.com/mcprice30/wmn/sensor"
)

func main() {

	config.LoadConfig("config_test.txt", "Sensor")

	// Listen for data from the sensors.
	hub := sensor.CreateSensorHub("localhost:5001", 0x0002)
	hub.Listen()
}
