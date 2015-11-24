// Event hooks between packages
package observer

import (
	"log"
)

var callbacks map[string][]func(interface{})

func init() {
	callbacks = make(map[string][]func(interface{}))
}

func Observe(channel string, callback func(interface{})) {
	log.Printf("Observe: %s", channel)
	callbacks[channel] = append(callbacks[channel], callback)
}

func Notify(channel string, message interface{}) {
	log.Printf("Notify: %s", channel)
	for _, callback := range callbacks[channel] {
		callback(message)
	}
}
