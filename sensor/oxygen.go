package sensor

import (
	"math/rand"
	"time"

	"github.com/mcprice30/wmn/data"
)

const OxygenInterval = "2000ms"

type OxygenSensor struct {
	interval time.Duration
	ticker   *time.Ticker
	id       byte
}

func CreateOxygenSensor() *OxygenSensor {
	return &OxygenSensor{
		interval: intervalFromString(OxygenInterval),
	}
}

func (s *OxygenSensor) Interval() time.Duration {
	return s.interval
}

func (s *OxygenSensor) GetData() data.SensorData {
	pct := rand.Float64() * 100.0
	defer s.incId()
	return data.CreateOxygenData(s.id, pct)
}

func (s *OxygenSensor) incId() {
	s.id++
}
