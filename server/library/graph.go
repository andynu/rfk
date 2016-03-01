package library

import (
	"fmt"
	"log"
	"math"
	"sort"
	"strings"
	"github.com/andynu/rfk/server/observer"
)

const debug = false

type Graph struct{}

type Node interface {
	Links() []*Node
	Link(*Node) error
}

type pathNode struct {
	Path  string
	links map[*Node]bool
}

func (n *pathNode) Links() []*Node {
	var nodes []*Node
	for node, _ := range n.links {
		nodes = append(nodes, node)
	}
	return nodes
}

func (n *pathNode) Link(o *Node) error {
	if n.links == nil {
		n.links = make(map[*Node]bool)
	}
	if debug {
		log.Printf("Link: %v, %v", &n, o)
		if n.links[o] {
			log.Printf("skip link: %v, %v", n, *o)
		} else {
			log.Printf("link nodes: %v, %v", n, *o)
		}
	}

	n.links[o] = true
	return nil
}

func (n pathNode) String() string {
	return fmt.Sprintf("[pathNode %s]", n.Path)
}

func (g *Graph) LoadGraph(songs []*Song) (roots []Node) {
	log.Println("Loading graph")
	songNodeCount, pathNodeCount, visitCount := 0, 0, 0
	var lastDepth int
	var lastNode *Node
	byPath := make(map[string]*Node)
	rootPaths := make(map[string]bool)

	traverseSongPaths(songs, func(partialPath string, pathDepth int, song *Song) {
		visitCount++
		depthDiff := math.Abs(float64(lastDepth - pathDepth))
		if depthDiff > 2 {
			lastNode = nil
		}

		// Link song
		if lastNode != nil && song != nil {

			song.Link(lastNode)
			songNode := Node(song)
			(*lastNode).Link(&songNode)
			if debug {
				log.Printf("add song: %v, %v => %v", &lastNode, &song, song)
			}
			songNodeCount++

		} else {

			// Lookup || Create PathNode
			var node Node
			var nodePtr *Node = byPath[partialPath]
			if nodePtr != nil {
				node = *nodePtr
			} else {
				node = &pathNode{Path: partialPath}
				if debug {
					log.Printf("add node: %v => %v", &node, node)
				}
				byPath[partialPath] = &node
				pathNodeCount++
			}

			if pathDepth == 0 {
				rootPaths[partialPath] = true
			}

			// Link paths
			if lastNode != nil {
				node.Link(lastNode)

				(*lastNode).Link(&node)
			}
			lastDepth = pathDepth
			lastNode = &node
		}

	})
	log.Printf("Loading graph complete : songs=%d pathNodes=%d visits=%d", songNodeCount, pathNodeCount, visitCount)

	//for path, node := range byPath {
	//	log.Printf("%v : %v => %v", path, &node, node)
	//}

	for path, _ := range rootPaths {
		node := *byPath[path]
		roots = append(roots, node)
	}

	observer.Notify("graph.loaded", struct{}{})

	return roots
}

func traverseSongPaths(
	songs []*Song,
	visitor func(partialPath string, pathDepth int, song *Song)) {

	limited := debug
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
		if limited && i > 20 {
			break
		}
	}
}

// if the visit function returns false it will not traverse farther down that path.
func traverseGraph(root Node, visit func(node Node, depth int) bool) {
	visited := make(map[Node]bool)
	depth := 0
	_traverseGraph(root, depth, visited, visit)
}

func _traverseGraph(node Node, depth int, visited map[Node]bool, visit func(node Node, depth int) bool) {
	if node == nil || visited[node] {
		return
	}
	shouldRecurse := visit(node, depth)
	visited[node] = true
	if shouldRecurse {
		for _, link := range node.Links() {
			_traverseGraph(*link, depth+1, visited, visit)
		}
	}
}
func (song *Song) PathGraphImpress(impression int) {
	fimp := float64(impression)
	traverseGraph(song, func(node Node, depth int) bool {
		if depth > 3 {
			return false
		}
		s, ok := node.(*Song)
		if ok {
			s.Lock()
			s.Rank += fimp / math.Pow(2.0, float64(depth))
			s.Unlock()
		}
		return true
	})
}

func CountImpressionsByDepth(song *Song) {
	depthCounts := make(map[int]int)
	traverseGraph(song, func(node Node, depth int) bool {
		if depth > 5 {
			return false
		}
		//fmt.Printf("%d\t%s\n", depth, node)
		depthCounts[depth]++
		return true
	})

	var sortedKeys []int
	for depth, _ := range depthCounts {
		sortedKeys = append(sortedKeys, depth)
	}
	sort.Ints(sortedKeys)

	for _, depth := range sortedKeys {
		fmt.Printf("%d\t%d\n", depth, depthCounts[depth])
	}
}
