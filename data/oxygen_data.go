package data

import (
	"fmt"
)

const OxygenDataType byte = 2
const OxygenDataSize byte = 10

type OxygenData struct {
	id         byte
	percentage float64
}

func CreateOxygenData(id byte, percentage float64) *OxygenData {
	return &OxygenData{
		id:         id,
		percentage: percentage,
	}
}

func (d *OxygenData) Id() byte {
	return d.id
}

func (d *OxygenData) Type() byte {
	return OxygenDataType
}

func (d *OxygenData) String() string {
	return fmt.Sprintf("[%d]: %f", d.id, d.percentage)
}

func (d *OxygenData) ToBytes() []byte {
	out := make([]byte, OxygenDataSize)
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

type OxygenUnmarshaller struct{}

func (u *OxygenUnmarshaller) FromBytes(in []byte) SensorData {
	return &HeartRateData{
		id:        in[1],
		heartRate: bytesToFloat64(in[2:]),
	}
}

func (u *OxygenUnmarshaller) NumBytes() int {
	return int(OxygenDataSize)
}
