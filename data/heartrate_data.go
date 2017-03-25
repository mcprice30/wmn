package data

import (
	"fmt"
)

const HeartRateDataType byte = 0
const HeartRateDataSize byte = 10

type HeartRateData struct {
	id byte
	heartRate float64
}

func CreateHeartRateData(id byte, heartRate float64) *HeartRateData {
	return &HeartRateData {
		id: id,
		heartRate: heartRate,
	}
}

func (d *HeartRateData) Id() byte {
	return d.id;
}

func (d *HeartRateData) Type() byte {
	return HeartRateDataType
}

func (d *HeartRateData) String() string {
	return fmt.Sprintf("[%d]: %f", d.id, d.heartRate)
}

func (d *HeartRateData) ToBytes() []byte {
	out := make([]byte, HeartRateDataSize)
	idx := 0
	out[idx] = d.Type()
	idx++
	out[idx] = d.Id()
	idx++
	for _, b := range float64ToBytes(d.heartRate) {
		out[idx] = b
		idx++
	}
	return out
}

type HeartRateUnmarshaller struct {}

func (u *HeartRateUnmarshaller) FromBytes(in []byte) SensorData {
	return &HeartRateData {
		id: in[1],
		heartRate: bytesToFloat64(in[2:]),
	}
}

func (u *HeartRateUnmarshaller) NumBytes() int {
	return int(HeartRateDataSize)
}

