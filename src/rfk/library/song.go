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
}

func (s *Song) String() string {
	return fmt.Sprintf("[Song %s %s]", s.Hash, s.Path)
}
