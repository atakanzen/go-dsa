// Source https://codetree.dev/golang-graph-traversal/
//  Made some tweaks

package graph

import (
	"errors"
	"fmt"
)

type Vertex struct {
	Key      int
	Vertices map[int]*Vertex
}

type Graph struct {
	Vertices map[int]*Vertex
	directed bool
}

type node struct {
	v    *Vertex
	next *node
}

type queue struct {
	head *node
	tail *node
}

func NewVertex(key int) *Vertex {
	return &Vertex{
		Key:      key,
		Vertices: map[int]*Vertex{},
	}
}

func NewGraph(directed bool) *Graph {
	return &Graph{
		Vertices: map[int]*Vertex{},
		directed: directed,
	}
}

func (g *Graph) AddVertices(vertices []int) {
	for _, key := range vertices {
		v := NewVertex(key)
		g.Vertices[key] = v
	}
}

func (g *Graph) AddEdges(keys map[int]int) error {
	for k1, k2 := range keys {
		v1 := g.Vertices[k1]
		v2 := g.Vertices[k2]

		if v1 == nil || v2 == nil {
			return errors.New("vertice do not exist")
		}

		if _, ok := v1.Vertices[v2.Key]; ok {
			return nil
		}

		v1.Vertices[v2.Key] = v2
		if !g.directed && v1.Key != v2.Key {
			v2.Vertices[v1.Key] = v1
		}

		g.Vertices[v1.Key] = v1
		g.Vertices[v2.Key] = v2

	}

	return nil
}

func (q *queue) enqueue(v *Vertex) {
	n := &node{v: v}

	if q.tail == nil {
		q.head = n
		q.tail = n
		return
	}

	q.tail.next = n
	q.tail = n
}

func (q *queue) dequeue() *Vertex {
	n := q.head

	if n == nil {
		return nil
	}

	q.head = q.head.next

	if q.head == nil {
		q.tail = nil
	}

	return n.v
}

func BFS(g *Graph, startVertex *Vertex) {
	queue := &queue{}
	visitedVertices := map[int]bool{}

	currentVertext := startVertex
	for {
		fmt.Printf("Visited => %v\n", currentVertext.Key)
		visitedVertices[currentVertext.Key] = true

		for key, v := range currentVertext.Vertices {
			if !visitedVertices[key] {
				queue.enqueue(v)
			}
		}

		currentVertext = queue.dequeue()

		if currentVertext == nil {
			break
		}
	}
}
