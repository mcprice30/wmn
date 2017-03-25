package data

import (
	"fmt"
)

const LocationDataType byte = 1
const LocationDataSize byte = 18

type LocationData struct {
	id  byte
	lat float64
	lon float64
}

func CreateLocationData(id byte, lat float64, lon float64) *LocationData {
	return &LocationData{
		id:  id,
		lat: lat,
		lon: lon,
	}
}

func (d *LocationData) Id() byte {
	return d.id
}

func (d *LocationData) Type() byte {
	return LocationDataType
}

func (d *LocationData) String() string {
	return fmt.Sprintf("[%d]: %f, %f", d.id, d.lat, d.lon)
}

func (d *LocationData) ToBytes() []byte {
	out := make([]byte, LocationDataSize)
	idx := 0
	out[idx] = d.Type()
	idx++
	out[idx] = d.Id()
	idx++
	for _, b := range float64ToBytes(d.lat) {
		out[idx] = b
		idx++
	}
	for _, b := range float64ToBytes(d.lon) {
		out[idx] = b
		idx++
	}
	return out
}

type LocationUnmarshaller struct{}

func (u *LocationUnmarshaller) FromBytes(in []byte) SensorData {
	return &LocationData{
		id:  in[1],
		lat: bytesToFloat64(in[2:10]),
		lon: bytesToFloat64(in[10:]),
	}
}

func (u *LocationUnmarshaller) NumBytes() int {
	return int(LocationDataSize)
}
