package sensor

import (
	"fmt"
	"net"
	"os"

	"github.com/mcprice30/wmn/data"
	"github.com/mcprice30/wmn/network"
	"github.com/mcprice30/wmn/transport"
)

// SensorHub represents the part of the codebase respsonsible for listening to
// the various sensors, before building and sending data packets to the fire
// chief.
type SensorHub struct {
	listenAddr string

	transmitDst data.ManetAddr

	currPacketBytes int
	currPacket      *data.DataPacket

	rc *transport.ReliableSender
}

// CreateSensorHub will create a sensor hub that listens on the given address
// for incoming sensor packets.
func CreateSensorHub(listenAddr, dstName string) *SensorHub {
	hub := &SensorHub{
		listenAddr:      listenAddr,
		transmitDst:     network.GetAddrFromHostname(dstName),
		currPacketBytes: data.PacketHeaderBytes,
	}
	hub.currPacket = hub.newTransmitPacket()
	return hub
}

// Listen will cause the given sensor hub to listen to incoming packets
// building and sending data packets to the fire chief.
func (hub *SensorHub) Listen() {
	hub.rc = transport.CreateReliableSender()
	conn, err := createConnectionToSensor(hub.listenAddr)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer conn.Close()
	hub.readFromConnectionToSensor(conn)
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
func (hub *SensorHub) readFromConnectionToSensor(conn net.Conn) {
	for {
		buffer := make([]byte, 64)
		if n, err := conn.Read(buffer); err != nil {
			fmt.Fprintf(os.Stderr, "Cannot read data sent: %s\n", err)
		} else {
			hub.addToTransmitPacket(data.SensorDataFromBytes(buffer[:n]))
		}
	}
}

func (hub *SensorHub) addToTransmitPacket(sd data.SensorData) {
	if hub.currPacketBytes+sd.NumBytes() > data.MaxPacketBytes {
		hub.rc.Transmit(hub.currPacket)
		hub.currPacketBytes = data.PacketHeaderBytes
		hub.currPacket = hub.newTransmitPacket()
	} else {
		hub.currPacketBytes += sd.NumBytes()
		hub.currPacket.Body = append(hub.currPacket.Body, sd)
	}
}

func (hub *SensorHub) newTransmitPacket() *data.DataPacket {
	return &data.DataPacket{
		Header: data.PacketHeader{
			DestinationAddress: hub.transmitDst,
			PacketType:         data.PacketTypeData,
		},
		Body: []data.SensorData{},
	}
}
