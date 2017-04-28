package data

import (
	"fmt"
)

// DataPacket encapsulates a packet containing actual data to be sent across
// the network.
type HelloPacket struct {
	// Header consists of the data packet's header, with information for both
	// the network and transport layers.
	Header PacketHeader

	// The number of bi-directional links in the packet
	NumBidirectional uint8

	// The list of bi-directional links this node has.
	BidirectionalLinks []ManetAddr

	// The number of addresses heard by this node.
	NumHeard uint8

	// HeardLinks
	HeardLinks []ManetAddr

	// The number of MPRs for this node.
	NumMPR uint8

	// MPRLinks
	MPRLinks []ManetAddr
}

// NumBytes returns the number of bytes in a given hello packet, including the
// header, once marshalled to bytes for transmission.
func (p *HelloPacket) NumBytes() int {
	out := PacketHeaderBytes
	// add NumMPR, NumHeard and NumBidirectional
	out += 3
	// add BidirectionalLinks and heard links
	out += 2 * (len(p.BidirectionalLinks) + len(p.HeardLinks) + len(p.MPRLinks))
	return out
}

// ToBytes will marshall the given data packet into bytes, to be sent across
// the network.
func (p *HelloPacket) ToBytes() []byte {
	p.Header.NumBytes = uint8(p.NumBytes())
	out := make([]byte, p.NumBytes())
	idx := 0
	for _, b := range p.Header.ToBytes() {
		out[idx] = b
		idx++
	}

	out[idx] = p.NumBidirectional
	idx++;

	for _, addr := range p.BidirectionalLinks {
		for _, b := range uint16ToBytes(uint16(addr)) {
			out[idx] = b
			idx++
		}
	}

	out[idx] = p.NumHeard
	idx++

	for _, addr := range p.HeardLinks {
		for _, b := range uint16ToBytes(uint16(addr)) {
			out[idx] = b
			idx++
		}
	}

	out[idx] = p.NumMPR
	idx++

	for _, addr := range p.MPRLinks {
		for _, b := range uint16ToBytes(uint16(addr)) {
			out[idx] = b
			idx++
		}
	}

	return out
}

func HelloPacketFromBytes(in []byte) *HelloPacket {
	header := *PacketHeaderFromBytes(in[:PacketHeaderBytes])
	idx := PacketHeaderBytes
	//fmt.Printf("UNMARSHALLING: %v\n", in)
	numBidirectional := in[idx]
	bidirectionalLinks := make([]ManetAddr, numBidirectional)
	//fmt.Printf("# Bi: %d at %d", in[idx], idx)
	idx++

	for i := uint8(0); i < numBidirectional; i++ {
		bidirectionalLinks[i] = ManetAddr(bytesToUint16(in[idx:idx+2]))
		idx += 2
	}

	numHeard := in[idx]
	heardLinks := make([]ManetAddr, numHeard)
	//fmt.Printf("# Uni: %d at %d", in[idx], idx)
	idx++

	for i := uint8(0); i < numHeard; i++ {
		heardLinks[i] = ManetAddr(bytesToUint16(in[idx:idx+2]))
		idx += 2
	}

	numMPR := in[idx]
	mprLinks := make([]ManetAddr, numMPR)
	//fmt.Printf("# MPR: %d at %d\n", in[idx], idx)
	idx++

	for i := uint8(0); i < numMPR; i++ {
		//fmt.Printf("Reading %d and %d\n", idx, idx+1)
		mprLinks[i] = ManetAddr(bytesToUint16(in[idx:idx+2]))
		idx += 2
	}

	return &HelloPacket{
		Header: header,
		NumBidirectional: numBidirectional,
		BidirectionalLinks: bidirectionalLinks,
		NumHeard: numHeard,
		HeardLinks: heardLinks,
		NumMPR: numMPR,
		MPRLinks: mprLinks,
	}
}

func (p *HelloPacket) String() string {
	out := fmt.Sprintf("Address: 0x%04x\n", p.Header.SourceAddress)
	out += fmt.Sprintf("Unidirectional: %v\n", p.HeardLinks)
	out += fmt.Sprintf("Bidirectional: %v\n", p.BidirectionalLinks)
	out += fmt.Sprintf("MPR: %v", p.MPRLinks)
	return out
}
