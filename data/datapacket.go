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

// DataPacket encapsulates a packet containing actual data to be sent across
// the network.
type DataPacket struct {
	// Header consists of the data packet's header, with information for both
	// the network and transport layers.
	Header PacketHeader

	// Body consists of the data from all the sensors that need to be transmitted
	// across the network.
	Body []SensorData
}

// NumBytes returns the number of bytes in a given data packet, including the
// header, once marshalled to bytes for transmission.
func (p *DataPacket) NumBytes() int {
	out := PacketHeaderBytes
	for _, data := range p.Body {
		out += data.NumBytes()
	}
	return out
}

// ToBytes will marshall the given data packet into bytes, to be sent across
// the network.
func (p *DataPacket) ToBytes() []byte {
	p.Header.NumBytes = uint8(p.NumBytes())
	out := make([]byte, p.NumBytes())
	idx := 0
	for _, b := range p.Header.ToBytes() {
		out[idx] = b
		idx++
	}

	for _, entry := range p.Body {
		for _, b := range entry.ToBytes() {
			out[idx] = b
			idx++
		}
	}

	return out
}

func DataPacketFromBytes(in []byte) *DataPacket {
	header := *PacketHeaderFromBytes(in[:PacketHeaderBytes])
	body := []SensorData{}
	idx := PacketHeaderBytes
	for idx < len(in) {
		res := SensorDataFromBytes(in[idx:])
		idx += res.NumBytes()
		body = append(body, res)
	}
	return &DataPacket{
		Header: header,
		Body:   body,
	}
}
