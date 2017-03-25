package sensor

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/mcprice30/wmn/data"
)

type OxygenSensor struct {
	interval time.Duration
	ticker   *time.Ticker
	id       byte
}

func CreateOxygenSensor(ms int) *OxygenSensor {
	interval, err := time.ParseDuration(fmt.Sprintf("%dms", ms))
	if err != nil {
		panic(err)
	}
	return &OxygenSensor{
		interval: interval,
		id:       0,
	}
}

func (s *OxygenSensor) GetData() data.SensorData {
	pct := rand.Float64() * 100.0
	defer s.incId()
	return data.CreateOxygenData(s.id, pct)
}

func (s *OxygenSensor) incId() {
	s.id++
}

func (s *OxygenSensor) Wait() {
	if s.ticker != nil {
		<-s.ticker.C
	} else {
		panic("Cannot 'Wait()' until 'Start()' is called")
	}
}

func (s *OxygenSensor) Start() {
	s.id = 0
	if s.ticker != nil {
		s.ticker.Stop()
	}
	s.ticker = time.NewTicker(s.interval)
}
