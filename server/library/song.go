// Song representation
package library

import (
	"fmt"
	"sync"
)

type Song struct {
	Hash string
	Path string
	Rank float64
	pathNode
	sync.Mutex
	SongMeta
}

type SongMeta struct {
	Artist string
	Album  string
	Title  string
	genre  string
}

func (s *Song) String() string {
	return fmt.Sprintf("[Song %s %s]", s.Hash, s.Path)
}

func (s *Song) Meta() SongMeta {
	var meta SongMeta
	if s.Path != "" {
		meta, _ = metaData(s.Path)
	}
	return meta
}
