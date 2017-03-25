package sensor

import (
	"fmt"
)

type Sender struct {
	sensor Sensor
}

func CreateSender(s Sensor) *Sender {
	return &Sender{
		sensor: s,
	}
}

func (s *Sender) Run() {
	s.sensor.Start()
	for {
		s.sensor.Wait()
		fmt.Println(s.sensor.GetData())
	}
}
