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
package data

import (
	"fmt"
)

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

	// ToBytes will marshall this data point into bytes in order to be put in a
	// packet.
	ToBytes() []byte
}

// SensorUnmarshaller is used to unmarshal bytes from the network into a
// a specific SensorData implementation.
type SensorUnmarshaller interface {

	// FromBytes will extract a SensorData point from the recieved bytes.
	FromBytes([]byte) SensorData

	// NumBytes indicates the number of bytes this SensorUnmarshaller must read
	// to extract a SensorData object.
	NumBytes() int
}

// ByteUnmarshaller is used to unmarshall recieved bytes into the appropriate
// sensor data types.
type ByteUnmarshaller struct {

	// Separate instances for each type to be resolved to.
	gasUnmarshaller       *GasUnmarshaller
	heartRateUnmarshaller *HeartRateUnmarshaller
	locationUnmarshaller  *LocationUnmarshaller
	oxygenUnmarshaller    *OxygenUnmarshaller
}

// CreateByteUnmarshaller instantiates a new ByteUnmarshaller.
func CreateByteUnmarshaller() *ByteUnmarshaller {
	return &ByteUnmarshaller{
		gasUnmarshaller:       &GasUnmarshaller{},
		heartRateUnmarshaller: &HeartRateUnmarshaller{},
		locationUnmarshaller:  &LocationUnmarshaller{},
		oxygenUnmarshaller:    &OxygenUnmarshaller{},
	}
}

// Unmarshall will attempt to unmarshall the given bytes to a SensorData
// instance. If unsuccessful, it will return nil.
func (u *ByteUnmarshaller) Unmarshal(in []byte) SensorData {
	switch in[0] {
	case GasDataType:
		return u.gasUnmarshaller.FromBytes(in)
	case HeartRateDataType:
		return u.heartRateUnmarshaller.FromBytes(in)
	case LocationDataType:
		return u.locationUnmarshaller.FromBytes(in)
	case OxygenDataType:
		return u.oxygenUnmarshaller.FromBytes(in)
	default:
		return nil
	}
}
