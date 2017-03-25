// Package sensor represents a sensor that is sending data to the communication
// hub.
package main

import (
	"os"

	"github.com/mcprice30/wmn/sensor"
)

func main() {
	streams := map[string]sensor.SensorStream{
		"heartrate": sensor.CreateHeartRateSensor(),
		"location":  sensor.CreateLocationSensor(),
		"oxygen":    sensor.CreateOxygenSensor(),
		"gas":       sensor.CreateGasSensor(),
	}

	sensor.Run(streams[os.Args[1]])
}
