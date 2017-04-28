package data

type ManetAddr uint16

const BroadcastAddress ManetAddr = 0xFFFF

// MaxPacketBytes indicates the maximum number of bytes that can be sent in
// a packet.
const MaxPacketBytes = 128

// MaxTTL is the largest possible value in the TTL field.
const MaxTTL = 0x1F

// PacketHeaderBytes indicates the number of bytes that a PacketHeader is
// marshalled to.
const PacketHeaderBytes = 12

// PacketTypeData indicates that the given packet actually contains data that
// is being transmitted across the network.
const PacketTypeData = 0

// PacketTypeAck indicates that the given packet is an acknowledgement for
// another packet that was sent across the network.
const PacketTypeAck = 1

// PacketTypeControl indicates that the given packet is a control packet used
// in diagnosing or providing more information about the network.
const PacketTypeHello = 2

// PacketHeader represents the header of a packet to be sent across the network.
// This header will be used at both the transport and network layers.
type PacketHeader struct {
	// SourceAddress is a value that uniquely identifies the source of the
	// transmission.
	SourceAddress ManetAddr

	// Destination address is a value that uniquely identifies the intended
	// destination of this transmission.
	DestinationAddress ManetAddr

	// PreviousHop indicates is a value that uniquely identifies the node that
	// we recieved this packet from.
	PreviousHop ManetAddr

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

	SendKey uint16
}

// ToBytes will marshall the given packet header into bytes, to be sent across
// the network.
func (h *PacketHeader) ToBytes() []byte {
	out := make([]byte, PacketHeaderBytes)
	idx := 0
	for _, b := range uint16ToBytes(uint16(h.SourceAddress)) {
		out[idx] = b
		idx++
	}
	for _, b := range uint16ToBytes(uint16(h.DestinationAddress)) {
		out[idx] = b
		idx++
	}
	for _, b := range uint16ToBytes(uint16(h.PreviousHop)) {
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
	idx++
	for _, b := range uint16ToBytes(h.SendKey) {
		out[idx] = b
		idx++
	}
	return out
}

// PacketHeaderFromBytes will unmarshall the recieived bytes into a packet
// header, after being received from the network.
func PacketHeaderFromBytes(in []byte) *PacketHeader {
	out := &PacketHeader{}
	out.SourceAddress = ManetAddr(bytesToUint16(in[0:2]))
	out.DestinationAddress = ManetAddr(bytesToUint16(in[2:4]))
	out.PreviousHop = ManetAddr(bytesToUint16(in[4:6]))
	out.PacketType, out.TTL = splitTypeAndTTL(in[6])
	out.SequenceNumber = bytesToUint16(in[7:9])
	out.NumBytes = in[9]
	out.SendKey = bytesToUint16(in[10:12])
	return out
}
