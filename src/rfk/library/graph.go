package library

import (
	"log"
	"strings"
)

type Node interface {
	Links() []*Node
	Link(Node) error
}

type NodeSet map[*Node]bool

func (ns NodeSet) Nodes() []Node {
	var nodes []Node
	for node, _ := range ns {
		nodes = append(nodes, *node)
	}
	return nodes
}

func (ns NodeSet) Add(node *Node) {
	ns[node] = true
}

type PathNode struct {
	Path  string
	links []*Node
}

func (n PathNode) Links() []*Node {
	return n.links
}

func (n PathNode) Link(o Node) error {
	n.links = append(n.links, &o)
	return nil
}

func LoadGraph(songs []*Song) (roots []Node) {
	log.Println("Loading graph")
	var lastNode Node
	byPath := make(map[string]*Node)
	rootPaths := make(map[string]bool)

	traverseSongPaths(songs, func(partialPath string, pathDepth int, song *Song) {
		// Lookup || Create PathNode
		var node Node
		var nodePtr *Node = byPath[partialPath]

		if nodePtr != nil {
			node = *nodePtr
		} else {
			node = PathNode{Path: partialPath}
			byPath[partialPath] = &node
		}

		if pathDepth == 0 {
			rootPaths[partialPath] = true
		}

		// Link paths
		if lastNode != nil {
			node.Link(lastNode)
			lastNode.Link(node)
		}

		// Link song
		if song != nil {
			var songNode Node = song
			node.Link(songNode)
			songNode.Link(node)
		}

		lastNode = node
	})
	log.Println("Loading graph complete")

	//for path, node := range byPath {
	//	log.Printf("%v : %v => %v", path, &node, node)
	//}

	for path, _ := range rootPaths {
		node := *byPath[path]
		roots = append(roots, node)
	}

	return roots
}

func traverseSongPaths(
	songs []*Song,
	visitor func(partialPath string, pathDepth int, song *Song)) {

	limited := false
	i := 0
	for _, song := range songs {
		partialPath := ""
		dirs := strings.Split(song.Path, "/")
		for pathDepth, dir := range dirs {
			if partialPath != "/" {
				partialPath += "/"
			}
			partialPath += dir
			//log.Printf("i:%d %d path:%s", pathDepth, partialPath)
			leafSong := song
			if song.Path != partialPath {
				leafSong = nil
			}
			visitor(partialPath, pathDepth, leafSong)
		}
		i++
		if limited && i > 10 {
			break
		}
	}
}
