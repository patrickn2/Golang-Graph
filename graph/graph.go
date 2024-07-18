package graph

import "fmt"

type Graph struct {
	AdjacencyList map[string]map[string]struct{}
}

func NewGraph() *Graph {
	return &Graph{
		AdjacencyList: map[string]map[string]struct{}{},
	}
}

func (g *Graph) AddVertex(name string) {
	if _, ok := g.AdjacencyList[name]; ok {
		return
	}
	g.AdjacencyList[name] = make(map[string]struct{})
}

func (g *Graph) AddEdge(v1, v2 string) {
	_, ok := g.AdjacencyList[v1]
	_, ok2 := g.AdjacencyList[v2]

	if !ok {
		g.AddVertex(v1)
	}
	if !ok2 {
		g.AddVertex(v2)
	}

	g.AdjacencyList[v1][v2] = struct{}{}
	g.AdjacencyList[v2][v1] = struct{}{}

}

func (g *Graph) Display() {
	for v1, value := range g.AdjacencyList {
		fmt.Printf("%s -> ", v1)
		count := 0
		for v2 := range value {
			if count > 0 {
				fmt.Printf(", ")
			}
			fmt.Printf("%s", v2)
			count++
		}
		fmt.Println()
	}
}

func (g *Graph) HasEdge(v1, v2 string) bool {
	_, ok := g.AdjacencyList[v1]
	if !ok {
		return false
	}

	_, ok = g.AdjacencyList[v1][v2]

	if !ok {
		return false
	}

	_, ok = g.AdjacencyList[v2]
	if !ok {
		return false
	}

	_, ok = g.AdjacencyList[v2][v1]

	return ok
}

func (g *Graph) RemoveEdge(v1, v2 string) {
	delete(g.AdjacencyList[v1], v2)
	delete(g.AdjacencyList[v2], v1)
}

func (g *Graph) RemoveVertex(v string) {
	if _, ok := g.AdjacencyList[v]; !ok {
		return
	}
	for key := range g.AdjacencyList[v] {
		g.RemoveEdge(key, v)
	}
	delete(g.AdjacencyList, v)
}
