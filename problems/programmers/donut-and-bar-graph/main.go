package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// edges := [][]int{{2, 3}, {4, 3}, {1, 1}, {2, 1}}
	// edges := [][]int{{4, 11}, {1, 12}, {8, 3}, {12, 7}, {4, 2}, {7, 11}, {4, 8}, {9, 6}, {10, 11}, {6, 10}, {3, 5}, {11, 1}, {5, 3}, {11, 9}, {3, 8}}
	edges := [][]int{{4, 3}, {2, 3}, {1, 1}, {2, 1}} // [2, 1, 1, 0]
	fmt.Println(solution(edges))
}

func solution(edges [][]int) []int {
	// initNode := edges[0][0]
	graph := NewGraph()
	for _, edge := range edges {
		graph.AddEdge(edge[0], edge[1])
	}
	initNode := graph.InitNode()

	donutCount := 0
	barCount := 0
	eightCount := 0
	for _, neighbor := range graph.Edges[initNode] {
		graph.Reset()
		graph.Traverse(neighbor)

		if graph.NumNodes == graph.NumEdges {
			donutCount++
		} else if graph.NumNodes == graph.NumEdges+1 {
			barCount++
		} else {
			size := graph.NumNodes / 2
			n := 2*size + 1
			e := 2*size + 2
			if n == graph.NumNodes && e == graph.NumEdges {
				eightCount++
			}
		}
		fmt.Println("stat", graph.NumNodes/2, neighbor, graph.NumNodes, graph.NumEdges)
	}
	fmt.Printf("%d, %d, %d", donutCount, barCount, eightCount)

	return []int{initNode, donutCount, barCount, eightCount}
}

type Graph struct {
	Edges    map[int][]int
	Visited  map[int]bool
	NumNodes int
	NumEdges int

	noInitNode map[int]struct{}
}

func NewGraph() *Graph {
	return &Graph{
		Edges:      make(map[int][]int),
		Visited:    make(map[int]bool),
		noInitNode: make(map[int]struct{}),
	}
}

func (g *Graph) AddEdge(from, to int) {
	g.Edges[from] = append(g.Edges[from], to)
	g.noInitNode[to] = struct{}{}
}

func (g *Graph) InitNode() int {
	maxNode := 0
	numEdges := 0
	for k, v := range g.Edges {
		if _, ok := g.noInitNode[k]; ok {
			continue
		}
		if maxNode == 0 {
			maxNode = k
			numEdges = len(v)
			continue
		}

		if numEdges < len(v) {
			maxNode = k
			numEdges = len(v)
			continue
		}
	}
	return maxNode
}

func (g *Graph) Traverse(node int) {
	g.Visited[node] = true
	g.NumNodes++
	for _, edge := range g.Edges[node] {
		g.NumEdges++
		if g.Visited[edge] {
			continue
		}
		g.Traverse(edge)
	}
}

func (g *Graph) Reset() {
	g.Visited = make(map[int]bool)
	g.NumEdges = 0
	g.NumNodes = 0
}

func (g *Graph) String() string {
	b, _ := json.MarshalIndent(g, "", "  ")
	return string(b)
}
