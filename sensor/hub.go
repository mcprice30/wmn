package sensor

import (
	"fmt"
	"net"
	"os"

	"github.com/mcprice30/wmn/data"
)

// SensorHubListen is responsible for acting as the sensor hub, listening on
// the given address for incoming data packets.
//
// Eventually, this will be responsible for building and sending data packets
// to the Fire Chief.
func SensorHubListen(address string) {
	conn, err := createConnectionToSensor(address)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	defer conn.Close()
	readFromConnectionToSensor(conn)
}

// createConnectionToSensor will attempt listen to data from sensors at the
// given address. If successful, it will return the connection. In all other
// cases, it will return nil and an error indicating what occurred.
func createConnectionToSensor(address string) (net.Conn, error) {
	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		return nil, fmt.Errorf("Cannot resolve address (%s): %s", address, err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return nil, fmt.Errorf("Cannot listen for packets: %s", err)
	}

	return conn, nil
}

// readFromConnectionToSensor will repeatedly listen to data sent by various
// sensors.
func readFromConnectionToSensor(conn net.Conn) {
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
