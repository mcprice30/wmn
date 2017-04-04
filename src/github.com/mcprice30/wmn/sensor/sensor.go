// Package sensor contains implementations of the various sensors that send
// data to the Sensor Hub.
//
// Currently, this consists of:
//
// GasSensor:		Monitors the levels of toxic gas. Data is sent as data.GasData
// packets, at an interval of 250ms.
//
// HeartRateSensor: Monitors the wearer's heart rate. Data is sent as
// data.HeartRateData packets, at an interval of 1000ms.
//
// LocationSensor: Monitors the current location of the wearer. Data is sent
// as data.LocationData packets, at an interval of 500ms.
//
// OxygenSensor: Monitors the oxygen level of the wearer's air-pack. Data is
// sent as data.OxygenData packets, at an interval of 2000ms.
package sensor

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/mcprice30/wmn/data"
)

// SensorStream is the interface that all data streams must implement in order
// to be sent to the sensor hub.
type SensorStream interface {
	// GetData will poll a new data point from the sensor. It is the
	// responsiblity of the caller to ensure the appropriate send ratefor the
	// given stream; GetData should not attempt to regulate the interval.
	GetData() data.SensorData

	// Interval indicates the duration to wait between calling the sensor.
	Interval() time.Duration
}

// Run will simulate a SensorStream sending data. It will generate a new data
// point from the SensorStream and send it to the sensor hub, at an interval
// dictated by the given SensorStream.
func Run(s SensorStream, from, to string) {
	srcAddr, err := net.ResolveUDPAddr("udp", from)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot resolve address (%s): %s\n", from, err)
		return
	}

	dstAddr, err := net.ResolveUDPAddr("udp", to)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot resolve address (%s): %s\n", to, err)
		return
	}

	conn, err := net.DialUDP("udp", srcAddr, dstAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot dial address (%v): %s\n", to, err)
		return
	}

	ticker := time.NewTicker(s.Interval())
	for {
		<-ticker.C
		if _, err := conn.Write(s.GetData().ToBytes()); err != nil {
			fmt.Fprintf(os.Stderr, "Error sending data to hub: %s\n", err)
		}
	}
}

// intervalFromString will parse intervalStr, a string describing an interval
// as described in time.ParseDuration. In the event that intervalStr
// is invalid, this will panic.
func intervalFromString(intervalStr string) time.Duration {
	if interval, err := time.ParseDuration(intervalStr); err == nil {
		return interval
	} else {
		panic(err)
	}
}
