package sensor

import (
	"github.com/mcprice30/wmn/data"
)

type Sensor interface {
	GetData() data.SensorData
	Wait()
	Start()
}
