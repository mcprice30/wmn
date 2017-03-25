// Package sensor represents a sensor that is sending data to the communication
// hub.
package main

import (
	"fmt"
	"net"
	"os"

	"github.com/mcprice30/wmn/data"
	"github.com/mcprice30/wmn/sensor"
)

func main() {
	/*
		streams := map[string]sensor.SensorStream{
			"heartrate": sensor.CreateHeartRateSensor(),
			"location":  sensor.CreateLocationSensor(),
			"oxygen":    sensor.CreateOxygenSensor(),
			"gas":       sensor.CreateGasSensor(),
		}

		sensor.Run(streams[os.Args[1]])

	*/
	go sensorHubListen("localhost:5001")
	go sensor.Run(sensor.CreateHeartRateSensor(), "localhost:5002", "localhost:5001")
	go sensor.Run(sensor.CreateLocationSensor(), "localhost:5003", "localhost:5001")
	go sensor.Run(sensor.CreateOxygenSensor(), "localhost:5004", "localhost:5001")
	go sensor.Run(sensor.CreateGasSensor(), "localhost:5005", "localhost:5001")

	for {
	}
}

func sensorHubListen(address string) {
	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot resolve address (%s): %s\n", address, err)
		return
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot listen for packets: %s\n", err)
		return
	}
	defer conn.Close()

	unmarshaller := data.CreateByteUnmarshaller()

	for {
		buffer := make([]byte, 64)
		if n, err := conn.Read(buffer); err != nil {
			fmt.Fprintf(os.Stderr, "Cannot read data sent: %s\n", err)
		} else {
			fmt.Println(unmarshaller.Unmarshal(buffer[:n]))
		}
	}
}
