package sensor

import (
	"math/rand"
	"time"

	"github.com/mcprice30/wmn/data"
)

const LocationInterval = "500ms"

type LocationSensor struct {
	interval time.Duration
	ticker   *time.Ticker
	id       byte
}

func CreateLocationSensor() *LocationSensor {
	return &LocationSensor{
		interval: intervalFromString(LocationInterval),
	}
}

func (s *LocationSensor) GetData() data.SensorData {
	lat := rand.Float64()*80.0 + 100.0
	lon := rand.Float64()*80.0 + 100.0
	defer s.incId()
	return data.CreateLocationData(s.id, lat, lon)
}

func (s *LocationSensor) Interval() time.Duration {
	return s.interval
}

func (s *LocationSensor) incId() {
	s.id++
}
