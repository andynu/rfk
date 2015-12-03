// Environment sampling
package env

import (
	"log"
	"sync"
	"time"
)

var sensors_mu sync.Mutex
var sensors []Sensor
var samples []Sample

type Sample struct {
	Timestamp  time.Time
	SensorName string
	Value      string
}

type Sensor interface {
	Sample() []Sample
}

func RegisterSensor(sensor Sensor) {
	sensors_mu.Lock()
	sensors = append(sensors, sensor)
	sensors_mu.Unlock()
}

func Prime() {
	for _, sensor := range sensors {
		for _, sample := range sensor.Sample() {
			samples = append(samples, sample)
		}
	}
}

func LogFull() {
	for _, sample := range samples {
		log.Printf("%v\t%q\t%q", sample.Timestamp, sample.SensorName, sample.Value)
	}
}
