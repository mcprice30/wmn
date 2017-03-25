package data

import (
	"fmt"
)

const GasDataType byte = 3
const GasDataSize byte = 10

type GasData struct {
	id         byte
	percentage float64
}

func CreateGasData(id byte, percentage float64) *GasData {
	return &GasData{
		id:         id,
		percentage: percentage,
	}
}

func (d *GasData) Id() byte {
	return d.id
}

func (d *GasData) Type() byte {
	return GasDataType
}

func (d *GasData) String() string {
	return fmt.Sprintf("[%d]: %f", d.id, d.percentage)
}

func (d *GasData) ToBytes() []byte {
	out := make([]byte, GasDataSize)
	idx := 0
	out[idx] = d.Type()
	idx++
	out[idx] = d.Id()
	idx++
	for _, b := range float64ToBytes(d.percentage) {
		out[idx] = b
		idx++
	}
	return out
}

type GasUnmarshaller struct{}

func (u *GasUnmarshaller) FromBytes(in []byte) SensorData {
	return &GasData{
		id:         in[1],
		percentage: bytesToFloat64(in[2:]),
	}
}

func (u *GasUnmarshaller) GasBytes() int {
	return int(OxygenDataSize)
}
