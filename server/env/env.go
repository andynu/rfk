// Environment sampling
package env

import (
	"log"
	"sync"
	"time"

	"github.com/andynu/rfk/server/karma"
)

var sensors_mu sync.Mutex
var sensors []Sensor
var samples []Sample

func init() {
	updateEnv()
}

func Updater() {
	tick := time.NewTicker(time.Hour).C
	for {
		<-tick
		log.Printf("env: start update")
		updateEnv()
		log.Printf("env: end update")
	}
}

func updateEnv() {
	oldSamples := samples
	newSamples := senseEnv()

	diffSamples := SamplesDifference(newSamples, oldSamples)
	logSamples(diffSamples)
	samples = newSamples
}

func SamplesDifference(a []Sample, b []Sample) []Sample {
	var diffSamples []Sample

	//	log.Printf("mySamples %v", a)
	//	log.Printf("otherSamples %v", b)

	for _, mySample := range a {
		isNew := true
		for _, otherSample := range b {
			if mySample.SensorName == otherSample.SensorName && mySample.Value == otherSample.Value {
				isNew = false
			}
		}
		//		log.Printf("env: new=%t (%s, %s)", isNew, mySample.SensorName,mySample.Value)
		if isNew {
			diffSamples = append(diffSamples, mySample)
		}
	}
	return diffSamples
}

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
	samples = senseEnv()
}

func senseEnv() []Sample {
	var curSamples []Sample
	for _, sensor := range sensors {
		for _, sample := range sensor.Sample() {
			curSamples = append(curSamples, sample)
		}
	}
	return curSamples
}

func logSamples(s []Sample) {
	for _, sample := range s {
		karma.LogEnv(sample.SensorName, sample.Value)
	}
}
