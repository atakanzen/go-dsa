package search_test

import (
	"algorithms/search"
	"testing"

	"github.com/stretchr/testify/assert"
)

func makeList(min, max int) []int {
	a := make([]int, max)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func TestBinarySearchFound(t *testing.T) {
	list := makeList(1, 500000)
	item := 221017

	position := search.BinarySearch(list, item)
	want := 221016

	assert.Equal(t, want, position)
}

func TestBinarySearchNotFound(t *testing.T) {
	list := makeList(0, 1)
	item := 311098

	position := search.BinarySearch(list, item)

	assert.Equal(t, -1, position)
}
