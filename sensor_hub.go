package main

import (
	"fmt"
	"os"

	"github.com/mcprice30/wmn/config"
	"github.com/mcprice30/wmn/sensor"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Use:", os.Args[0], "<hostname> <listenport>")
		os.Exit(127)
	}

	config.LoadConfig("config_test.txt", os.Args[1])

	// Listen for data from the sensors.
	listenAddr := fmt.Sprintf("localhost:%s", os.Args[2])
	config.ListenForErrorChanges("error_test.txt")
	hub := sensor.CreateSensorHub(listenAddr, "Display")
	hub.Listen()
}
