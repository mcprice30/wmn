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

// MaxPacketBytes indicates the maximum number of bytes that can be sent in
// a packet.
const MaxPacketBytes = 128

// PacketHeaderBytes indicates the number of bytes that a PacketHeader is
// marshalled to.
const PacketHeaderBytes = 10

// PacketTypeData indicates that the given packet actually contains data that
// is being transmitted across the network.
const PacketTypeData = 0

// PacketTypeAck indicates that the given packet is an acknowledgement for
// another packet that was sent across the network.
const PacketTypeAck = 1

// PacketTypeControl indicates that the given packet is a control packet used
// in diagnosing or providing more information about the network.
const PacketTypeControl = 2

// PacketHeader represents the header of a packet to be sent across the network.
// This header will be used at both the transport and network layers.
type PacketHeader struct {
	// SourceAddress is a value that uniquely identifies the source of the
	// transmission.
	SourceAddress uint16

	// Destination address is a value that uniquely identifies the intended
	// destination of this transmission.
	DestinationAddress uint16

	// PreviousHop indicates is a value that uniquely identifies the node that
	// we recieved this packet from.
	PreviousHop uint16

	// TTL is the number of hops remaining on this packet before it is dropped,
	// used to prevent an improperly routed packet from staying in the network
	// forever. Though this is an 8 bit field, only the lower 5 bits will be
	// transmitted on the network.
	TTL uint8

	// PacketType is used to indicates what type of packet this is (for example,
	// an acknowledgement, control, or data packet). Though this is an 8 bit
	// field, only the lower 3 bits will be transmited on the network.
	PacketType uint8

	// SequenceNumber identifies where this packet's data falls relative to all
	// incoming packets. This is used both by the Flooding Protocol, and in the
	// transport layer for re-assembling data receieved.
	SequenceNumber uint16

	// NumBytes indicates the total size of the packet (including the header).
	NumBytes uint8
}

// ToBytes will marshall the given packet header into bytes, to be sent across
// the network.
func (h *PacketHeader) ToBytes() []byte {
	out := make([]byte, PacketHeaderBytes)
	idx := 0
	for _, b := range uint16ToBytes(h.SourceAddress) {
		out[idx] = b
		idx++
	}
	for _, b := range uint16ToBytes(h.DestinationAddress) {
		out[idx] = b
		idx++
	}
	for _, b := range uint16ToBytes(h.PreviousHop) {
		out[idx] = b
		idx++
	}
	out[idx] = combineTypeAndTTL(h.PacketType, h.TTL)
	idx++
	for _, b := range uint16ToBytes(h.SequenceNumber) {
		out[idx] = b
		idx++
	}
	out[idx] = h.NumBytes
	return out
}

// PacketHeaderFromBytes will unmarshall the recieived bytes into a packet
// header, after being received from the network.
func PacketHeaderFromBytes(in []byte) *PacketHeader {
	out := &PacketHeader{}
	out.SourceAddress = bytesToUint16(in[0:2])
	out.DestinationAddress = bytesToUint16(in[2:4])
	out.PreviousHop = bytesToUint16(in[4:6])
	out.PacketType, out.TTL = splitTypeAndTTL(in[6])
	out.SequenceNumber = bytesToUint16(in[7:9])
	out.NumBytes = in[9]
	return out
}

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
	unmarshaller := &ByteUnmarshaller{}
	header := *PacketHeaderFromBytes(in[:PacketHeaderBytes])
	body := []SensorData{}
	idx := PacketHeaderBytes
	for idx < len(in) {
		res := unmarshaller.Unmarshal(in[idx:])
		idx += res.NumBytes()
		body = append(body, res)
	}
	return &DataPacket{
		Header: header,
		Body:   body,
	}
}
