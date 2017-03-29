// Package data contains data types that are common across the project,
// allowing easy access.
//
// It currently contains the following data types:
//
// SensorData: an interface for encapsulating data from various sensors,
// and implementations for data from the toxic gas, heart rate, location,
// and oxygen level sensors.
//
// SensorUnmarshaller: an interface for retrieving data from the raw bytes
// that it will be sent across the network in, and implementations for data
// from the toxic gas, heart rate, location, and oxygen level sensors.
//
// ByteUnmarshaller: ByteUnmarshaller delegates to various SensorUnmarshallers
// to automatically decode bytes received from the network.
//
// PacketHeader: PacketHeader contains the information stored inside the header
// of a packet sent over the network.
//
// Packet: Packet encapsulates the entirety of the packet to be sent across
// the network.
package data

import (
	"fmt"
)

const NumSensorTypes = 4

// SensorData defines the various types of data that can be obtained from
// a sensor.
type SensorData interface {

	// SensorData inherits String() string from fmt.Stringer.
	fmt.Stringer

	// Id indicates the sequence ID of this element of data among all data points
	// generated from this sensor.
	Id() byte

	// Type indicates what sensor this element of data came from, as defined
	// by <SensorName>Type constants in the package.
	Type() byte

	// NumBytes indicates how many bytes this data point will be marshalled to.
	NumBytes() int

	// ToBytes will marshall this data point into bytes in order to be put in a
	// packet.
	ToBytes() []byte
}
