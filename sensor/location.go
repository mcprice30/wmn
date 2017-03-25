package sensor

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/mcprice30/wmn/data"
)

type LocationSensor struct {
	interval time.Duration
	ticker   *time.Ticker
	id       byte
}

func CreateLocationSensor(ms int) *LocationSensor {
	interval, err := time.ParseDuration(fmt.Sprintf("%dms", ms))
	if err != nil {
		panic(err)
	}
	return &LocationSensor{
		interval: interval,
		id:       0,
	}
}

func (s *LocationSensor) GetData() data.SensorData {
	lat := rand.Float64()*80.0 + 100.0
	lon := rand.Float64()*80.0 + 100.0
	defer s.incId()
	return data.CreateLocationData(s.id, lat, lon)
}

func (s *LocationSensor) incId() {
	s.id++
}

func (s *LocationSensor) Wait() {
	if s.ticker != nil {
		<-s.ticker.C
	} else {
		panic("Cannot 'Wait()' until 'Start()' is called")
	}
}

func (s *LocationSensor) Start() {
	s.id = 0
	if s.ticker != nil {
		s.ticker.Stop()
	}
	s.ticker = time.NewTicker(s.interval)
}
