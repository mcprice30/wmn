package data

import (
	"encoding/binary"
)

var ByteOrder = binary.BigEndian

type SensorData interface {
	Id() byte
	Type() byte
	ToBytes() []byte
	String() string
}

type SensorUnmarshaller interface {
	FromBytes([]byte) SensorData
	NumBytes() int
}
