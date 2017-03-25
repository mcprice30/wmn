// Package sensor represents a sensor that is sending data to the communication
// hub.
package main

import (
	"github.com/mcprice30/wmn/sensor"
)

func main() {

	// Listen for data from the sensors.
	go sensor.SensorHubListen("localhost:5001")

	// Launch each sensor in a separate goroutine (thread).
	go sensor.Run(sensor.CreateHeartRateSensor(), "localhost:5002",
		"localhost:5001")
	go sensor.Run(sensor.CreateLocationSensor(), "localhost:5003",
		"localhost:5001")
	go sensor.Run(sensor.CreateOxygenSensor(), "localhost:5004",
		"localhost:5001")
	go sensor.Run(sensor.CreateGasSensor(), "localhost:5005",
		"localhost:5001")

	// Spin forever.
	for {
	}
}
