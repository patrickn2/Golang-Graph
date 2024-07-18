package main

import (
	"fmt"

	"github.com/patrickn2/Golang-Graph/graph"
)

func main() {
	g := graph.NewGraph()
	g.AddVertex("A")
	g.AddVertex("B")
	g.AddVertex("C")

	g.AddEdge("A", "B")
	g.AddEdge("B", "C")
	g.Display()

	fmt.Println("A -> B", g.HasEdge("A", "B"))
	fmt.Println("A -> C", g.HasEdge("A", "C"))
	fmt.Println("B -> C", g.HasEdge("B", "C"))
	fmt.Println("C -> A", g.HasEdge("C", "A"))

	// g.RemoveEdge("A", "B")

	// fmt.Println("Removing A -> B", g.HasEdge("A", "B"))

	fmt.Println("Removing Vertex B")
	g.RemoveVertex("B")
	g.Display()
}
