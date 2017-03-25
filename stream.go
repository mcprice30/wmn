// Package sensor represents a sensor that is sending data to the communication
// hub.
package main

import (
	"os"

	"github.com/mcprice30/wmn/sensor"
)

func main() {
	sensors := map[string]sensor.Sensor{
		"heartrate": sensor.CreateHeartRateSensor(1000),
		"location":  sensor.CreateLocationSensor(500),
		"oxygen":    sensor.CreateOxygenSensor(2000),
		"gas":       sensor.CreateGasSensor(250),
	}

	s := sensor.CreateSender(sensors[os.Args[1]])
	s.Run()
}
