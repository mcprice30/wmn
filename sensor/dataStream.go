package sensor

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/mcprice30/wmn/data"
)

type Sensor interface {
	GetData() data.SensorData
	Wait()
	Start()
}

type HeartRateSensor struct {
	interval time.Duration
	ticker *time.Ticker
	id byte
}

func CreateHeartRateSensor(ms int) *HeartRateSensor {
	interval, err := time.ParseDuration(fmt.Sprintf("%dms", ms))
	if err != nil {
		panic(err)
	}
	return &HeartRateSensor {
		interval: interval,
		id: 0,
	}
}

func (s *HeartRateSensor) GetData() data.SensorData {
	hr := rand.Float64() * 80.0 + 100.0
	defer s.incId()
	return data.CreateHeartRateData(s.id, hr)
}

func (s *HeartRateSensor) incId() {
	s.id++
}

func (s *HeartRateSensor) Wait() {
	if s.ticker != nil {
		<-s.ticker.C
	} else {
		panic("Cannot 'Wait()' until 'Start()' is called")
	}
}

func (s *HeartRateSensor) Start() {
	s.id = 0
	if s.ticker != nil {
		s.ticker.Stop()
	}
	s.ticker = time.NewTicker(s.interval)
}
