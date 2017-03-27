package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mcprice30/wmn/sensor"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "<sensor name>")
		fmt.Println("Available Sensors: heartrate, location, oxygen, gas")
		os.Exit(127)
	}

	switch strings.ToLower(os.Args[1]) {

	case "heartrate":
		sensor.Run(sensor.CreateHeartRateSensor(), "localhost:5002",
			"localhost:5001")
	case "location":
		sensor.Run(sensor.CreateLocationSensor(), "localhost:5003",
			"localhost:5001")
	case "oxygen":
		sensor.Run(sensor.CreateOxygenSensor(), "localhost:5004",
			"localhost:5001")
	case "gas":
		sensor.Run(sensor.CreateGasSensor(), "localhost:5005",
			"localhost:5001")
	default:
		fmt.Println("Unknown sensor", os.Args[1])
		fmt.Println("Available Sensors: heartrate, location, oxygen, gas")
		os.Exit(127)
	}
}
