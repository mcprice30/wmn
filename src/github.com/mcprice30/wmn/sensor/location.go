package sensor

import (
	"math/rand"
	"time"

	"github.com/mcprice30/wmn/data"
)

// LocationSensorInterval indicates how often data is generated by the GPS
// (location) sensor, before being sent to the sensor hub.
const LocationSensorInterval = "500ms"

// LocationSensor represents the location sensor attached to the first
// responder. It implements sensor.SensorStream, allowing for the data it
// generates to be sent to the sensor hub.
type LocationSensor struct {
	// interval indicates how often data is generated.
	interval time.Duration
	// id is the id of the upcoming data segment, used in sequencing data.
	id byte
}

// CreateLocationSensor will create a new instance of LocationSensor.
func CreateLocationSensor() *LocationSensor {
	return &LocationSensor{
		interval: intervalFromString(LocationSensorInterval),
	}
}

// Interval indicates how regularly data is generated by the sensor.
// It implements SensorStream.
func (s *LocationSensor) Interval() time.Duration {
	return s.interval
}

// GetData will generate a new data point from the sensor.
// It implements SensorStream.
func (s *LocationSensor) GetData() data.SensorData {
	lat := rand.Float64()*80.0 + 100.0
	lon := rand.Float64()*80.0 + 100.0
	defer s.incrementId()
	return data.CreateLocationData(s.id, lat, lon)
}

// incrementId will increase the packet id, used in sequencing data.
func (s *LocationSensor) incrementId() {
	s.id++
}