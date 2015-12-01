// Environment sampling
package env

import (
	"log"
	"sync"
)

var sensors_mu sync.Mutex
var sensors []Sensor
var samples []*Sample

type Sample struct {
	Timestamp  int
	SensorName string
	Value      string
}

type Sensor interface {
	Sample() *Sample
}

func RegisterSensor(sensor Sensor) {
	sensors_mu.Lock()
	sensors = append(sensors, sensor)
	sensors_mu.Unlock()
}

func Prime() {
	for _, sensor := range sensors {
		samples = append(samples, sensor.Sample())
	}
}

func LogFull() {
	for _, sample := range samples {
		log.Printf("%v\t%q\t%q", sample.Timestamp, sample.SensorName, sample.Value)
	}
}
