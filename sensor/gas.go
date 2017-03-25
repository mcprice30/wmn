package sensor

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/mcprice30/wmn/data"
)

type GasSensor struct {
	interval time.Duration
	ticker   *time.Ticker
	id       byte
}

func CreateGasSensor(ms int) *GasSensor {
	interval, err := time.ParseDuration(fmt.Sprintf("%dms", ms))
	if err != nil {
		panic(err)
	}
	return &GasSensor{
		interval: interval,
		id:       0,
	}
}

func (s *GasSensor) GetData() data.SensorData {
	pct := rand.Float64() * 15.0
	defer s.incId()
	return data.CreateGasData(s.id, pct)
}

func (s *GasSensor) incId() {
	s.id++
}

func (s *GasSensor) Wait() {
	if s.ticker != nil {
		<-s.ticker.C
	} else {
		panic("Cannot 'Wait()' until 'Start()' is called")
	}
}

func (s *GasSensor) Start() {
	s.id = 0
	if s.ticker != nil {
		s.ticker.Stop()
	}
	s.ticker = time.NewTicker(s.interval)
}
