package main

// Karma Module
type KarmaModule struct{}

// Initialization
func (m *KarmaModule) Init() {
	// launch updater go routine
	go m.updateTicker()
}

// Internals
func (m *KarmaModule) updateTicker() {
	// run update() every now and then.
}

func (m *KarmaModule) update() {
	// update the current state.
}

// KarmaView
func (m *KarmaModule) RelevantSongs() []Impression {
	var relevant []Impression
	return relevant
}

// KarmaWriter
func (m *KarmaModule) Impress(i Impression) {
}
