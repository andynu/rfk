// Song representation
package library

import "fmt"

type Song struct {
	Hash string
	Path string
	Rank float64
	pathNode
}

func (s *Song) String() string {
	return fmt.Sprintf("[Song %s %s]", s.Hash, s.Path)
}
