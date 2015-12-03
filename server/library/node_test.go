package library

import (
	"testing"
)

func TestSingeLink(t *testing.T) {
	var a, b Node
	a = &pathNode{Path: "a"}
	b = &pathNode{Path: "b"}

	a.Link(&b)
	a.Link(&b)
	b.Link(&a)
	b.Link(&a)

	if a.Links()[0] != &b {
		t.Error(`a.Link should include b`)
	}

	if len(a.Links()) != 1 {
		t.Error(`len a.Links() should be one`)
	}
}

func TestTraverse(t *testing.T) {
	var a, b, c Node
	a = &pathNode{Path: "a"}
	b = &pathNode{Path: "b"}
	c = &pathNode{Path: "c"}

	a.Link(&b)
	b.Link(&c)
	c.Link(&a)

	expectedNodeCount := 3
	nodeCount := 0

	traverseGraph(a, func(node Node, depth int) bool {
		nodeCount++
		return true
	})
	if expectedNodeCount != nodeCount {
		t.Error(`Expected a different number of nodes`)
	}
}
