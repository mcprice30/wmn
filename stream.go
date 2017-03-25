// Package sensor represents a sensor that is sending data to the communication
// hub.
package main

import (
	"fmt"
	"os"

	"github.com/mcprice30/wmn/sensor"
)

func main() {
	fmt.Println(os.Args)

	s := sensor.CreateSender(sensor.CreateHeartRateSensor(10))
	s.Run()
}
