package sensor

import (
	"math/rand"
	"time"

	"github.com/mcprice30/wmn/data"
)

const GasSensorInterval = "250ms"

type GasSensor struct {
	interval time.Duration
	ticker   *time.Ticker
	id       byte
}

func CreateGasSensor() *GasSensor {
	return &GasSensor{
		interval: intervalFromString(GasSensorInterval),
	}
}

func (s *GasSensor) Interval() time.Duration {
	return s.interval
}

func (s *GasSensor) GetData() data.SensorData {
	pct := rand.Float64() * 15.0
	defer s.incId()
	return data.CreateGasData(s.id, pct)
}

func (s *GasSensor) incId() {
	s.id++
}
