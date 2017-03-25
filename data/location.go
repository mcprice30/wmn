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

// LocationDataType indicates what the first byte of LocationData structs will
// be when marshalled into bytes for transmission.
const LocationDataType byte = 1

// LocationDataSize indicates the size of a LocationData object when marshalled
// into bytes.
const LocationDataSize = 18

// LocationData represents a measurement of the first responder's position. It
// implements SensorData.
type LocationData struct {

	// id indicates where this data point falls relative to all other data points
	// from the sensor.
	id byte

	// lat indicates the latitude of the first responder.
	lat float64

	// lon indicates the longitude of the first responder.
	lon float64
}

// CreateLocationData will instantiate and return a new LocationData object with
// the given id, representing a measurement of the first responder's location
// at the given latitude and longitude.
func CreateLocationData(id byte, lat float64, lon float64) *LocationData {
	return &LocationData{
		id:  id,
		lat: lat,
		lon: lon,
	}
}

// Id indicates the sequence id of this element of data among all data points
// generated by this sensor, as defined by SensorData.
func (d *LocationData) Id() byte {
	return d.id
}

// Type returns a value that uniquely identifies location sensors, as defined
// by SensorData.
func (d *LocationData) Type() byte {
	return LocationDataType
}

// String will return a string representation of the first responder's position,
// as defined by fmt.Stringer.
func (d *LocationData) String() string {
	return fmt.Sprintf("Location [%d]: %f, %f", d.id, d.lat, d.lon)
}

// NumBytes returns the number of bytes that a LocationData object is marshalled
// to, as defined by SensorData.
func (u *LocationData) NumBytes() int {
	return LocationDataSize
}

// ToBytes will marshall this data point into a slice of bytes, which can be
// transmitted across the network, as defined by SensorData.
func (d *LocationData) ToBytes() []byte {
	out := make([]byte, LocationDataSize)
	idx := 0
	out[idx] = d.Type()
	idx++
	out[idx] = d.Id()
	idx++
	for _, b := range float64ToBytes(d.lat) {
		out[idx] = b
		idx++
	}
	for _, b := range float64ToBytes(d.lon) {
		out[idx] = b
		idx++
	}
	return out
}

// LocationUnmarshaller implements SensorUnmarshaller, and is used to umarshall
// recieved bytes into a LocationData object.
type LocationUnmarshaller struct{}

// FromBytes takes hte given input bytes and returns a new LocationData object
// made from the data stored in the bytes, as defined by SensorUnmarshaller.
func (u *LocationUnmarshaller) FromBytes(in []byte) SensorData {
	return &LocationData{
		id:  in[1],
		lat: bytesToFloat64(in[2:10]),
		lon: bytesToFloat64(in[10:LocationDataSize]),
	}
}
