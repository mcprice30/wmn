package main

import (
	"fmt"
	"os"

	"github.com/mcprice30/wmn/config"
	"github.com/mcprice30/wmn/sensor"
)

func main() {

	if len(os.Args) != 5 {
		fmt.Println("Use:", os.Args[0], "<config file> <error file> <hostname> <listenport>")
		os.Exit(127)
	}

	config.LoadConfig(os.Args[1], os.Args[3])

	// Listen for data from the sensors.
	config.ListenForErrorChanges(os.Args[2])
	listenAddr := fmt.Sprintf("localhost:%s", os.Args[4])
	hub := sensor.CreateSensorHub(listenAddr, "Display")
	hub.Listen()
}
