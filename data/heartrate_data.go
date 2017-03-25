package data

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
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
	out[0] = d.Type()
	out[1] = d.Id()
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, ByteOrder, d.heartRate); err != nil {
		panic(err)
	}
	for i, b := range buf.Bytes() {
		out[i+2] = b
	}
	return out
}

type HeartRateUnmarshaller struct {}

func (u *HeartRateUnmarshaller) FromBytes(in []byte) SensorData {
	id := in[1]
	bits := ByteOrder.Uint64(in[2:])
	heartRate := math.Float64frombits(bits)
	return &HeartRateData {
		id: id,
		heartRate: heartRate,
	}
}

func (u *HeartRateUnmarshaller) NumBytes() int {
	return int(HeartRateDataSize)
}

