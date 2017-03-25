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
	"bytes"
	"encoding/binary"
	"math"
)

// ByteOrder indicates the endianness the data should take when transmitted
// across the network.
var ByteOrder = binary.BigEndian

// float64ToBytes converts the given float64 to a slice of 8 bytes equivalent
// to the float.
func float64ToBytes(in float64) []byte {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, ByteOrder, in); err != nil {
		panic(err)
	}
	return buf.Bytes()
}

// bytesToFloat64 takes a slice of bytes and returns the float64 they represent.
func bytesToFloat64(in []byte) float64 {
	bits := ByteOrder.Uint64(in)
	return math.Float64frombits(bits)
}
