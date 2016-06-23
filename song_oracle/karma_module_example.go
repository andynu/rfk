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

// ------------------------------------------------

type ImpressionPipe struct {
	readers []chan Impression
}

func (p *ImpressionPipe) ReaderChan() chan Impression {
	c := make(chan Impression)
	p.readers = append(p.readers)
	return c
}
func (p *ImpressionPipe) Write(impression Impression) {
	for _, reader := range p.readers {
		reader <- impression
	}
}
