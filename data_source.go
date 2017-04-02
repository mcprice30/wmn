package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mcprice30/wmn/sensor"
)

func main() {

	if len(os.Args) != 4 {
		fmt.Println("Usage:", os.Args[0], "<sensor name> <generate port> <hub port>")
		fmt.Println("Available Sensors: heartrate, location, oxygen, gas")
		os.Exit(127)
	}

	srcAddr := fmt.Sprintf("localhost:%s", os.Args[2])
	dstAddr := fmt.Sprintf("localhost:%s", os.Args[3])

	switch strings.ToLower(os.Args[1]) {

	case "heartrate":
		sensor.Run(sensor.CreateHeartRateSensor(), srcAddr, dstAddr)
	case "location":
		sensor.Run(sensor.CreateLocationSensor(), srcAddr, dstAddr)
	case "oxygen":
		sensor.Run(sensor.CreateOxygenSensor(), srcAddr, dstAddr)
	case "gas":
		sensor.Run(sensor.CreateGasSensor(), srcAddr, dstAddr)
	default:
		fmt.Println("Unknown sensor", os.Args[1])
		fmt.Println("Available Sensors: heartrate, location, oxygen, gas")
		os.Exit(127)
	}
}
