package sensor

import (
	"fmt"
	"time"

	"github.com/mcprice30/wmn/data"
)

type SensorStream interface {
	GetData() data.SensorData
	Interval() time.Duration
}

func Run(s SensorStream) {
	ticker := time.NewTicker(s.Interval())
	for {
		<-ticker.C
		fmt.Println(s.GetData())
	}
}

func intervalFromString(intervalStr string) time.Duration {
	if interval, err := time.ParseDuration(intervalStr); err == nil {
		return interval
	} else {
		panic(err)
	}
}
