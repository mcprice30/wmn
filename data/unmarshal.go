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

// SensorDataFromBytes will attempt to unmarshall the given bytes to a
// SensorData instance. If unsuccessful, it will return nil.
func SensorDataFromBytes (in []byte) SensorData {
	switch in[0] {
	case GasDataType:
		return GasDataFromBytes(in)
	case HeartRateDataType:
		return HeartRateDataFromBytes(in)
	case LocationDataType:
		return LocationDataFromBytes(in)
	case OxygenDataType:
		return OxygenDataFromBytes(in)
	default:
		return nil
	}
}
