package graph_test

import (
	"algorithms/graph"
	"testing"

	"github.com/stretchr/testify/assert"
)

var g = *graph.NewGraph(true)

func TestShouldAddVertices(t *testing.T) {
	g.AddVertices([]int{3, 2, 5})

	assert.NotEmpty(t, g.Vertices)
	assert.NotNil(t, g.Vertices[3])
}

func TestShouldAddEdges(t *testing.T) {
	g.AddVertices([]int{3, 2, 5})
	g.AddEdges(map[int]int{
		3: 5,
		2: 3,
		5: 2,
	})
	v2Key := g.Vertices[5].Key

	assert.NotNil(t, g.Vertices[3].Vertices[v2Key])
}
