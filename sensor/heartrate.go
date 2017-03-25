package sensor

import (
	"math/rand"
	"time"

	"github.com/mcprice30/wmn/data"
)

const HeartRateInterval = "1000ms"

type HeartRateSensor struct {
	interval time.Duration
	ticker   *time.Ticker
	id       byte
}

func CreateHeartRateSensor() *HeartRateSensor {
	return &HeartRateSensor{
		interval: intervalFromString(HeartRateInterval),
	}
}

func (s *HeartRateSensor) GetData() data.SensorData {
	hr := rand.Float64()*80.0 + 100.0
	defer s.incId()
	return data.CreateHeartRateData(s.id, hr)
}

func (s *HeartRateSensor) Interval() time.Duration {
	return s.interval
}

func (s *HeartRateSensor) incId() {
	s.id++
}
