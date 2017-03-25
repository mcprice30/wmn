package data

import (
	"fmt"
)

// HeartRateDataType indicates what the first byte of HeartRateData will be
// when marshalled into bytes for transmission.
const HeartRateDataType byte = 0

// HeartRateDataSize indicates the size of a HeartRateData object when
// marshalled to bytes.
const HeartRateDataSize int = 10

// HeartRateData represents a heart rate measuremtn from the heart rate
// monitor on the first responder. It implements SensorData.
type HeartRateData struct {

	// id indicates where this data point falls relative to all other data points
	// from the sensor.
	id byte

	// heartRate indicates the first responder's heart rate.
	heartRate float64
}

// CreateHeartRateData will instantiate a new HeartRateData object with the
// given id, representing a measurement from the first responder's heart rate
// monitor.
func CreateHeartRateData(id byte, heartRate float64) *HeartRateData {
	return &HeartRateData{
		id:        id,
		heartRate: heartRate,
	}
}

// Id indiates the sequence id of this element of data among all data points
// generated by this sensor, as defined by SensorData.
func (d *HeartRateData) Id() byte {
	return d.id
}

// Type returns a value that uniquely identifies heart rate sensors, as defined
// by SensorData.
func (d *HeartRateData) Type() byte {
	return HeartRateDataType
}

// String will return a string representation of the heart rate measurement, as
// defined by fmt.Stringer.
func (d *HeartRateData) String() string {
	return fmt.Sprintf("Heart Rate [%d]: %f", d.id, d.heartRate)
}

// NumBytes returns the number of bytes that a HeartRateData object is
// marshalled to, as defined by SensorData
func (u *HeartRateData) NumBytes() int {
	return HeartRateDataSize
}

// ToBytes will marshall this measurement into a slice of bytes, which can be
// transmitted across the network, as defined by SensorData.
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

// HeartRateUnmarshaller implements SensorUnmarshaller, and is used to
// unmarshall recieved bytes into a HeartRateData object.
type HeartRateUnmarshaller struct{}

// HeartRateDataFromBytes takes the given input bytes and returns a new
// HeartRateData object made from the data stored in the bytes.
func HeartRateDataFromBytes(in []byte) SensorData {
	return &HeartRateData{
		id:        in[1],
		heartRate: bytesToFloat64(in[2:HeartRateDataSize]),
	}
}
