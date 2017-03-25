package data

import (
	"bytes"
	"encoding/binary"
	"math"
)

func float64ToBytes(in float64) []byte {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, ByteOrder, in); err != nil {
		panic(err)
	}
	return buf.Bytes()
}

func bytesToFloat64(in []byte) float64 {
	bits := ByteOrder.Uint64(in[2:])
	return math.Float64frombits(bits)
}

