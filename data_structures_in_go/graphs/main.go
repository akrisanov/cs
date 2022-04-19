package main

import "fmt"

// Graph represents an adjacency list graph
type Graph struct {
	vertices []*Vertex
}

// Vertex represents a graph vertex
type Vertex struct {
	key      int
	adjacent []*Vertex
}

// AddVertex adds a Vertex to the graph
func (g *Graph) AddVertex(key int) {
	if contains(g.vertices, key) {
		err := fmt.Errorf("Vertex %v already exists in the graph", key)
		fmt.Println(err.Error())
		return
	}
	g.vertices = append(g.vertices, &Vertex{key: key})
}

// AddEdge adds an edge to the graph
func (g *Graph) AddEdge(from, to int) {
	// get vertex
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	// check error
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("(%v -> %v) invalid edge", from, to)
		fmt.Println(err.Error())
		return
	}
	if contains(fromVertex.adjacent, to) {
		err := fmt.Errorf("(%v -> %v) edge already exists", from, to)
		fmt.Println(err.Error())
		return
	}
	// add edge
	fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
}

// getVertex returns a pointer to the vertex with a key integer
func (g *Graph) getVertex(key int) *Vertex {
	for i, v := range g.vertices {
		if v.key == key {
			return g.vertices[i]
		}
	}
	return nil
}

// contains
func contains(s []*Vertex, key int) bool {
	for _, v := range s {
		if key == v.key {
			return true
		}
	}
	return false
}

// Print will print the adjacent list for each vertex of the graph
func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex %v: ", v.key)
		for _, v := range v.adjacent {
			fmt.Printf(" %v ", v.key)
		}
	}
	fmt.Println()
}

func main() {
	graph := &Graph{}

	for i := 0; i < 5; i++ {
		graph.AddVertex(i)
	}

	graph.AddEdge(1, 2)
	graph.AddEdge(1, 2)
	graph.AddEdge(6, 2)
	graph.AddEdge(3, 2)

	graph.Print()
}
