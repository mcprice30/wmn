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

// uint16ToBytes converts the given uint16 to a slice of 2 bytes equivalent to
// the input.
func uint16ToBytes(in uint16) []byte {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, ByteOrder, in); err != nil {
		panic(err)
	}
	return buf.Bytes()
}

// bytesToUint16 takes a slice of bytes and returns the uint16 they represent.
func bytesToUint16(in []byte) uint16 {
	return ByteOrder.Uint16(in)
}

// combineTypeAndTTL will combine a PacketType field's lowest 3 bits and a
// TTL field's lowest 5 bits and combine them to a single byte with the
// PacketType's bits, followed by the TTL's bits.
func combineTypeAndTTL(packetType, ttl uint8) byte {
	out := ttl & 0x1F
	out |= (packetType & 0x07) << 5
	return out
}

// splitTypeAndTTL will split a merged type and TTL (as described in
// combineTypeAndTTL) and will return the two separate fields it represents.
func splitTypeAndTTL(in byte) (packetType, ttl uint8) {
	packetType = (in & 0xE0) >> 5
	ttl = (in & 0x1F)
	return
}
